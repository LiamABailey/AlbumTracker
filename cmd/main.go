package main

import (
	"github.com/LiamABailey/AlbumTracker/pkg/albumdata"
	"time"
)

func main() {
	// wait for mogno service to start
	time.Sleep(time.Second * 10)
	mc, _ := albumdata.NewMongoConnect()
	// we connect to the mongo server as part of the client-build step
	srv := albumdata.NewAPIServer(mc)
	// run the API indefinitely
	srv.Run(":8080")
}
