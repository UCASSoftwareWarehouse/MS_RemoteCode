package model

import (
	"context"
	"log"
	"time"
)

type Project struct {
	Id                 int64     `xorm:"pk autoincr BIGINT"`
	ProjectName        string    `xorm:"not null VARCHAR(100)"`
	UserId             int64     `xorm:"not null index BIGINT"`
	Tags               string    `xorm:"VARCHAR(20)"`
	CodeAddr           string    `xorm:"CHAR(12)"`
	BinaryAddr         string    `xorm:"CHAR(12)"`
	License            string    `xorm:"VARCHAR(50)"`
	UpdateTime         time.Time `xorm:"not null TIMESTAMP"`
	ProjectDescription string    `xorm:"TEXT"`
}

func (p *Project) TableName() string {
	return "user"
}

func (p *Project) FindProjectById(ctx context.Context, id int) (*Project, error) {
	db := GetDb().Model(Project{})
	project := &Project{}
	err := db.Where("id=?", id).Find(project).Error
	if err != nil {
		log.Printf("project FindProjectById err:%+v", err)
		return nil, err
	}
	return project, nil
}
