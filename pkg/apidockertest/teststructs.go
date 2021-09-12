package apidockertest

import (
  "go.mongodb.org/mongo-driver/bson/primitive"
)

// when writing a fruit to the DB, we allow mongod to add the _id field
type FruitWritable struct {
  Name    string              `bson:"name"`
  Origin  string              `bson:"origin"`
  Price   int                 `bson:"number"`
}

// when reading, we recieve the _id field
type FruitReadable struct {
  ID      primitive.ObjectID  `bson:"_id"`
  Name    string              `bson:"name"`
  Origin  string              `bson:"origin"`
  Price   float64             `bson:"price"`
}


func (f FruitReadable) Copy() FruitReadable {
  return FruitReadable{f.ID, f.Name, f.Origin, f.Price}
}