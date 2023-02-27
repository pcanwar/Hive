package models

type Owner struct {
	Address  string `json:"address"`
	IsSigned bool   `json:"isSigned"`
	Artworks []Artwork
}

type Artwork struct {
	ID                 string  `json:"_id"`
	Title              string  `json:"title"`
	Price              float64 `json:"price"`
	IsAvailableForSale bool    `json:"is_available_for_sale"`
}
