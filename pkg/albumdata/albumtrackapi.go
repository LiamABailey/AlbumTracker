package albumdata

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	svr.router.GET("/albums/search",svr.searchAlbums)
	svr.router.DELETE("/albums",svr.deleteAlbumByID)
	return svr
}

func (svr *APIServer) Run(address string) error {
  return svr.router.Run(address)
}

func (svr *APIServer) Use(hf gin.HandlerFunc) {
	svr.router.Use(hf)
}

func (svr *APIServer) addAlbum(ctx *gin.Context) {
	var album AlbumWritable
	// bind the json body into the AlbumWritable
	if err := ctx.ShouldBindJSON(&album); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// set the date created
	album.SetDateAdded()
	/// attempt to write to Mongo
	if err := svr.connector.AddAlbum(album); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// send a 201 status on success
	ctx.JSON(http.StatusCreated, "Album Posted")
}

func (svr *APIServer) deleteAlbumByID(ctx *gin.Context) {
	// get the "id" from the context
	id, _ := idFromContext(ctx)
	// attempt to delete from Mongo
	dr, err := svr.connector.DeleteAlbumByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, dr)
}

func (svr *APIServer) searchAlbums(ctx *gin.Context) {
	var aquery AlbumQuery
	if err := ctx.Bind(&aquery); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// attempt to perform the search
	queryres, err := svr.connector.SearchAlbums(aquery)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.IndentedJSON(http.StatusOK, queryres)

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

//create the mongo ObjectID primitive from the hex-id string
func idFromContext(ctx *gin.Context) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(ctx.Param("id"))
}
