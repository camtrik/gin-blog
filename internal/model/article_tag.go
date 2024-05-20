package model

import "gorm.io/gorm"

type ArticleTag struct {
	*Model
	ArticleId uint32 `json:"article_id"`
	TagId     uint32 `json:"tag_id"`
}

func (at ArticleTag) TableName() string {
	return "blog_article_tag"
}

func (at ArticleTag) Create(db *gorm.DB) error {
	return db.Table(at.TableName()).Create(&at).Error
}

func (at ArticleTag) Delete(db *gorm.DB) error {
	return db.Table(at.TableName()).Where("article_id = ? AND tag_id = ? AND is_del = ?", at.ArticleId, at.TagId, 0).Delete(&at).Error
}

func (at ArticleTag) DeleteAllTagFromArticle(db *gorm.DB) error {
	return db.Table(at.TableName()).Where("article_id = ? AND is_del = ?", at.ArticleId, 0).Delete(&at).Error
}

// Get First tag of the article
// TODO: get all tags of the article
func (at ArticleTag) GetByAId(db *gorm.DB) (ArticleTag, error) {
	var articleTag ArticleTag
	err := db.Table(at.TableName()).Where("article_id = ? AND is_del = ?", at.ArticleId, 0).First(&articleTag).Error
	return articleTag, err
}
