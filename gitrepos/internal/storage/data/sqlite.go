package data

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// SQLiteDB sqlite database
type SQLiteDB struct {
	client *gorm.DB
	Driver Driver
	Config DBConfig
}

// Init init database client
func (s *SQLiteDB) Init() {
	s.client = nil
}

// Close close database client
func (s *SQLiteDB) Close() error {
	sqlDB, err := s.client.DB()
	if err != nil {
		return errors.WithStack(err)
	}
	err = sqlDB.Close()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
