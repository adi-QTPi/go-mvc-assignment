package middleware

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
	"golang.org/x/crypto/bcrypt"
)

func RequiredEntries(entries ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var err error
			contentType := r.Header.Get("Content-Type")
			if strings.HasPrefix(contentType, "multipart/form-data") {
				const maxMemory = 32 << 20
				err = r.ParseMultipartForm(maxMemory)
			} else {
				err = r.ParseForm()
			}
			if err != nil {
				http.Error(w, fmt.Sprintf("Error parsing the form %v", err), http.StatusInternalServerError)
				return
			}
			ok := true
			for _, v := range entries {
				if r.Form.Get(v) == "" {
					ok = false
				}
			}

			if ok {
				next.ServeHTTP(w, r)
				return
			}

			popup := util.Popup{
				Msg:     "Enter ALL the fields carefully",
				IsError: true,
			}

			requestFrom := r.Referer()
			util.InsertPopupInFlash(w, r, popup)
			util.RedirectToSite(w, r, requestFrom)

			// var responseJson util.StandardResponseJson
			// responseJson.Msg = "Data reading failed"
			// responseJson.ErrDescription = fmt.Sprintf("Enter the given fields in x-www-form-urlencoded format only : %v", entries)
			// util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusBadRequest)
		})
	}
}

func CheckIfUserExists(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		userName := r.Form.Get("user_name")
		// password := r.Form.Get("password")
		var userId string
		isThere, err, userId := models.GetUserByUsername(userName)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error running the query for get user by username: %v", err), http.StatusInternalServerError)
			return
		}
		if !isThere {
			toLoginPage := util.Popup{
				Msg:     "User Not Found... Signup to Create New Account !",
				IsError: true,
			}

			err := util.InsertPopupInFlash(w, r, toLoginPage)
			if err != nil {
				http.Error(w, fmt.Sprintf("error in identifyUser middleware : %v", err), http.StatusInternalServerError)
				return
			}

			util.RedirectToSite(w, r, "/login")

			return
		}

		r = util.PutInContext(r, "user_id", userId)

		next.ServeHTTP(w, r)
	})
}

func CheckPassword(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := util.ExtractFromContext(r, "user_id")

		r.ParseForm()
		password := r.Form.Get("password")

		passwordHash := models.FetchHashedPassword(userId)

		err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
		if err != nil {
			toLoginPage := util.Popup{
				Msg:     "Incorrect Password... Try Again",
				IsError: true,
			}

			err := util.InsertPopupInFlash(w, r, toLoginPage)
			if err != nil {
				http.Error(w, fmt.Sprintf("error in identifyUser middleware : %v", err), http.StatusInternalServerError)
				return
			}

			util.RedirectToSite(w, r, "/login")

			return
		}

		next.ServeHTTP(w, r)
	})
}

func IdentifyUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("jwt_token")

		if err != nil {
			toLoginPage := util.Popup{
				Msg:     "Log In to continue !",
				IsError: false,
			}

			err := util.InsertPopupInFlash(w, r, toLoginPage)
			if err != nil {
				http.Error(w, fmt.Sprintf("error in identifyUser middleware : %v", err), http.StatusInternalServerError)
				return
			}

			util.RedirectToSite(w, r, "/login")
			return
		}
		tokenValue := token.Value

		xUser := util.DecryptJwtToken(w, r, tokenValue)

		r = util.PutUserInContext(r, xUser)

		next.ServeHTTP(w, r)
	})
}

func AddUserInfoInContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("jwt_token")

		if err == nil {
			tokenValue := token.Value

			xUser := util.DecryptJwtToken(w, r, tokenValue)

			r = util.PutUserInContext(r, xUser)
		}

		next.ServeHTTP(w, r)
	})
}

func RestrictToRoles(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			xUser := util.ExtractUserFromContext(r)

			if role := xUser.Role; slices.Contains(allowedRoles, role) {
				next.ServeHTTP(w, r)
				return
			}

			popup := util.Popup{
				Msg:     "You are NOT Authorised to Access this Service",
				IsError: true,
			}

			// requestFrom := r.Referer()
			// fmt.Println(requestFrom)
			util.InsertPopupInFlash(w, r, popup)
			util.RedirectToSite(w, r, "/error")

			// var responseJson util.StandardResponseJson
			// responseJson.Msg = "User is Unauthorised"
			// responseJson.ErrDescription = fmt.Sprintf("this page is for users with roles %v only.", allowedRoles)
			// util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusForbidden)
		})
	}
}
