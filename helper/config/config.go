package config

import (
	. "docker-site/helper"
	"errors"
	"os"
	"strconv"
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

	sessionTimeConverted, err := strconv.ParseUint(sessionTime, 10, 64)
	if err != nil {
		return nil, err
	}

	var dataBaseType DataBaseType
	switch dataBase {
	case "LOCAL":
		dataBaseType = LOCAL
		break
	case "POSTGRES_SQL":
		dataBaseType = POSTGRES_SQL
		break
	case "MYSQL":
		dataBaseType = MYSQL
		break
	default:
		return nil, errors.New("")
	}

	result := &Conf{
		sessionTime:      uint(sessionTimeConverted),
		dataBase:         dataBaseType,
		connectionString: connectionString,
	}

	return result, nil
}
