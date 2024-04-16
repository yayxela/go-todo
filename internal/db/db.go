package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/yayxela/go-todo/internal/config"
)

// Список коллекций бд
const (
	TaskCollection = "tasks"
)

type IDB interface {
	GetDB() *mongo.Database
	Disconnect(ctx context.Context) error
}

type db struct {
	*mongo.Client
	dbName string
}

func New(ctx context.Context, config config.DBConfig) (IDB, error) {
	co := options.Client().ApplyURI(config.GetConnectionString())
	client, err := mongo.Connect(ctx, co)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &db{
		Client: client,
		dbName: config.DBName,
	}, nil
}

// GetDB ...
// Получение подключения к бд
func (d *db) GetDB() *mongo.Database {
	return d.Database(d.dbName)
}

func (d *db) Disconnect(ctx context.Context) error {
	return d.Client.Disconnect(ctx)
}

// GetOID ...
// Получение primitive.ObjectID из строки
func GetOID(id string) primitive.ObjectID {
	oid, _ := primitive.ObjectIDFromHex(id)
	return oid
}
