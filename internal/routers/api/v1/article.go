package v1

import (
	"github.com/camtrik/gin-blog/global"
	"github.com/camtrik/gin-blog/internal/service"
	"github.com/camtrik/gin-blog/pkg/app"
	convert "github.com/camtrik/gin-blog/pkg/covert"
	"github.com/camtrik/gin-blog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	param := service.GetArticleRequest{Id: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	article, err := svc.GetArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleFail)
		return
	}

	response.ToResponse(article)
}

func (a Article) List(c *gin.Context) {
	param := service.ArticleListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	articlesWithTag, articleCount, err := svc.GetArticleList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetArticleList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleListFail)
		return
	}

	response.ToResponseList(articlesWithTag, articleCount)
}

func (a Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}

	response.ToResponse(gin.H{})
}
func (a Article) Update(c *gin.Context) {
	param := service.UpdateArticleRequest{Id: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}

	response.ToResponse(gin.H{})
}

func (a Article) Delete(c *gin.Context) {
	param := service.DeleteArticleRequest{Id: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}

	response.ToResponse(gin.H{})
}

func (a Article) AddTag(c *gin.Context) {
	param := service.AddArticleTagRequest{ArticleId: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.AddArticleTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.AddArticleTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorAddArticleTagFail)
		return
	}

	response.ToResponse(gin.H{})
}

func (a Article) DeleteTag(c *gin.Context) {
	param := service.DeleteArticleTagRequest{ArticleId: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteArticleTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteArticleTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleTagFail)
		return
	}

	response.ToResponse(gin.H{})
}
