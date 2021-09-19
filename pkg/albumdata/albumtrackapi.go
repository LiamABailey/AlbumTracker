package albumtrack

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
	//route GET, POST, and DELETE methods

	return svr
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
func idFromContext(ctx *gin.Context) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(ctx.Param("id"))
}
