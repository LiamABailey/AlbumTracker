package main

import (
	"github.com/LiamABailey/AlbumTracker/pkg/albumdata"
	"github.com/gin-gonic/contrib/static"
	"time"
)

func main() {
	// wait for mogno service to start
	time.Sleep(time.Second * 3)
	mc, _ := albumdata.NewMongoConnect()
	// we connect to the mongo server as part of the client-build step
	srv := albumdata.NewAPIServer(mc)
	// set the homepage
  srv.Use(static.Serve("/",static.LocalFile("../app", true)))
	// run the API indefinitely
	srv.Run(":8080")
}
