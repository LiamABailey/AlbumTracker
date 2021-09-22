package albumdata

import (
	"github.com/gin-gonic/gin"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type APIServer struct {
	connector *MongoConnect
	router    *gin.Engine
}

func NewAPIServer(mc *MongoConnect) *APIServer {
	svr := &APIServer{connector: mc}
	svr.router = gin.Default()
	svr.router.POST("/albums",svr.addAlbum)

	return svr
}

func (srv *APIServer) Run(address string) error {
  return srv.router.Run(address)
}

func (srv *APIServer) addAlbum(ctx * gin.Context) {
	var album AlbumWritable
	// bind the json body into the AlbumWritable
	if err := ctx.ShouldBindJSON(&album); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// set the date created
	album.SetDateAdded()
	/// attempt to write to Mongo
	if err := srv.connector.AddAlbum(album); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// send a 201 status on success
	ctx.JSON(http.StatusCreated, nil)
}

//TODO : Need functions for the following
//GET Album
//GET Albums conforming to query
//POST new album
//Potenitailly update an album
//DELETE album

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

//create the mongo ObjectID primitive from the hex-id string
//func idFromContext(ctx *gin.Context) (primitive.ObjectID, error) {
//	return primitive.ObjectIDFromHex(ctx.Param("id"))
//}
