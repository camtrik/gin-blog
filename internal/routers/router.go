package routers

import (
	"log"
	"net/http"
	"time"

	_ "github.com/camtrik/gin-blog/docs"
	"github.com/camtrik/gin-blog/global"
	"github.com/camtrik/gin-blog/internal/middleware"
	"github.com/camtrik/gin-blog/internal/routers/api"
	v1 "github.com/camtrik/gin-blog/internal/routers/api/v1"
	"github.com/camtrik/gin-blog/pkg/limiter"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/net/context/ctxhttp"
)

var methodLimiter = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.Default()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.RateLimiter(methodLimiter))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Tracing())
	// r.Use(middleware.ContextTimeout(time.Second * 3))
	r.Use(middleware.Translation())
	// swagger blog
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	article := v1.NewArticle()
	tag := v1.NewTag()

	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	r.POST("/auth", api.GetAuth)
	r.GET("/auth", api.GetAuth)

	// test panic
	r.GET("/panic", func(c *gin.Context) {
		panic("test panic")
	})

	// test timeout
	r.GET("/timeout", func(c *gin.Context) {
		_, err := ctxhttp.Get(c.Request.Context(), http.DefaultClient, "https://httpbin.org/delay/10")
		if err != nil {
			log.Printf("ctxhttp.Get err: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.String(http.StatusOK, "Request successful")
	})

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
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
