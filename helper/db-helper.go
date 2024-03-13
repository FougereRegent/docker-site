package helper

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _db *gorm.DB = nil

func CreateDatabase(name string) (*gorm.DB, error) {
	var err error
	_db, err = gorm.Open(sqlite.Open(name), &gorm.Config{})
	return _db, err
}

func GetDb() (*gorm.DB, error) {
	if _db == nil {
		return nil, errors.New("Database doesn't exist")
	}
	return _db, nil
}
