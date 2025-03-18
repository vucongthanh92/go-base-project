package models

type CreateCategoryReq struct {
	Name string `json:"category_name" binding:"required"`
}
