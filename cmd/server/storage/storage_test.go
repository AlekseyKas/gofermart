package storage_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.gofermart/cmd/server/storage"
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

// func TestDatabase_GetUser(t *testing.T) {
// 	type fields struct {
// 		Con   *pgxpool.Pool
// 		Loger logrus.FieldLogger
// 		DBURL string
// 		Ctx   context.Context
// 	}
// 	type args struct {
// 		u User
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    string
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			d := &Database{
// 				Con:   tt.fields.Con,
// 				Loger: tt.fields.Loger,
// 				DBURL: tt.fields.DBURL,
// 				Ctx:   tt.fields.Ctx,
// 			}
// 			got, err := d.GetUser(tt.args.u)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Database.GetUser() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("Database.GetUser() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
