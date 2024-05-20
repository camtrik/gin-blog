package dao

import "github.com/camtrik/gin-blog/internal/model"

func (d *Dao) GetArticleTagByAId(articleId uint32) (model.ArticleTag, error) {
	articleTag := model.ArticleTag{ArticleId: articleId}
	return articleTag.GetByAId(d.engine)
}

func (d *Dao) CreateArticleTag(articleId, tagId uint32, createdBy string) error {
	articleTag := model.ArticleTag{
		Model:     &model.Model{CreatedBy: createdBy},
		ArticleId: articleId,
		TagId:     tagId,
	}
	return articleTag.Create(d.engine)
}

func (d *Dao) DeleteArticleTag(articleId, tagId uint32) error {
	articleTag := model.ArticleTag{
		ArticleId: articleId,
		TagId:     tagId,
	}
	return articleTag.Delete(d.engine)
}

func (d *Dao) DeleteAllTagFromArticle(articleId uint32) error {
	articleTag := model.ArticleTag{
		ArticleId: articleId,
	}
	return articleTag.DeleteAllTagFromArticle(d.engine)
}
