package model

type Category struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	ImageID       string `json:"image_ID"`
	Price         string `json:"price"`
	CurrencyID    string `json:"currency_ID"`
	Rating        string `json:"rating"`
	CategoryID    string `json:"category_ID"`
	Specification string `json:"specification"`
	CreatedAt     string `json:"created_At"`
	UpdatedAt     string `json:"updated_At"`
}
