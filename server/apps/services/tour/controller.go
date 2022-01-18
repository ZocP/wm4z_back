package tour

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wm4z_back/config"
)

type TourController struct {
	log    *zap.Logger
	config config.Config
	db     *gorm.DB
}

func (t *TourController) GetHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func InitTourController(config config.Config, log *zap.Logger) *TourController {

	dsn := "zocp:Student@725@tcp(rm-uf60p6k023ue0dsmiio.mysql.rds.aliyuncs.com:3306)/wm4z"
	//dsn := getDSN(config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("connecting to database: ", zap.Error(err))
	}
	return &TourController{
		config: config,
		db:     db,
	}
}

func getDSN(config config.Config) string {
	un := config.Services.About.DB.UserName
	pc := config.Services.About.DB.Password
	prtc := config.Services.About.DB.Protocol
	url := config.Services.About.DB.URL
	dn := config.Services.About.DB.DBName
	return un + ":" + pc + "@" + prtc + "(" + url + ")/" + dn
}
