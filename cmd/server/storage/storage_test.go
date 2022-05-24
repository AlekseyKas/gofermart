package storage_test

import (
	"context"
	"testing"

	"go.gofermart/cmd/server/storage"
	"go.gofermart/internal/helpers"
)

func TestDatabase_InitDB(t *testing.T) {

	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, _ := context.WithCancel(context.Background())

			DBURL, id, _ := helpers.StartDB()

			storage.IDB = &storage.DB
			storage.IDB.InitDB(ctx, DBURL)
			defer helpers.StopDB(id)
			if err := storage.IDB.InitDB(ctx, DBURL); (err != nil) != tt.wantErr {
				t.Errorf("Database.InitDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
