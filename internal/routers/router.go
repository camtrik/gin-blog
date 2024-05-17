package routers

import (
	_ "github.com/camtrik/gin-blog/docs"
	v1 "github.com/camtrik/gin-blog/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// swagger blog
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	article := v1.NewArticle()
	tag := v1.NewTag()

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
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.DELETE("/articles/:id", article.Delete)
	}

	return r
}
