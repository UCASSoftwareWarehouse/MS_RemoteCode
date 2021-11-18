package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"remote_code/config"
	"testing"
)

func TestCode_Insert(t *testing.T) {
	config.InitConfigDefault()
	InitEngine()
	code := Code{}
	data := Code{
		FileName:   "numpy",
		ProjectID:  1,
		FileType:   0,
		FileSize:   0,
		Content:    nil,
		UpdateTime: primitive.Timestamp{},
		ChildFiles: nil,
	}
	res, err := code.Insert(context.Background(), &data)
	log.Printf("res=%+v,err=%+v", res, err)
}

func TestCode_FindOne(t *testing.T) {
	config.InitConfigDefault()
	InitEngine()
	code := Code{}
	data := Code{
		FileName: "numpy",
	}
	res, err := code.FindOne(context.Background(), &data)
	log.Printf("res=%+v,err=%+v", res, err)
}

func TestCode_FindAll(t *testing.T) {
	config.InitConfigDefault()
	InitEngine()
	code := Code{}
	data := Code{
		//FileName:   "numpy",
	}
	res, err := code.FindAll(context.Background(), &data)
	log.Printf("res=%+v,err=%+v", res, err)
}
