package middleware

import (
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
	"golang.org/x/crypto/bcrypt"
)

func VerifyDuplicatePassword(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form in verify duplicate password middleware", http.StatusBadRequest)
			return
		}

		pwd := r.Form.Get("pwd")
		re_pwd := r.Form.Get("re_pwd")

		if pwd == re_pwd {
			next.ServeHTTP(w, r)
			return
		}

		popup := util.Popup{
			Msg:     "Passwords don't Match, try again",
			IsError: true,
		}

		util.InsertPopupInFlash(w, r, popup)

		requestFrom := r.Referer()
		util.InsertPopupInFlash(w, r, popup)
		util.RedirectToSite(w, r, requestFrom)

		// var responseJson util.StandardResponseJson

		// responseJson.ErrDescription = "the 2 passwords dont match."
		// responseJson.Msg = "failed to signup"

		// util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusBadRequest)
	})
}

func HashPasword(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form in hash password middleware", http.StatusBadRequest)
			return
		}

		password := r.Form.Get("pwd")

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
		if err != nil {
			http.Error(w, "Failed to hash password, error from bcrypt in the middleware hash password", http.StatusBadRequest)
			return
		}

		r.Form.Set("pwd", string(hashedPassword))

		next.ServeHTTP(w, r)
	})
}
