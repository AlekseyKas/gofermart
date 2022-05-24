package middlewarecustom

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
	"go.gofermart/cmd/server/storage"
)

func CheckCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		if req.URL.Path == "/api/user/register" || req.URL.Path == "/api/user/login" {
			next.ServeHTTP(rw, req)
		}
		var cookie *http.Cookie
		for _, cook := range req.Cookies() {
			if cook.Name == "gofermart" {
				cookie = cook
			}
		}
		if cookie != nil {
			defer req.Body.Close()
			out, err := ioutil.ReadAll(req.Body)
			if err != nil {
				logrus.Error("Error read body: ", err)
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
			u := storage.Users
			err = json.Unmarshal(out, &u)
			if err != nil {
				logrus.Error("Error unmarshal body: ", err)
				rw.WriteHeader(http.StatusBadRequest)
			}
			cookieDB, err := storage.DB.GetUser(u)
			if err != nil {
				logrus.Error(err)
			}
			if cookie.Value == cookieDB {
				next.ServeHTTP(rw, req)
			} else {
				rw.WriteHeader(http.StatusForbidden)
				return
			}
		}

	})
}
