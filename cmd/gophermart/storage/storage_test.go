package storage_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.gofermart/cmd/gophermart/storage"
	"go.gofermart/internal/helpers"
)

func TestDatabase_InitDB(t *testing.T) {

	t.Run("Init database", func(t *testing.T) {
		ctx, _ := context.WithCancel(context.Background())
		DBURL, id, _ := helpers.StartDB()
		defer helpers.StopDB(id)

		storage.IDB = &storage.DB
		err := storage.IDB.InitDB(ctx, DBURL)
		require.NoError(t, err)

	})

}

func TestDatabase_User(t *testing.T) {
	type args struct {
		u storage.User
	}
	type want struct {
		cookie storage.Cookie
	}
	tests := []struct {
		name string
		// fields  fields

		args args
		want storage.Cookie
	}{
		{
			name: "first",
			args: args{storage.User{
				Login:    "user",
				Password: "password",
			}},
			want: storage.Cookie{
				Name:   "gofermart",
				Value:  "xxx",
				MaxAge: 864000,
			},
		},
		{
			name: "second",
			args: args{storage.User{
				Login:    "user2",
				Password: "password",
			}},
			want: storage.Cookie{
				Name:   "gofermart",
				Value:  "xxxxx",
				MaxAge: 99999,
			},
		},
		// TODO: Add test cases.
	}
	ctx, _ := context.WithCancel(context.Background())
	DBURL, id, _ := helpers.StartDB()
	defer helpers.StopDB(id)

	storage.IDB = &storage.DB
	storage.IDB.InitDB(ctx, DBURL)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cookie, err := storage.DB.CreateUser(tt.args.u)
			require.NoError(t, err)
			require.NotEmpty(t, cookie.Value)
			require.NotEmpty(t, cookie.MaxAge)
			require.NotEqual(t, cookie.MaxAge, tt.want.MaxAge)
			require.Equal(t, cookie.Name, tt.want.Name)

			s, errg := storage.DB.GetUser(tt.args.u)
			require.NoError(t, errg)
			require.Equal(t, s, cookie.Value)

		})
	}
}

func TestHashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "first",
			args: args{"test"},
		},
		{
			name: "second",
			args: args{"testtesttest"},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := storage.HashPassword(tt.args.password)
			require.Equal(t, 60, len(s))
			require.NoError(t, err)
			result := storage.CheckPasswordHash(tt.args.password, s)
			if !result {
				t.Errorf("Invalid hash %s password %s", s, tt.args.password)
			}
		})
	}
}
