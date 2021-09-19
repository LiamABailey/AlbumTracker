package main

import (
	"github.com/LiamABailey/LearnDocker/pkg/albumtrack"
	"time"
)

func main() {
	// wait for mogno service to start
	time.Sleep(time.Second * 10)
	mc, _ := apidockertest.NewMongoConnect()
	// we connect to the mongo server as part of the client-build step
	srv := apidockertest.NewAPIServer(mc)
	// run the API indefinitely
	srv.Run(":8080")
}
