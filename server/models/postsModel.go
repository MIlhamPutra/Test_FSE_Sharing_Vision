package models

import "time"

type PostsModel struct {
	Id   int    `json:"id" validate:"required" sql:"AUTO_INCREMENT"`
	Title string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	Category string `json:"category" validate:"required"`
	CreatedDate time.Time `json:"created_date" validate:"required"`
	UpdatedDate time.Time `json:"updated_date" validate:"required"`
	Status string `json:"status" validate:"required"`
}