package calendar

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"wm4z_back/server/apps"
	"wm4z_back/server/apps/content"
)

func CalendarHandler(content *content.Content) gin.HandlerFunc {
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

		ok, result := getMonth(content.Db, m)

		if !ok {
			ctx.JSON(http.StatusNotFound, apps.ErrorResponse(fmt.Errorf("didn't find any record")))
			return
		}

		ctx.JSON(http.StatusOK, apps.SuccessResponse(result))
	}
}

func parseMonth(s string) (bool, int) {
	if r, err := strconv.Atoi(s); err == nil {
		return true, r
	}
	return false, -1
}

func getMonth(db *gorm.DB, month int) (bool, interface{}) {
	var record []Event
	db.Where("Month = ? ", month).Find(&record)
	if len(record) != 0 {
		return true, record
	}
	return false, nil
}
