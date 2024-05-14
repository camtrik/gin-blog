package model

type Model struct {
	Id         int32  `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	CreatedOn  int32  `json:"created_on"`
	DeletedOn  int32  `json:"deleted_on"`
	IsDel      int8   `json:"is_del"`
	ModifiedBy string `json:"modified_by"`
	ModifiedOn int32  `json:"modified_on"`
}
