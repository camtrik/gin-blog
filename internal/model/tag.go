package model

import "github.com/camtrik/gin-blog/pkg/app"

type Tag struct {
	*Model
	Name   string `json:"name"`
	Status uint8  `json:"status"`
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}
