package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "Failed to get tag list")
	ErrorCreateTagFail  = NewError(20010002, "Failed to create tag")
	ErrorUpdateTagFail  = NewError(20010003, "Failed to update tag")
	ErrorDeleteTagFail  = NewError(20010004, "Failed to delete tag")
	ErrorCountTagFail   = NewError(20010005, "Failed to count tags")
)

var (
	ErrorGetArticleFail       = NewError(20020001, "Failed to get article")
	ErrorGetArticleListFail   = NewError(20020002, "Failed to get article list")
	ErrorCreateArticleFail    = NewError(20020003, "Failed to create article")
	ErrorUpdateArticleFail    = NewError(20020004, "Failed to update article")
	ErrorDeleteArticleFail    = NewError(20020005, "Failed to delete article")
	ErrorAddArticleTagFail    = NewError(20020006, "Failed to add article tag")
	ErrorDeleteArticleTagFail = NewError(20020007, "Failed to delete article tag")

	ErrorUploadFileFail = NewError(20030001, "Failed to upload file")
)
