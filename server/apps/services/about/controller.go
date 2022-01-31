package about

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"net/http"
	"strconv"
	"wm4z_back/server/apps"
	"wm4z_back/server/apps/content"
)

func AboutHandler(ctn *content.Content) gin.HandlerFunc {
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

		ok, result := getAbouts(ctn.Db, from, to)
		if !ok {
			ctx.JSON(http.StatusNotFound, apps.ErrorResponse(fmt.Errorf("didn't found match record")))
			return
		}

		ctx.JSON(http.StatusOK, apps.SuccessResponse(result))
	}
}

func getAbouts(db *gorm.DB, from int, to int) (bool, []About) {
	var records []About
	db.Where("Numbers BETWEEN ? AND ?", from, to).Find(&records)
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
