package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"

	"go.gofermart/cmd/server/storage"
	"go.gofermart/internal/middlewarecustom"
)

func Router(r chi.Router) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middlewarecustom.CheckCookie)

	//регистрация пользователя
	r.Post("/api/user/register", register())
	// аутентификация пользователя
	r.Post("/api/user/login", login())
	// загрузка пользователем номера заказа для расчёта
	r.Post("/api/user/orders", setOrder())
	// запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа
	r.Post("/api/user/balance/withdraw", withdrawOrder())
	// получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях
	r.Get("/api/user/orders", getOrders())
	// получение текущего баланса счёта баллов лояльности пользователя
	r.Get("/api/user/balance", getBalance())
	// получение информации о выводе средств с накопительного счёта пользователем
	r.Get("/api/user/balance/withdrawals", withdraw())

}

func register() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()

		if !strings.Contains(req.Header.Get("Content-Type"), "application/json") {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

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
		if u.Login == "" || u.Password == "" {
			logrus.Error("Wrong format of user or password.")
			rw.WriteHeader(http.StatusBadRequest)
		}
		err409 := errors.New("ERROR: duplicate key value violates unique constraint \"users_login_key\" (SQLSTATE 23505)")
		cookie, err := storage.DB.CreateUser(u)

		switch {
		case err == nil:
			logrus.Info("User added: ", u.Login)
			http.SetCookie(rw, &http.Cookie{Name: cookie.Name, Value: cookie.Value, MaxAge: cookie.MaxAge})
			rw.WriteHeader(http.StatusOK)
		case err.Error() == err409.Error():
			logrus.Error("User already exist: ", u.Login)
			rw.WriteHeader(http.StatusConflict)
		default:
			rw.WriteHeader(http.StatusInternalServerError)
		}
	}
}

//login users
func login() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		if !strings.Contains(req.Header.Get("Content-Type"), "application/json") {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
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
		if u.Login == "" || u.Password == "" {
			logrus.Error("Wrong format of user or password.")
			rw.WriteHeader(http.StatusBadRequest)
		}
		// err := storage.DB.AuthUser(u)

	}
}

func setOrder() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {

	}
}

func withdrawOrder() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {

	}
}

func withdraw() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {

	}
}
func getOrders() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {

	}
}

func getBalance() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		logrus.Info("balance")
	}
}
