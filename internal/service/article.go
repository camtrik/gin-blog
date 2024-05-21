package service

import (
	"github.com/camtrik/gin-blog/internal/dao"
	"github.com/camtrik/gin-blog/internal/model"
	"github.com/camtrik/gin-blog/pkg/app"
)

type GetArticleRequest struct {
	Id uint32 `form:"id" binding:"required,gte=1"`
}

// article information with first tag
// TODO: get all tags
type ArticleWithTag struct {
	Id            uint32     `json:"id"`
	Title         string     `json:"title"`
	Desc          string     `json:"desc"`
	Content       string     `json:"content"`
	CoverImageUrl string     `json:"cover_image_url"`
	State         uint8      `json:"state"`
	Tag           *model.Tag `json:"tag"`
}

type ArticleListRequest struct {
	TagId uint32 `form:"tag_id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title         string `form:"title" binding:"required,min=3,max=100"`
	Desc          string `form:"desc" binding:"required,min=3,max=255"`
	Content       string `form:"content" binding:"required,min=3"`
	CoverImageUrl string `form:"cover_image_url"`
	CreatedBy     string `form:"created_by" binding:"required,min=3,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	Id            uint32 `form:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"max=100"`
	Desc          string `form:"desc" binding:"max=255"`
	Content       string `form:"content"`
	CoverImageUrl string `form:"cover_image_url"`
	State         uint8  `form:"state" binding:"oneof=0 1"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteArticleRequest struct {
	Id uint32 `form:"id" binding:"required,gte=1"`
}

type AddArticleTagRequest struct {
	ArticleId uint32 `form:"article_id" binding:"required,gte=1"`
	TagId     uint32 `form:"tag_id" binding:"required,gte=1"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
}

type DeleteArticleTagRequest struct {
	ArticleId uint32 `form:"article_id" binding:"required,gte=1"`
	TagId     uint32 `form:"tag_id" binding:"required,gte=1"`
}

func (svc *Service) GetArticle(param *GetArticleRequest) (*ArticleWithTag, error) {
	article, err := svc.dao.GetArticle(param.Id, 0)
	if err != nil {
		return nil, err
	}

	articleTag, err := svc.dao.GetArticleTagByAId(article.Id)
	if err != nil {
		return nil, err
	}

	tag, err := svc.dao.GetTag(articleTag.TagId)
	if err != nil {
		return nil, err
	}

	return &ArticleWithTag{
		Id:            article.Id,
		Title:         article.Title,
		Desc:          article.Desc,
		Content:       article.Content,
		CoverImageUrl: article.CoverImageUrl,
		Tag:           &model.Tag{Model: &model.Model{Id: tag.Id}, Name: tag.Name},
	}, nil
}

func (svc *Service) GetArticleList(param *ArticleListRequest, pager *app.Pager) ([]*ArticleWithTag, int64, error) {
	articleCount, err := svc.dao.CountArticleByTag(param.TagId, param.State)
	if err != nil {
		return nil, 0, err
	}

	articlesInfo, err := svc.dao.ListArticleByTag(param.TagId, param.State, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var articlesWithTag []*ArticleWithTag
	for _, articleInfo := range articlesInfo {
		articlesWithTag = append(articlesWithTag, &ArticleWithTag{
			Id:      articleInfo.Id,
			Title:   articleInfo.Title,
			Desc:    articleInfo.Desc,
			Content: articleInfo.Content,
			Tag:     &model.Tag{Model: &model.Model{Id: articleInfo.TagId}, Name: articleInfo.TagName},
		})
	}
	return articlesWithTag, articleCount, nil
}

func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	return svc.dao.CreateArticle(&dao.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		CreatedBy:     param.CreatedBy,
	})
}

func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error {
	return svc.dao.UpdateArticle(&dao.Article{
		Id:            param.Id,
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		ModifiedBy:    param.ModifiedBy,
	})
}

func (svc *Service) DeleteArticle(param *DeleteArticleRequest) error {
	err := svc.dao.DeleteArticle(param.Id)
	if err != nil {
		return err
	}

	err = svc.dao.DeleteAllTagFromArticle(param.Id)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) AddArticleTag(param *AddArticleTagRequest) error {
	return svc.dao.CreateArticleTag(param.ArticleId, param.TagId, param.CreatedBy)
}

func (svc *Service) DeleteArticleTag(param *DeleteArticleTagRequest) error {
	return svc.dao.DeleteArticleTag(param.ArticleId, param.TagId)
}
