package albumdata

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// When writing an album to the DB, we allow mongod to add the _id field
type AlbumWritable struct {
	//album characteristics, including name, band, genre, and others
	Name      string    `bson:"Name" json:"Name"`
	Band      string    `bson:"Band" json:"Band"`
	Genre     string    `bson:"Genre" json:"Genre"`
	Year      int       `bson:"Year" json:"Year"`
	DateAdded time.Time `bson:"DateAdded"`
}

func (aw *AlbumWritable) SetDateAdded() {
	aw.DateAdded = time.Now()
}

// When reading, we recieve the _id field
// note that we can't just use composition of AlbumWritable, as we
// can't use the promoted fields in struct literals
type AlbumReadable struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	DateAdded time.Time          `bson:"DateAdded" json:"DateAdded"`
	Name      string             `bson:"Name" json:"Name"`
	Band      string             `bson:"Band" json:"Band"`
	Genre     string             `bson:"Genre" json:"Genre"`
	Year      int                `bson:"Year" json:"Year"`
}

func (f AlbumReadable) Copy() AlbumReadable {
	return AlbumReadable{f.ID, f.DateAdded, f.Name, f.Band, f.Genre, f.Year}
}
