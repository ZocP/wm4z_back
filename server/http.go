package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"wm4z_back/config"
	"wm4z_back/server/apps"
	"wm4z_back/server/apps/services/about"
)

type HTTPServer struct {
	appsController map[string]apps.AppController
	config         config.Config
	log            *zap.Logger
	engine         *gin.Engine
}

func (s *HTTPServer) Run() error {
	if err := s.engine.Run(); err != nil {
		return err
	}
	return nil
}

func (s *HTTPServer) Stop() {
	panic("implement me")
}

func InitHTTPServer(config config.Config, logger *zap.Logger) Server {
	s := &HTTPServer{
		config: config,
		log:    logger,
		engine: gin.Default(),
	}
	s.regControllers(config)
	s.regHandlers()
	return s
}

func (s *HTTPServer) regControllers(config config.Config) {
	s.appsController = make(map[string]apps.AppController)
	s.appsController["about"] = about.InitAboutController(config)
}

func (s *HTTPServer) regHandlers() {
	s.engine.GET("/about", s.appsController["about"].GetHandler())
}
