package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"wm4z_back/config"
	"wm4z_back/server/apps/content"
	"wm4z_back/server/apps/services/about"
	"wm4z_back/server/apps/services/calendar"
	"wm4z_back/server/apps/services/tour"
)

type HTTPServer struct {
	config config.Config
	log    *zap.Logger
	engine *gin.Engine
	ctn    *content.Content
}

func (s *HTTPServer) Run() error {
	if err := s.engine.Run(); err != nil {
		return err
	}
	return nil
}

func (s *HTTPServer) Stop() {
	s.log.Sync()
}

func InitHTTPServer(config config.Config, logger *zap.Logger) Server {
	s := &HTTPServer{
		config: config,
		log:    logger,
		engine: gin.Default(),
		ctn:    content.InitContent(config, logger),
	}
	gin.SetMode(gin.ReleaseMode)
	s.regHandlers()
	return s
}

func (s *HTTPServer) regHandlers() {
	s.engine.Use(Cors())
	s.engine.GET("/about", about.AboutHandler(s.ctn))
	s.engine.GET("/tour", tour.TourHandler(s.ctn))
	s.engine.GET("/calendar", calendar.CalendarHandler(s.ctn))
}

//跨域

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers,Authorization,User-Agent, Keep-Alive, Content-Type, X-Requested-With,X-CSRF-Token,AccessToken,Token")
		c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.Header("Access-Control-Max-Age", "600")
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
