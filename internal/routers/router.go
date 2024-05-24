package routers

import (
	"net/http"

	_ "github.com/camtrik/gin-blog/docs"
	"github.com/camtrik/gin-blog/global"
	"github.com/camtrik/gin-blog/internal/middleaware"
	"github.com/camtrik/gin-blog/internal/routers/api"
	v1 "github.com/camtrik/gin-blog/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleaware.Translation())
	// swagger blog
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	article := v1.NewArticle()
	tag := v1.NewTag()

	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	r.POST("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags", tag.List)
		apiv1.POST("/tags", tag.Create)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.DELETE("/tags/:id", tag.Delete)

		apiv1.GET("/articles", article.List)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.POST("/articles", article.Create)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.DELETE("/articles/:id", article.Delete)

		apiv1.POST("/articles/:id/tags", article.AddTag)
		apiv1.DELETE("/articles/:id/tags", article.DeleteTag)
	}

	return r
}
