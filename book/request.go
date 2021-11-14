package book

import "encoding/json"

type BookRequest struct {
	Name_product string      `json:"name_product" binding:"required"`
	Image_url    string      `json:"image_url" binding:"required"`
	Description  string      `json:"description" binding:"required"`
	Price        json.Number `json:"price" binding:"required,number"`
	Name_user    string      `json:"name_user" binding:"required"`
	Email_user   string      `json:"email_user" binding:"required"`
}
