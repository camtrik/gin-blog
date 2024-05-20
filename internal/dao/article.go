package dao

import (
	"github.com/camtrik/gin-blog/internal/model"
	"github.com/camtrik/gin-blog/pkg/app"
)

type Article struct {
	Id         uint32 `json:"id"`
	TagId      uint32 `json:"tag_id"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	State      uint8  `json:"state"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
}

func (d *Dao) GetArticle(id uint32, state uint8) (model.Article, error) {
	article := model.Article{Model: &model.Model{Id: id}, State: state}
	return article.Get(d.engine)
}

func (d *Dao) CountArticleByTag(tagId uint32, state uint8) (int64, error) {
	article := model.Article{State: state}
	return article.CountByTag(d.engine, tagId)
}

func (d *Dao) ListArticleByTag(tagId uint32, state uint8, page, pageSize int) ([]*model.ArticleInfo, error) {
	article := model.Article{State: state}
	return article.ListByTag(d.engine, tagId, app.GetPageOffset(page, pageSize), pageSize)
}

func (d *Dao) CreateArticle(param *Article) error {
	article := model.Article{
		Title:   param.Title,
		Desc:    param.Desc,
		Content: param.Content,
		State:   param.State,
		Model:   &model.Model{CreatedBy: param.CreatedBy},
	}
	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(param *Article) error {
	article := model.Article{Model: &model.Model{Id: param.Id}}
	values := map[string]interface{}{
		"modified_by": param.ModifiedBy,
		// TODO: check what this state does
		"state": param.State,
	}
	if param.Title != "" {
		values["title"] = param.Title
	}
	if param.Desc != "" {
		values["desc"] = param.Desc
	}
	if param.Content != "" {
		values["content"] = param.Content
	}
	return article.Update(d.engine, values)
}

func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{Id: id}}
	return article.Delete(d.engine)
}
