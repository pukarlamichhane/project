package model

type Model struct {
	ID          int    `json:"id"`
	ImgURL      string `json:"imgurl"`
	Price       string `json:"price"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Order struct {
	Item     string `json:"selectedItem"`
	Quantity string `json:"quantity"`
	Name     string `json:"person"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}
