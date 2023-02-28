package data

import (
	"fmt"
)

type Driver int

// driver
const (
	SQLite Driver = iota
	MySQL
)

// Database storage database
type Database interface {
	Init()
	Close() error
}

// DBConfig config
type DBConfig interface {
	GetDBUrl() string
	GetDBName() string
}

// Open open a database client
func Open(conf DBConfig, dri Driver) (Database, error) {
	switch dri {
	case SQLite:
		database := SQLiteDB{}
		database.Driver = SQLite
		database.Config = conf
		database.Init()
		return &database, nil
	case MySQL:
		return nil, fmt.Errorf("not support mysql")
	}
	return nil, fmt.Errorf("Unknown database: %s", conf.GetDBName())
}
