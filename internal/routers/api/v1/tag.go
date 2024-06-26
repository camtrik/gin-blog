package v1

import (
	"github.com/camtrik/gin-blog/global"
	"github.com/camtrik/gin-blog/internal/service"
	"github.com/camtrik/gin-blog/pkg/app"
	convert "github.com/camtrik/gin-blog/pkg/covert"
	"github.com/camtrik/gin-blog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// @Summary Get multiple tags
// @Produce json
// @Param name query string false "Tag Name" maxlength(100)
// @Param state query int false "State" Enums(0, 1) default(1)
// @Param page query int false "Page Number"
// @Param page_size query int false "Page Size"
// @Success 200 {object} model.TagSwagger "success"
// @Failure 400 {object} errcode.Error "invalide params"
// @Failure 500 {object} errcode.Error "inside error"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf(c, "svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(tags, totalRows)
}

// @Summary Create a tag
// @Produce json
// @Param name body string true "Tag Name"
// @Param created_by body string true "Created By"
// @Param state body int false "State" Enums(0, 1) default(1)
// @Success 200 {object} model.TagSwagger "success"
// @Failure 400 {object} errcode.Error "invalide params"
// @Failure 500 {object} errcode.Error "inside error"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary Update a tag
// @Produce json
// @Param id path int true "Tag ID"
// @Param name body string false "Tag Name"
// @Param state body int false "State" Enums(0, 1)
// @Param modified_by body string true "Modified By"
// @Success 200 {object} model.TagSwagger "success"
// @Failure 400 {object} errcode.Error "invalide params"
// @Failure 500 {object} errcode.Error "inside error"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{Id: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary Delete a tag
// @Produce json
// @Param id path int true "Tag ID"
// @Success 200 {object} model.TagSwagger "success"
// @Failure 400 {object} errcode.Error "invalide params"
// @Failure 500 {object} errcode.Error "inside error"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{Id: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}
	response.ToResponse(gin.H{})
}
