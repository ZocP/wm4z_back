package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"wm4z_back/config"
	"wm4z_back/server"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	s := initDependencies()
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
	//defer s.Stop()
}

func initDependencies() server.Server {
	log, _ := zap.NewDevelopment()
	config := config.InitConfig()
	return server.InitHTTPServer(config, log)
}
