package models

type Pagination struct {
	Page  int `form:"page,default=1" json:"page" binding:"omitempty,min=1,max=100"`
	Limit int `form:"limit,default=10" json:"limit" binding:"omitempty,min=1,max=50"`
}
