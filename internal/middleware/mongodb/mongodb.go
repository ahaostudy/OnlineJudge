package mongodb

import (
	"main/internal/common/ctxt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo *mongo.Client

func InitMongoDB() error {
	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	// 连接mongodb
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	Mongo = client

	// ping
	if err = Mongo.Ping(ctx, nil); err != nil {
		return err
	}

	return nil
}

func Disconnect() {
	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	_ = Mongo.Disconnect(ctx)
}
