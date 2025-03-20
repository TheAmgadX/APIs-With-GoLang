package middleware

import (
	"errors"
	"net/http"

	"github.com/TheAmgadX/APIs-With-GoLang/api"
	"github.com/TheAmgadX/APIs-With-GoLang/internal/tools"
	log "github.com/sirupsen/logrus"
)

var ErrUnAuthorized error = errors.New("Invalid Username Or Token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		token := r.Header.Get("Authorization")

		var err error

		if username == "" || token == "" {
			log.Error(ErrUnAuthorized)
			api.RequestErrorHandler(w, ErrUnAuthorized)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()

		if err != nil {
			log.Error("Internal Server Error!")
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails

		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(ErrUnAuthorized)
			api.RequestErrorHandler(w, ErrUnAuthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
