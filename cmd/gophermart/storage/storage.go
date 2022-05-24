package storage

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"time"

	"go.gofermart/cmd/gophermart/database"
	"go.gofermart/cmd/gophermart/storage/migrations"
	"go.gofermart/internal/config/migrate"
	"golang.org/x/crypto/bcrypt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type Cookie struct {
	Name   string `json:"Name,omitempty"`
	Value  string `json:"Value,omitempty"`
	Path   string `json:"Path,omitempty"`
	MaxAge int    `json:"MaxAge,omitempty"`
}

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Database struct {
	Con   *pgxpool.Pool
	Loger logrus.FieldLogger
	DBURL string
	Ctx   context.Context
}

var DB Database
var IDB Storage
var Users User

type Storage interface {
	InitDB(ctx context.Context, DBURL string) error
	CreateUser(u User) (cookie Cookie, err error)
	GetUser(u User) (string, error)
	// ReconnectDB() error
}

func (d *Database) GetUser(u User) (string, error) {
	var login string
	var password string
	row := d.Con.QueryRow(d.Ctx, "SELECT * FROM users WHERE login = $1", u.Login)
	row.Scan(&login, &password)
	valhash := hmac.New(sha256.New, []byte(login+password))
	hh := fmt.Sprintf("%x", valhash.Sum(nil))
	return hh, nil
}

func (d *Database) InitDB(ctx context.Context, DBURL string) error {
	DB.Ctx = ctx
	DB.DBURL = DBURL
	DB.Loger = logrus.New()
loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case <-time.After(2 * time.Second):
			var err error
			DB.Con, err = database.Connect(ctx, DBURL, d.Loger)
			if err != nil {
				DB.Loger.Error("Error conncet to DB: ", err)
				continue
			}
			break loop
		}
	}
	err := migrate.MigrateFromFS(ctx, DB.Con, &migrations.Migrations, DB.Loger)
	if err != nil {
		DB.Loger.Error("Error migration: ", err)
		return err
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (d *Database) CreateUser(u User) (cookie Cookie, err error) {
	// logrus.Info(u)
	// var err error
	hash, _ := HashPassword(u.Password)
	valhash := hmac.New(sha256.New, []byte(u.Login+hash))
	hh := fmt.Sprintf("%x", valhash.Sum(nil))
	switch d.Con {
	case nil:
		return cookie, err
	default:
		cookie := Cookie{
			Name:   "gofermart",
			Value:  hh,
			MaxAge: 86400,
		}
		_, err := d.Con.Exec(d.Ctx, "INSERT INTO users (login, password) VALUES($1,$2)", u.Login, hash)
		if err != nil {
			d.Loger.Error("Error create user: ", err)
			return cookie, err
		}
		return cookie, err
	}
}

func (d *Database) ReconnectDB() error {
	var err error
	for i := 0; i < 5; i++ {
		select {
		case <-d.Ctx.Done():
			return nil
		case <-time.After(2 * time.Second):
			DB.Con, err = database.Connect(d.Ctx, d.DBURL, d.Loger)
			if err != nil {
				d.Loger.Error("Error conncet to DB: ", err)
			} else {
				return nil
			}
		}
	}
	return err
}
