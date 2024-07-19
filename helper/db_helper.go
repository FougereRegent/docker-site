package helper

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DataBaseType int

const (
	LOCAL DataBaseType = iota
	POSTGRES_SQL
	MYSQL
)

var _db *gorm.DB = nil

type _sqliteConnection struct{}
type _postgresConnection struct{}
type _mysqlConnection struct{}

type IConnectionFabrique interface {
	CreateConnection(connectionString string) (*gorm.DB, error)
}

func (o *_sqliteConnection) CreateConnection(connectionString string) (*gorm.DB, error) {
	var err error
	_db, err = gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	return _db, err
}

func (o *_postgresConnection) CreateConnection(connectionString string) (*gorm.DB, error) {
	var err error
	_db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	return _db, err
}

func (o *_mysqlConnection) CreateConnection(connectionString string) (*gorm.DB, error) {
	var err error
	_db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	return _db, err
}

func CreateDatabase(connectionString string, databaseType DataBaseType) (*gorm.DB, error) {
	var fabrique IConnectionFabrique

	switch databaseType {
	case POSTGRES_SQL:
		fabrique = &_postgresConnection{}
		break
	case MYSQL:
		fabrique = &_mysqlConnection{}
		break
	case LOCAL:
		fabrique = &_sqliteConnection{}
		connectionString = "./docker-site.db"
		break
	}

	return fabrique.CreateConnection(connectionString)
}

func GetDb() (*gorm.DB, error) {
	if _db == nil {
		return nil, errors.New("Database doesn't exist")
	}
	return _db, nil
}
