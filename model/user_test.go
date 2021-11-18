package model

import (
	"context"
	"log"
	"remote_code/config"
	"testing"
)

func TestUser_FindUserById(t *testing.T) {
	config.InitConfigDefault()
	InitGorm()
	user := User{}
	u, err := user.FindUserById(context.Background(), 1)
	log.Printf("id=%+v,err=%+v", u, err)
}
