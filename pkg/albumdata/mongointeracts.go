package albumdata

import (
	"context"
	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

// define system environment variable names
const (
	CONNSTR    = "MONGOCONN"
	DB         = "MONGODB"
	COLLECTION = "MONGOCOLLECTION"
)

// Client-containing struct, supports method attachment for cleaner code.
type MongoConnect struct {
	Client *mongo.Client
}

// Create a new Mongo client
func NewMongoConnect() (*MongoConnect, error) {
	M := MongoConnect{}
	cliops := options.Client().ApplyURI(os.Getenv(CONNSTR))
	client, err := mongo.Connect(context.TODO(), cliops)
	M.Client = client
	return &M, err
}

// Disconnect the mongo client
func (M *MongoConnect) DisconnectMongoClient() error {
	err := M.Client.Disconnect(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func (M *MongoConnect) AddAlbum(a AlbumWritable) error {
	coll := M.getCollection()
	_, err := coll.InsertOne(context.TODO(), a)
	if err != nil {
		return err
	}
	return nil
}

//TODO : mongo functions to support capabilities specified in albumtrackapi.go

// connector to collections
func (M *MongoConnect) getCollection() *mongo.Collection {
	return M.Client.Database(os.Getenv(DB)).Collection(os.Getenv(COLLECTION))
}
