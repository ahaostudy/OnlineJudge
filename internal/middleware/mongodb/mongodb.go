package mongodb

import (
	"fmt"
	"main/config"
	"main/internal/common/ctxt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database
var MongoContesScore *mongo.Collection

func InitMongoDB() error {
	conf := config.ConfMongodb

	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	// 连接mongodb
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", conf.Username, conf.Password, conf.Host, conf.Port)))
	if err != nil {
		return err
	}

	// ping
	if err = client.Ping(ctx, nil); err != nil {
		return err
	}

	MongoClient = client
	MongoDB = client.Database(conf.Dbname)
	MongoContesScore = MongoDB.Collection("contest_score")

	return nil
}

func Disconnect() {
	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	_ = MongoClient.Disconnect(ctx)
}
