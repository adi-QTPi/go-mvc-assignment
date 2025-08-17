package middleware

import (
	"fmt"
	"net/http"
	"slices"
	"strings"
	"unicode"

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
		})
	}
}

func CheckIfUserExists(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		userName := r.Form.Get("user_name")
		var userId string
		isThere, _, userId := models.GetUserByUsername(userName)
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
			util.InsertPopupInFlash(w, r, popup)
			util.RedirectToSite(w, r, "/error")
		})
	}
}

func PasswordStrengthTest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form in verify duplicate password middleware", http.StatusBadRequest)
			return
		}
		pwd := r.Form.Get("pwd")

		popup := util.Popup{
			Msg:     "Password must be 10 characters long, combination of sMALL, Capital and Spec1al characters..",
			IsError: true,
		}

		var (
			hasUpperCase bool
			hasLowerCase bool
			hasDigit     bool
			hasSpecial   bool
		)

		for _, char := range pwd {
			if unicode.IsUpper(char) {
				hasUpperCase = true
			} else if unicode.IsLower(char) {
				hasLowerCase = true
			} else if unicode.IsDigit(char) {
				hasDigit = true
			} else {
				hasSpecial = true
			}
		}
		if len(pwd) < 10 || !hasUpperCase || !hasLowerCase || !hasDigit || !hasSpecial {
			util.InsertPopupInFlash(w, r, popup)
			util.RedirectToSite(w, r, "/signup")
		}
		next.ServeHTTP(w, r)
	})
}
