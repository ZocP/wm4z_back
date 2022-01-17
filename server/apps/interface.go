package apps

import "github.com/gin-gonic/gin"

type AppController interface {
	GetHandler() gin.HandlerFunc
}
