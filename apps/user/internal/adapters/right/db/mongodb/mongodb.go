package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/adapters/right/db"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoApi struct {
	url        string
	ctx        context.Context
	connection *mongo.Client
	l          logger.CustomLogger
}

func NewMongoApi(ctx context.Context, dbPort int, dbIp, dbUserName, dbPassword, collectionName string, l logger.CustomLogger) db.DbPort {
	ctx = context.WithValue(ctx, "database", collectionName)
	url := fmt.Sprintf("mongodb://%s:%s@%s:%d", dbUserName, dbPassword, dbIp, dbPort)
	return &mongoApi{
		url: url,
		l:   l,
		ctx: ctx,
	}
}

// StartDbConnection implements ports.DbPort.
func (mApi *mongoApi) StartDbConnection() error {
	clientOptions := options.Client().ApplyURI(mApi.url)
	client, err := mongo.Connect(mApi.ctx, clientOptions)
	if err != nil {
		return err
	}
	mApi.connection = client

	// Check the connection
	ctx, cancel := context.WithTimeout(mApi.ctx, time.Second*5)
	err = client.Ping(ctx, nil)
	defer cancel()
	if err != nil {
		return err
	}

	log.Println("Database Connected.")
	return nil
}

func (mApi *mongoApi) CloseDbConnection() error {
	err := mApi.connection.Disconnect(mApi.ctx)
	if err != nil {
		return err
	}
	return nil
}
