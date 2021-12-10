package model

import (
	"context"
	"github.com/jinzhu/gorm"
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

// get softmatadata by primary key projectid
func GetProjectById(db *gorm.DB, projectId uint64) (*Project, error) {
	project := new(Project)
	err := db.Where("ID = ?", projectId).First(project).Error
	if err != nil {
		log.Printf("get project by project id error, err=[%v]", err)
		return nil, err
	}
	log.Printf("get project info by id success!")
	return project, err
}
