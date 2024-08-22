package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-web-app/controller"
	"go-web-app/dao/mysql"
	"go-web-app/logger"
	"go-web-app/middlewares"
	"go-web-app/models"
	"net/http"
)

func Setup(mode, ClientUrl string, size int64, savedir string) *gin.Engine {
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(middlewares.Cors(ClientUrl))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.MaxMultipartMemory = size << 20
	//注册业务路由
	// 注册Swagger路由
	// 映射 Swagger UI 相关静态文件
	r.Static("/swagger-ui", "./docs/swagger-ui")
	r.StaticFile("/swagger.json", "./docs/swagger.json")

	url := ginSwagger.URL("/swagger.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.POST("/selectswitchchangvlan", controller.SelectSwitchChangeVlan)
	r.POST("/selectneighbors", controller.SelectNeighbors)
	r.POST("/selectinterfacedetail", controller.InterfaceDetail)

	r.POST("/selectswitchtotal", controller.SelectSwitchTotal)
	r.POST("/selectswitch", controller.SelectSwitchMac)
	r.POST("/download", controller.DownloadHandler)
	r.POST("/upload", func(ctx *gin.Context) {
		forms, err := ctx.MultipartForm()
		if err != nil {
			fmt.Println("error", err)
		}
		files := forms.File["file"]
		for _, v := range files {
			filelog := &models.Filelog{
				FileName: v.Filename,
				FileSize: v.Size,
				FileDir:  savedir + v.Filename,
			}
			fmt.Println(filelog)
			if err := ctx.SaveUploadedFile(v, fmt.Sprintf("%s%s", savedir, v.Filename)); err != nil {
				ctx.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
			}
			err = mysql.FileLogAdd(filelog)
		}
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
