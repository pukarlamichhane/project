package model

type Model struct {
	ID          int    `json:"id"`
	ImgURL      string `json:"imgurl"`
	Price       string `json:"price"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
