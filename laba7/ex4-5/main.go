package main

import (
	"context"
	"os"
	"os/signal"
	"psutiGoLabs/laba7/ex4-5/group"
	"psutiGoLabs/laba7/ex4-5/server"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	h := group.NewHandler(log)
	g := gin.New()
	h.InitRoutes(g)
	server := server.NewServer(g)

	go func() {
		server.Run()
	}()

	log.Info("server is running")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	server.Stop(context.Background())
	log.Info("server shutdown")
}
