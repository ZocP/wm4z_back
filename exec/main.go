package main

import (
	"go.uber.org/zap"
	"log"
	"wm4z_back/config"
	"wm4z_back/server"
)

func main() {
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
