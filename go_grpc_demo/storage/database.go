package storage

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "grpc.sqlite?cache=shared"
const ENV_PATH = "DB_PATH"

func newDatabase() (*sql.DB, error) {
	root := "./"
	if os.Getenv(ENV_PATH) != "" {
		root = os.Getenv(ENV_PATH)
		if !strings.HasSuffix(root, "/") {
			root = root + "/"
		}
	}
	dbUrl := "file:" + root + dbPath
	db, err := sql.Open("sqlite3", dbUrl)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitTables() (*sql.DB, error) {
	sqlStr := "CREATE TABLE IF NOT EXISTS ToDo(Id INTEGER PRIMARY KEY AUTOINCREMENT, Title TEXT, Description TEXT, Reminder INT)"
	db, err := newDatabase()
	if err != nil {
		return nil, fmt.Errorf("create database failed: " + err.Error())
	}
	_, err = db.Exec(sqlStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
