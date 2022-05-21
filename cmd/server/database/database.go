package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

func Connect(ctx context.Context, DBURL string) (conPool *pgxpool.Pool, err error) {
	conPool, err = pgxpool.Connect(ctx, DBURL)
	if err != nil {
		logrus.Error("Error pgx pool connect: ", err)
		return nil, err
	}
	logrus.Info("Connect to postrgress success!")
	return conPool, nil
}
