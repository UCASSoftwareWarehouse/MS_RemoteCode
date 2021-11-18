package model

import (
	"context"
	"log"
)

type User struct {
	Id       int64  `xorm:"pk autoincr BIGINT"`
	UserName string `xorm:"not null VARCHAR(100)"`
	Password string `xorm:"not null CHAR(32)"`
}

func (user *User) TableName() string {
	return "user"
}

func (u *User) FindUserById(ctx context.Context, id int) (*User, error) {
	db := GetDb().Model(User{})
	user := &User{}
	err := db.Where("id=?", id).Find(user).Error
	if err != nil {
		log.Fatalf("ser findUserById err:%+v", err)
		return nil, err
	}
	return user, nil
}
