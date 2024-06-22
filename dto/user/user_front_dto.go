package user

import (
	"time"
)

type UserFrontDTO struct {
	Id             uint64
	Name           string
	LastConnection *time.Time
}
