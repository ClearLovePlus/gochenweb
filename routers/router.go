package routers

import (
	"gochenweb/pkg/setting"
	v1 "gochenweb/routers/api"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	apiv1s := r.Group("api/v1")
	{
		//获取信息
		apiv1s.GET("/tags", v1.GetTags)
		apiv1s.POST("/tags/add", v1.AddTags)
		apiv1s.POST("/tags/update", v1.UpdateTags)
	}
	return r
}
