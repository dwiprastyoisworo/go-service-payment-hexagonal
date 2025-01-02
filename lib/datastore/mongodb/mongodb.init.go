package mongodb

import (
	"fmt"
	"go-service-payment-hexagonal/lib/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongodbInit struct {
	cfg *config.Config
}

func NewMongodbInit(cfg *config.Config) *MongodbInit {
	return &MongodbInit{cfg: cfg}
}

func (m MongodbInit) DbInit() (*mongo.Database, error) {
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", m.cfg.MongoDB.Username, m.cfg.MongoDB.Password, m.cfg.MongoDB.Url, m.cfg.MongoDB.Password)
	client, err := mongo.Connect(options.Client().ApplyURI(url))
	if err != nil {
		panic("Mongodb connection error")
	}
	return client.Database(m.cfg.MongoDB.Database), nil
}
