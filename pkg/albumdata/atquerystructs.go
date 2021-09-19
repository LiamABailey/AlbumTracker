package albumtrack

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


// master query struct
type AlbumQuery struct {
  // fields defining the search
  // name LIKE
  // genre (in list)
  // date criteria (added before date, added on or after date)
  // band name (in list)
  // return up to X results or pagination
}
