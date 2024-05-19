package dao

import "github.com/camtrik/gin-blog/internal/model"

func (d *Dao) CreateArticleTag(articleId, tagId uint32, createdBy string) error {
	articleTag := model.ArticleTag{
		Model:     &model.Model{CreatedBy: createdBy},
		ArticleId: articleId,
		TagId:     tagId,
	}
	return articleTag.Create(d.engine)
}
