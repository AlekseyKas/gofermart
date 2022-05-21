package storage

import (
	"context"
	"time"

	"go.gofermart/cmd/server/database"
	"go.gofermart/cmd/server/storage/migrations"
	"go.gofermart/internal/migrate"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type Database struct {
	con   *pgxpool.Pool
	loger logrus.FieldLogger
}

var DB Database
var IDB Storage

type Storage interface {
	InitDB(ctx context.Context, DBURL string) error
}

func (d *Database) InitDB(ctx context.Context, DBURL string) error {
loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case <-time.After(2 * time.Second):
			var err error
			DB.con, err = database.Connect(ctx, DBURL)
			if err != nil {
				logrus.Error("Error conncet to DB: ", err)
				continue
			}
			break loop
		}
	}
	err := migrate.MigrateFromFS(ctx, DB.con, &migrations.Migrations, DB.loger)
	if err != nil {
		logrus.Error("Error migration: ", err)
		return err
	}
	logrus.Info("end", d.con)
	return nil
}
