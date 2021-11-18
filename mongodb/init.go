package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"remote_code/config"
)

var mgoCli *mongo.Client

func InitEngine() {
	var err error
	log.Println(config.Conf)
	log.Printf("mongodb://root:root@" + config.Conf.MongodbAddr)
	clientOptions := options.Client().ApplyURI("mongodb://root:root@" + config.Conf.MongodbAddr)

	// 连接到MongoDB
	mgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func GetMgoCli() *mongo.Client {
	if mgoCli == nil {
		InitEngine()
	}
	return mgoCli
}

func GetDb(database string) *mongo.Database {
	client = GetMgoCli()
	return client.Database(database)
}

func GetCollection(table string) *mongo.Collection {
	return GetDb(table).Collection(table) // todo 一个db 一个collection
}
