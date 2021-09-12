package apidockertest

import (
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "context"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)


const (
  CONNSTR = "mongodb://localhost:27017"
  DB      = "food"
  FRUIT   = "fruits"
)

type MongoConnect struct {
  Client    *mongo.Client
}


func NewMongoClient() (*MongoConnect, error) {
  M := MongoConnect{}
  cliops := options.Client().ApplyURI(CONNSTR)
  client, err := mongo.Connect(context.TODO(), cliops)
  M.Client = client
  return &M, err
}

func (M *MongoConnect) DisconnectMongoClient() error {
  err := M.Client.Disconnect(context.TODO())
  if err != nil {
    return err
  }
  return nil
}

// adds a new Fruit to the db
func (M *MongoConnect) AddFruit(newfruit FruitWritable) error {
  coll := M.Client.Database(DB).Collection(FRUIT)
  _, err := coll.InsertOne(context.TODO(), newfruit)
  if err != nil {
    return err
  }
  return nil
}

// adds a slice of fruits to the database
func (M *MongoConnect) AddFruitMulti(newfruits []FruitWritable) error {
    //InsertMany takes an []interface{}
    insertables := make([]interface{}, len(newfruits))
    for i,v := range newfruits {
      insertables[i] = v
    }

    coll := M.Client.Database(DB).Collection(FRUIT)
    _, err := coll.InsertMany(context.TODO(), insertables)
    if err != nil {
      return err
    }
    return nil
}

// Returns fruit matching the provided ID.
func (M *MongoConnect) GetFruitByID(id primitive.ObjectID ) (FruitReadable, error) {
  var fruit FruitReadable
  filt := bson.D{primitive.E{Key: "_id", Value: id}}
  coll := M.Client.Database(DB).Collection(FRUIT)
  err := coll.FindOne(context.TODO(), filt).Decode(&fruit)
  return fruit, err
}

// Returns any fruit matching the provided name
func (M *MongoConnect) GetFruitByName(fruitname string) ([]FruitReadable, error) {
  var fruits []FruitReadable
  filt := bson.M{"name": "orange"}
  //bson.D{primitive.E{Key: "name", Value: fruitname}}
  coll := M.Client.Database(DB).Collection(FRUIT)
  curs, err := coll.Find(context.TODO(), filt)
  if err != nil {
    return fruits, err
  }
  defer curs.Close(context.TODO())
  for curs.Next(context.TODO()){
    var fruit FruitReadable
    if err = curs.Decode(&fruit); err == nil {
      fruits = append(fruits, fruit.Copy())
    }
  }
  return fruits, err
}

// Remove fruit by ID
//func (M *MongoConnect) RemoveFruitByID(id primitive.ObjectID) (error){
//
//}