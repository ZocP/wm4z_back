package calendar

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"wm4z_back/config"
	"wm4z_back/server/apps"
)

type CalendarController struct {
	log    *zap.Logger
	config config.Config
	db     *gorm.DB
}

func (c CalendarController) GetHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		month, exist := ctx.GetQuery("month")
		if !exist {
			ctx.JSON(http.StatusBadRequest, apps.ErrorResponse(fmt.Errorf("invalid query")))
			return
		}

		valid, m := parseMonth(month)
		if !valid {
			ctx.JSON(http.StatusBadRequest, apps.ErrorResponse(fmt.Errorf("invalid param")))
			return
		}

		ok, result := c.getMonth(m)

		if !ok {
			ctx.JSON(http.StatusNotFound, apps.ErrorResponse(fmt.Errorf("didn't find any record")))
			return
		}

		ctx.JSON(http.StatusOK, apps.SuccessResponse(result))
	}
}

func InitCalendarController(config config.Config, log *zap.Logger) *CalendarController {
	dsn := "zocp:Student@725@tcp(rm-uf60p6k023ue0dsmiio.mysql.rds.aliyuncs.com:3306)/wm4z"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("connecting to database: ", zap.Error(err))
	}
	return &CalendarController{
		config: config,
		db:     db,
	}
}

func parseMonth(s string) (bool, int) {
	if r, err := strconv.Atoi(s); err == nil {
		return true, r
	}
	return false, -1
}

func (c *CalendarController) getMonth(month int) (bool, interface{}) {
	var record []Event
	c.db.Where("Month = ? ", month).Find(&record)
	if len(record) != 0 {
		return true, record
	}
	return false, nil
}
