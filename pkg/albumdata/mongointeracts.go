package albumdata

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (M *MongoConnect) DeleteAlbumByID(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	// filter by the provided ID
	filt := bson.M{"_id": id}
	coll := M.getCollection()
	dr, err := coll.DeleteOne(context.TODO(), filt)
	return dr, err
}

func (M *MongoConnect) SearchAlbums(q AlbumQuery) ([]AlbumReadable, error) {
	//build the query as a filter via bson.M using various operators - $gt,  $or, $and
	var albums []AlbumReadable
	filt := bson.M{}
	if q.AlbumName != "" {
		// defaults to false
		if q.NameExactMatch {
			filt["Name"] = bson.M{"$eq": q.AlbumName}
		} else {
			// queried value is in name, case invariant
			filt["Name"] = bson.M{"$regex": q.AlbumName, "$options":"i"}
		}
	}

	if q.BandName != "" {
		// defaults to false
		if q.BandNameExactMatch {
			filt["Band"] = bson.M{"$eq": q.BandName}
		} else {
			// queried value is in name, case invariant
			filt["Band"] = bson.M{"$regex": q.BandName, "$options":"i"}
		}
	}

	//todo correct Genres to properly leverage list search
	if len(q.Genres) != 0 {
		filt["Genre"] = bson.M{"$in": q.Genres}
	}
	// support Year range via two parameters
	if q.YearStart != 0 {
		filt["Year"] = bson.M{"$gte":q.YearStart}
	}
	if q.YearEnd != 0 {
		if q.YearStart != 0 {
			filt["Year"] = bson.M{"$gte":q.YearStart,"$lte":q.YearEnd}
		} else {
			filt["Year"] = bson.M{"$lte":q.YearEnd}
		}

	}
	// support date added range via two parametersPer
	if !(q.DateAddedStart.IsZero()) {
		filt["DateAdded"] = bson.M{"$gte":q.DateAddedStart}
	}
	if !(q.DateAddedEnd.IsZero()) {
		if !(q.DateAddedStart.IsZero()) {
			filt["DateAdded"] = bson.M{"$lte":q.DateAddedEnd,"$gte":q.DateAddedStart}
		} else {
			filt["DateAdded"] = bson.M{"$lte":q.DateAddedEnd}
		}
	}

	coll := M.getCollection()
	resultcurs, err := coll.Find(context.TODO(), filt, fopt)
	if err != nil {
		return albums, err
	}
 	defer resultcurs.Close(context.TODO())
	// Decode each abum
	for resultcurs.Next(context.TODO()) {
		var album AlbumReadable
		if err = resultcurs.Decode(&album); err == nil {
			albums = append(albums, album.Copy())
		}
	}
	return albums, err

}

//TODO : mongo functions to support capabilities specified in albumtrackapi.go

// connector to collections
func (M *MongoConnect) getCollection() *mongo.Collection {
	return M.Client.Database(os.Getenv(DB)).Collection(os.Getenv(COLLECTION))
}
