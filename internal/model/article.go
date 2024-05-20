package model

import (
	"github.com/camtrik/gin-blog/pkg/app"
	"gorm.io/gorm"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

type ArticleInfo struct {
	Id      uint32 `json:"id"`
	TagId   uint32 `json:"tag_id"`
	TagName string `json:"tag_name"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) CountByTag(db *gorm.DB, tagId uint32) (int64, error) {
	var count int64
	err := db.Table(ArticleTag{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+a.TableName()+"` AS a ON at.article_id = a.id").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").
		Where("at.tag_id = ? AND a.state = ? AND a.is_del = ?", tagId, a.State, 0).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// TODO: Implement ArticleRow to include the tag information
func (a Article) ListByTag(db *gorm.DB, tagId uint32, pageOffset, pageSize int) ([]*ArticleInfo, error) {
	var articlesInfo []*ArticleInfo
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	err = db.Table(ArticleTag{}.TableName()+" AS at").
		Select("a.id as id, at.tag_id as tag_id, t.name AS tag_name, a.title as title, a.`desc` as `desc`, a.content as content").
		Joins("LEFT JOIN `"+a.TableName()+"` AS a ON at.article_id = a.id").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").
		Where("at.tag_id = ? AND a.state = ? AND a.is_del = ?", tagId, a.State, 0).
		Find(&articlesInfo).Error
	if err != nil {
		return nil, err
	}
	return articlesInfo, nil
}

func (a Article) Count(db *gorm.DB) (int64, error) {
	var count int64
	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}
	db = db.Where("state = ?", a.State)
	if err := db.Model(&a).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (a Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var articles []*Article
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}
	db = db.Where("state = ?", a.State)
	if err = db.Model(&a).Where("is_del = ?", 0).Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	db = db.Where("id = ? AND is_del = ?", a.Id, 0)
	err := db.First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, err
	}
	return article, nil
}

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Article) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(a).Where("id = ? AND is_del = ?", a.Id, 0).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

func (a Article) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", a.Id, 0).Delete(&a).Error
}
