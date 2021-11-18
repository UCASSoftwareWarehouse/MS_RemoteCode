package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// add file
//https://stackoverflow.com/questions/64918108/how-to-persist-a-file-much-less-than-16mb-in-mongodb-using-official-go-driver

type Code struct {
	FileID     primitive.ObjectID   `json:"file_id" bson:"_id,omitempty"`
	FileName   string               `json:"file_name" bson:"file_name"`
	ProjectID  uint64               `json:"project_id" bson:"project_id"`
	FileType   int                  `json:"file_type" bson:"file_type"`
	FileSize   int64                `json:"file_size" bson:"file_size"`
	Content    []byte               `json:"content,omitempty" bson:"content,omitempty"`
	UpdateTime primitive.Timestamp  `json:"update_time" bson:"update_time"`
	ChildFiles []primitive.ObjectID `json:"child_files,omitempty" bson:"child_files,omitempty"`
}

// TableName get sql table name.获取数据库表名
func (m *Code) TableName() string {
	return "code"
}

// get codefile column name in mongodb.获取数据库列名
var CodeColumns = struct {
	FileID     string
	FileName   string
	ProjectID  string
	FileType   string
	FileSize   string
	Content    string
	UpdateTime string
	ChildFiles string
}{
	FileID:     "_id",
	FileName:   "file_name",
	ProjectID:  "project_id",
	FileType:   "file_type",
	FileSize:   "file_size",
	Content:    "content",
	UpdateTime: "update_time",
	ChildFiles: "child_files",
}

var (
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
)

func (m *Code) Insert(context context.Context, code *Code) (*primitive.ObjectID, error) {
	collection = GetCollection(m.TableName())
	res, err := collection.InsertOne(context, code)
	if err != nil {
		log.Fatalf("code insert err:%+v", err)
		return nil, err
	}
	id := res.InsertedID.(primitive.ObjectID)
	return &id, nil
}

// todo
func BuildCodeCondition(object Code) map[string]string {
	condition := make(map[string]string)
	if len(object.FileName) != 0 {
		condition["file_name"] = object.FileName
	}
	return condition
}

func (m *Code) FindOne(ctx context.Context, code *Code) (*Code, error) {
	collection = GetCollection(m.TableName())
	condition := BuildCodeCondition(*code)
	log.Printf("buildCondition:%+v", condition)

	res := collection.FindOne(ctx, condition)
	codeItem := &Code{}
	err := res.Decode(codeItem)
	if err != nil {
		log.Fatalf("codefile FindOne err:%+v", err)
		return nil, err
	}
	return codeItem, nil
}

func (m *Code) FindAll(ctx context.Context, code *Code) ([]Code, error) {
	var results []Code
	var cursor *mongo.Cursor
	var err error
	condition := BuildCodeCondition(*code)
	collection = GetCollection(m.TableName())
	if cursor, err = collection.Find(context.TODO(), condition, options.Find().SetSkip(0), options.Find().SetLimit(10)); err != nil {
		log.Fatalf("codefile FindAll err:%+v", err)
		return nil, err
	}
	//延迟关闭游标
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Fatalf("codefile FindAll err:%+v", err)
		}
	}()
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatalf("codefile FindAll err:%+v", err)
		return nil, err
	}
	for _, result := range results {
		log.Println(result)
	}
	return results, nil
}
