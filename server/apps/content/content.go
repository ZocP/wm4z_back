package content

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wm4z_back/config"
)

type Content struct {
	Log    *zap.Logger
	Config config.Config
	Db     *gorm.DB
}

func InitContent(config config.Config, log *zap.Logger) *Content {
	dsn := "zocp:Student@725@tcp(rm-uf60p6k023ue0dsmiio.mysql.rds.aliyuncs.com:3306)/wm4z"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("connecting to database: ", zap.Error(err))
	}
	return &Content{
		Config: config,
		Db:     db,
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
