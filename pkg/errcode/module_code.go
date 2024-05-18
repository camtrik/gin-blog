package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "Failed to get tag list")
	ErrorCreateTagFail  = NewError(20010002, "Failed to create tag")
	ErrorUpdateTagFail  = NewError(20010003, "Failed to update tag")
	ErrorDeleteTagFail  = NewError(20010004, "Failed to delete tag")
	ErrorCountTagFail   = NewError(20010005, "Failed to count tags")
)

