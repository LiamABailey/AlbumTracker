package main

import (
	"github.com/LiamABailey/AlbumTracker/pkg/albumdata"
	"github.com/LiamABailey/AlbumTracker/pkg/spotifyintegration"
	"github.com/gin-gonic/contrib/static"
	"time"
)

func main() {
	// wait for mongo service to start
	time.Sleep(time.Second * 1)
	mc, _ := albumdata.NewMongoConnect()
	// we connect to the mongo server as part of the client-build step
	srv := albumdata.NewAPIServer(mc)
	// prep the spotify service
	spotifysrv := spotifyintegration.NewSpotifyServer()
	// set the homepage
  srv.Use(static.Serve("/",static.LocalFile("../app", true)))
	// run the API indefinitely
	wait := make(chan bool)
	go func() {
		srv.Run(":8080")
		wait <- true
	}()
	go func() {
		spotifysrv.Run(":8081")
		wait <- true
	}()
	<- wait
}
