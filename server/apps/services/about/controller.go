package about

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

type AboutController struct {
	log    *zap.Logger
	config config.Config
	db     *gorm.DB
}

func InitAboutController(config config.Config, log *zap.Logger) *AboutController {
	dsn := "zocp:Student@725@tcp(rm-uf60p6k023ue0dsmiio.mysql.rds.aliyuncs.com:3306)/wm4z"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("connecting to database: ", zap.Error(err))
	}
	return &AboutController{
		config: config,
		db:     db,
	}
}

func (a *AboutController) GetHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		f, exist1 := ctx.GetQuery("from")
		t, exist2 := ctx.GetQuery("to")
		if !exist1 || !exist2 {
			ctx.JSON(http.StatusBadRequest, apps.ErrorResponse(fmt.Errorf("invalid query")))
			return
		}

		from, to, valid := parseContent(f, t)
		if !valid {
			ctx.JSON(http.StatusBadRequest, apps.ErrorResponse(fmt.Errorf("invalid query parameters")))
		}

		ok, result := a.getAbouts(from, to)
		if !ok {
			ctx.JSON(http.StatusNotFound, apps.ErrorResponse(fmt.Errorf("didn't found match record")))
			return
		}

		ctx.JSON(http.StatusOK, apps.SuccessResponse(result))
	}
}

func (a *AboutController) getAbouts(from int, to int) (bool, []About) {
	var records []About
	a.db.Where("Numbers BETWEEN ? AND ?", from, to).Find(&records)
	if len(records) != 0 {
		return true, records
	}
	return false, nil
}

func parseContent(f string, t string) (int, int, bool) {
	from, err1 := strconv.Atoi(f)
	to, err2 := strconv.Atoi(t)
	if err1 != nil && err2 != nil {
		return -1, -1, false
	}
	if to-from < 0 || to-from > 50 {
		return -1, -1, false
	}
	return from, to, true
}
