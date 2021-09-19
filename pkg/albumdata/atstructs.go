package albumtrack

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// When writing an album to the DB, we allow mongod to add the _id field
type AlbumWritable struct {
	//album characteristics, including name, band, genre, and others
}

// When reading, we recieve the _id field
type AlbumReadable struct {
	ID     primitive.ObjectID `bson:"_id" json:"_id"`
	//same album characterstics
}

//
func (f AlbumReadable) Copy() AlbumReadable {
	return AlbumReadable{}//album fields}
}
