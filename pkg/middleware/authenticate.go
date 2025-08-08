package middleware

import (
	"fmt"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
	"golang.org/x/crypto/bcrypt"
)

func RequiredEntries(entries ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			fmt.Println("the list of required entries are : ", entries)

			next.ServeHTTP(w, r)
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
			var responseJson util.StandardResponseJson
			responseJson.Msg = "No such user exists, check username"

			util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusOK)

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

		pwd_hash := models.FetchHashedPassword(userId)

		err := bcrypt.CompareHashAndPassword([]byte(pwd_hash), []byte(password))
		if err != nil {
			var responseJson util.StandardResponseJson
			responseJson.ErrDescription = "Incorrect password"
			responseJson.Msg = "unable to login"

			util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusForbidden)

			return
		}

		next.ServeHTTP(w, r)
	})
}

func IdentifyUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("jwt_token")
		if err != nil {
			var responseJson util.StandardResponseJson
			responseJson.Msg = "cant use the service"
			responseJson.ErrDescription = "Need to be signed in to use this route"

			util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusForbidden)

			return
		}
		tokenValue := token.Value

		xUser := util.DecryptJwtToken(w, r, tokenValue)

		r = util.PutUserInContext(r, xUser)

		next.ServeHTTP(w, r)
	})
}

func AnotherMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("this is a middleware ")

		// xUser := util.ExtractUserFromContext(r)

		// fmt.Println(xUser)

		next.ServeHTTP(w, r)
	})
}
