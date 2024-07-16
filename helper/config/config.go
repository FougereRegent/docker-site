package config

import (
	"errors"
	"os"
)

type DataBaseType int

const (
	LOCAL DataBaseType = iota
	POSTGRES_SQL
	MYSQL
)

const (
	sessionTimeEnv      string = "DOCKER_SITE_SESSION_TIME"
	dataBaseEnv         string = "DOCKER_SITE_DATABASE_TYPE"
	connectionStringEnv string = "DOCKER_SITE_CONNECTION_STRING"
)

type Conf struct {
	sessionTime      uint
	dataBase         DataBaseType
	connectionString string
}

func (o *Conf) SessionTime() uint {
	return o.sessionTime
}

func (o *Conf) DataBase() DataBaseType {
	return o.dataBase
}

func (o *Conf) ConnectionString() string {
	return o.connectionString
}

func ReadConfFromEnv() (*Conf, error) {
	sessionTime := os.Getenv(sessionTimeEnv)
	dataBase := os.Getenv(dataBaseEnv)
	connectionString := os.Getenv(connectionStringEnv)

	if sessionTime == "" || dataBase == "" || connectionString == "" {
		return nil, errors.New("Environment variable not set")
	}

	return &Conf{}, nil
}
