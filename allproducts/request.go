package allproducts

import "encoding/json"

type AllProductRequest struct {
	Id           int         `json:"id"`
	Name_product string      `json:"name_product"`
	Image_url    string      `json:"image_url"`
	Description  string      `json:"description"`
	Price        json.Number `json:"price"`
	Name_user    string      `json:"name_user"`
	Email_user   string      `json:"email_user"`
	Category     string      `json:"category"`
}
