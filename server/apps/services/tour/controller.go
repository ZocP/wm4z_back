package tour

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"wm4z_back/server/apps"
	"wm4z_back/server/apps/content"
)

func TourHandler(ctn *content.Content) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		floor, exist := ctx.GetQuery("floor")
		fl, ok := parseFloor(floor)
		if !exist || !ok {
			ctx.JSON(http.StatusBadRequest, apps.ErrorResponse(fmt.Errorf("invalid query")))
			return
		}

		result, ok := search(ctn.Db, fl)
		if !ok {
			ctx.JSON(http.StatusOK, apps.ErrorResponse(fmt.Errorf("didn't found records match the requirement")))
			return
		}

		ctx.JSON(http.StatusOK, apps.SuccessResponse(result))
	}
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

func search(db *gorm.DB, floor int) (*Floor, bool) {

	var records []Position
	db.Where("FloorNumber = ?", floor).Find(&records)
	if len(records) != 0 {
		return &Floor{
			FloorNumber: floor,
			Positions:   records,
		}, true
	}
	return nil, false
}

func searchFloor(db *gorm.DB, f int) (int, bool) {
	var records []Position
	db.Where("FloorNumber = ?", f).Find(&records)
	if len(records) != 0 {
		return len(records), true
	}
	return -1, false
}

func parseFloor(floor string) (int, bool) {
	f, err := strconv.Atoi(floor)
	if f > 2 {
		return -1, false
	}
	if err != nil {
		return -1, false
	}
	return f, true
}
