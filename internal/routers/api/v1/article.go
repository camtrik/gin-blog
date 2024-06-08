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

// @Summary Get an article by id
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {object} model.ArticleSwagger "success"
// @Failure 400 {object} errcode.Error "invalide params"
// @Failure 500 {object} errcode.Error "inside error"
// @Router /api/v1/articles/{id} [get]
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

// @Summary Get a list of articles
// @Produce json
// @Param tag_id query int false "Tag ID"
// @Param title query string false "Title"
// @Param state query int false "State" Enums(0, 1)
// @Param page query int false "Page"
// @Param page_size query int false "Page Size"
// @Success 200 {object} model.ArticleSwagger "success"
// @Failure 400 {object} errcode.Error "invalide params"
// @Failure 500 {object} errcode.Error "inside error"
// @Router /api/v1/articles [get]
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

// @Summary Create an article
// @Produce json
// @Param title body string true "Title"
// @Param desc body string true "Desc"
// @Param content body string true "Content"
// @Param cover_image_url body string true "Cover Image URL"
// @Param state body int false "State" Enums(0, 1) default(1)
// @Param created_by body string true "Created By"
// @Success 200 {object} model.ArticleSwagger "success"
// @Failure 400 {object} errcode.Error "invalide params"
// @Failure 500 {object} errcode.Error "inside error"
// @Router /api/v1/articles [post]
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

// @Summary Update an article
// @Produce json
// @Param id path int true "Article ID"
// @Param title body string false "Title"
// @Param desc body string false "Desc"
// @Param content body string false "Content"
// @Param cover_image_url body string false "Cover Image URL"
// @Param state body int false "State" Enums(0, 1)
// @Param modified_by body string true "Modified By"
// @Success 200 {object} model.ArticleSwagger "success"
// @Failure 400 {object} errcode.Error "invalide params"
// @Failure 500 {object} errcode.Error "inside error"
// @Router /api/v1/articles/{id} [put]
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

// @Summary Delete an article
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {object} model.ArticleSwagger "success"
// @Failure 400 {object} errcode.Error "invalide params"
// @Failure 500 {object} errcode.Error "inside error"
// @Router /api/v1/articles/{id} [delete]
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

// @Summary Add a tag to an article
// @Produce json
// @Param id path int true "Article ID"
// @Param tag_id body int true "Tag ID"
// @Param created_by body string true "Created By"
// @Success 200 {object} model.ArticleSwagger "success"
// @Failure 400 {object} errcode.Error "invalide params"
// @Failure 500 {object} errcode.Error "inside error"
// @Router /api/v1/articles/{article_id}/tags [post]
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

// @Summary Delete a tag from an article
// @Produce json
// @Param id path int true "Article ID"
// @Param tag_id body int true "Tag ID"
// @Success 200 {object} model.ArticleSwagger "success"
// @Failure 400 {object} errcode.Error "invalide params"
// @Failure 500 {object} errcode.Error "inside error"
// @Router /api/v1/articles/{article_id}/tags [delete]
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
