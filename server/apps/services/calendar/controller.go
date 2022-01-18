package calendar

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"wm4z_back/config"
)

type CalendarController struct {
	log    *zap.Logger
	config config.Config
	db     *gorm.DB
}
