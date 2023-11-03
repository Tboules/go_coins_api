package middleware

import (
	"errors"
	"net/http"

	"github.com/Tboules/go_coins_api/api"
	"github.com/Tboules/go_coins_api/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizationError = errors.New("Invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username string = r.URL.Query().Get("username")
		var token string = r.Header.Get("Authorization")

		if username == "" || token == "" {
			log.Error(UnAuthorizationError)
			api.RequestErrorHandler(w, UnAuthorizationError)
			return
		}

		db, err := tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		loginDetails := (*db).GetUserLoginDetails(username)

		if loginDetails == nil || token != loginDetails.Token {
			log.Error(UnAuthorizationError)
			api.RequestErrorHandler(w, UnAuthorizationError)
			return
		}

		next.ServeHTTP(w, r)

	})
}
