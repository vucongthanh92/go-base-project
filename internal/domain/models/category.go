package models

import "time"

type CreateCategoryReq struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCategoryReq struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name" binding:"required"`
	UpdatedAt time.Time `json:"updated_at" binding:"required"`
}
