package storage_test

import (
	"database/sql"
	"go_grpc/storage"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestInitTables(t *testing.T) {
	tests := []struct {
		name    string
		want    *sql.DB
		wantErr bool
	}{
		{"create 1", nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := storage.InitTables()
			if (err != nil) != tt.wantErr {
				t.Errorf("InitTables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitTables() = %v, want %v", got, tt.want)
			}
		})
	}
}
