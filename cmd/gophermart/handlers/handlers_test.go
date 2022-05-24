package handlers

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"go.gofermart/cmd/server/storage"
	"go.gofermart/internal/helpers"
)

// Возможные коды ответа:
//     200 — пользователь успешно зарегистрирован и аутентифицирован;
//     400 — неверный формат запроса;
//     409 — логин уже занят;
//     500 — внутренняя ошибка сервера.
func Test_register(t *testing.T) {

	type want struct {
		contentType string
		statusCode  int
	}
	tests := []struct {
		name        string
		body        []byte
		method      string
		url         string
		contentType string
		want        want
	}{
		// TODO: Add test cases.
		{
			name:        "success register",
			body:        []byte(`{"login": "user1", "password": "password"}`),
			method:      "POST",
			url:         "/api/user/register",
			contentType: "application/json",
			want: want{
				contentType: "application/json",
				statusCode:  200,
			},
		},
		{
			name:        "register duplicated",
			body:        []byte(`{"login": "user1", "password": "password"}`),
			method:      "POST",
			url:         "/api/user/register",
			contentType: "application/json",
			want: want{
				contentType: "application/json",
				statusCode:  409,
			},
		},
		{
			name:        "wrong type#1",
			body:        []byte(`{"loginn": "user1", "sword": "password"}`),
			method:      "POST",
			url:         "/api/user/register",
			contentType: "application/json",
			want: want{
				contentType: "application/json",
				statusCode:  400,
			},
		},
		{
			name:        "wrong type#2",
			body:        []byte(`{"login": "user1", "password": "password"}`),
			method:      "POST",
			url:         "/api/user/register",
			contentType: "plain/txt",
			want: want{
				contentType: "application/json",
				statusCode:  400,
			},
		},
	}
	r := chi.NewRouter()
	r.Route("/", Router)
	ts := httptest.NewServer(r)
	defer ts.Close()
	ctx, _ := context.WithCancel(context.Background())

	DBURL, id, _ := helpers.StartDB()

	storage.IDB = &storage.DB
	storage.IDB.InitDB(ctx, DBURL)
	logrus.Info(DBURL)

	defer helpers.StopDB(id)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := tt.body
			buff := bytes.NewBuffer(body)
			req, err := http.NewRequest(tt.method, ts.URL+tt.url, buff)
			req.Header.Set("Content-Type", tt.contentType)
			require.NoError(t, err)
			resp, err := http.DefaultClient.Do(req)
			require.Equal(t, tt.want.statusCode, resp.StatusCode)

			require.NoError(t, err)
			defer resp.Body.Close()
		})
	}
}

// Возможные коды ответа:

//     200 — пользователь успешно аутентифицирован;
//     400 — неверный формат запроса;
//     401 — неверная пара логин/пароль;
//     500 — внутренняя ошибка сервера.
func Test_login(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
	}
	tests := []struct {
		name        string
		body        []byte
		method      string
		url         string
		contentType string
		want        want
	}{
		// TODO: Add test cases.
		{
			name:        "success login",
			body:        []byte(`{"login": "user1", "password": "password"}`),
			method:      "POST",
			url:         "/api/user/login",
			contentType: "application/json",
			want: want{
				contentType: "application/json",
				statusCode:  200,
			},
		},
		{
			name:        "wrong type#1",
			body:        []byte(`{"loginn": "user1111", "sword": "password"}`),
			method:      "POST",
			url:         "/api/user/login",
			contentType: "application/json",
			want: want{
				contentType: "application/json",
				statusCode:  401,
			},
		},
		{
			name:        "wrong type#2",
			body:        []byte(`{"login": "user1", "password": "password"}`),
			method:      "POST",
			url:         "/api/user/login",
			contentType: "plain/txt",
			want: want{
				contentType: "application/json",
				statusCode:  400,
			},
		},
	}
	r := chi.NewRouter()
	r.Route("/", Router)
	ts := httptest.NewServer(r)
	defer ts.Close()
	ctx, _ := context.WithCancel(context.Background())

	DBURL, id, _ := helpers.StartDB()

	storage.IDB = &storage.DB
	storage.IDB.InitDB(ctx, DBURL)
	logrus.Info(DBURL)
	defer helpers.StopDB(id)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := tt.body
			buff := bytes.NewBuffer(body)
			req, err := http.NewRequest(tt.method, ts.URL+tt.url, buff)
			req.Header.Set("Content-Type", tt.contentType)
			require.NoError(t, err)
			resp, err := http.DefaultClient.Do(req)
			require.Equal(t, tt.want.statusCode, resp.StatusCode)

			require.NoError(t, err)
			defer resp.Body.Close()
		})
	}
}
