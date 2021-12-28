package v1

import (
	model "gochenweb/models"
	"gochenweb/pkg/e"
	"gochenweb/pkg/setting"
	"gochenweb/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state, _ = com.StrTo(arg).Int()
		maps["state"] = state
	}
	code := e.SUCCESS
	data["lists"] = model.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = model.GetTotalTags(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddTags(c *gin.Context) {

}

func UpdateTags(c *gin.Context) {

}
