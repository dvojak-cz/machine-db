package repository

import (
	"context"
	"strconv"
	"sync"

	"git.recombee.net/jtrojak/server-db/config"
	"github.com/docker/cli/cli/registry/client"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository2 struct {
	MongoRepository mongoRepository
}

type Repository interface {
	Connect() error
	Disconnect() error
}

type mongoRepository struct {
	Client *client.Client
	connectOnce   sync.Once
	connectionFailed bool
}

func (r *repository2) Init() {
	log.Info("Repository initialized")
	
}

func (m *mongoRepository) Connect() error {
	log.Info("Connecting to MongoDB")

	opt := options.
		Client().
		SetAuth(options.Credential{Username: config.Configs.MongoDbConfig.Username, Password: config.Configs.MongoDbConfig.Password}).
		ApplyURI("mongodb://" + config.Configs.MongoDbConfig.Host + ":" + strconv.Itoa(config.Configs.MongoDbConfig.Port))

	if err := opt.Validate(); err != nil {
		log.WithError(err).Fatal("Failed to validate MongoDB connection options")
		return err
	}

	client, err := mongo.Connect(context.TODO(), opt)
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to MongoDB")
		return err
	}
	m.client = client

	log.Info("Connected to MongoDB")
	return nil
}

func (m *mongoRepository) Disconnect() error {
	log.Info("Closing repository")
	if err := m.client.Disconnect(context.Background()); err != nil {
		log.WithError(err).Error("Failed to disconnect from MongoDB")
		return err
	}
	return nil
}
