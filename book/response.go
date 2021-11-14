package book

type BookResponse struct {
	Id           int    `json:"id"`
	Name_product string `json:"name_product"`
	Image_url    string `json:"image_url"`
	Description  string `json:"description"`
	Price        int    `json:"price"`
	Name_user    string `json:"name_user"`
	Email_user   string `json:"email_user"`
}
