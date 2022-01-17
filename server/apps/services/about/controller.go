package about

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"wm4z_back/config"
)

type AboutController struct {
	log 	*zap.Logger
	config config.Config
	db 		*gorm.DB
}

func InitAboutController(config config.Config) *AboutController{
	dsn := "zocp:Student@725@tcp(rm-uf60p6k023ue0dsmiio.mysql.rds.aliyuncs.com"
	db, err := gorm.Open("zocp:Student@725@tcp(rm-uf60p6k023ue0dsmiio.mysql.rds.aliyuncs.com)",&gorm.Config{})
	return &AboutController{
		config: config,
		db : ,
	}
}

func (a *AboutController) GetHandler() gin.HandlerFunc {
	return func(ctx *gin.Context){
		content, exist := ctx.Get("fromTo")
		if !exist{
			ctx.JSON(http.StatusBadRequest,ErrorResponse)
			return
		}
		from ,to := parseContent(content)


	}
}

func (a *AboutController) getAbouts(from int, to int) Abouts{

}

type AboutResponse struct {
	Code int `json:"code"`
	ErrorMsg string `json:"error"`
	Data interface{}
}

func ErrorResponse(err error) *AboutResponse {
	return &AboutResponse{Code: -1, ErrorMsg: err.Error()}
}

func SuccessResponse(data interface{}) *AboutResponse {
	return &AboutResponse{Code: 0, ErrorMsg: "ok", Data: data}
}

func parseContent(ctn interface{})(int, int){
	return 1,5
}
