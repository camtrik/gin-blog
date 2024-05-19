package service

import (
	"github.com/camtrik/gin-blog/internal/model"
	"github.com/camtrik/gin-blog/pkg/app"
)

// article information with tag
type Article struct {
	Id      uint32     `json:"id"`
	Title   string     `json:"title"`
	Desc    string     `json:"desc"`
	Content string     `json:"content"`
	State   uint8      `json:"state"`
	Tag     *model.Tag `json:"tag"`
}

type ArticleListRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	TagId     uint32 `form:"tag_id" binding:"required,gte=1"`
	Title     string `form:"title" binding:"required,min=3,max=100"`
	Desc      string `form:"desc" binding:"required,min=3,max=255"`
	Content   string `form:"content" binding:"required,min=3"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Title      string `form:"title" binding:"min=3,max=100"`
	Desc       string `form:"desc" binding:"min=3,max=255"`
	Content    string `form:"content" binding:"min=3"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
	State      uint8  `form:"state" binding:"required,oneof=0 1"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

// func (svc *Service) GetArticleList(param *ArticleListRequest, pager *app.Pager) ([]*model.Article, error) {
	
// }
