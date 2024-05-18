package v1

import (
	"github.com/camtrik/gin-blog/global"
	"github.com/camtrik/gin-blog/pkg/app"
	"github.com/camtrik/gin-blog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {}

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
	param := struct {
		Name  string `form:"name" binding:"max=100"`
		State uint8  `form:"state,default=1" binding:"oneof=0 1"`
	}{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	// parameters validation error
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	response.ToResponse(gin.H{})
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
func (t Tag) Create(c *gin.Context) {}
func (t Tag) Update(c *gin.Context) {}
func (t Tag) Delete(c *gin.Context) {}
