package service

import (
	"github.com/satori/uuid"
)

func genToken(appId int32, username string) string {
	return uuid.NewV4().String()
}
