package handlers

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
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

	}
}

func login() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {

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

	}
}
