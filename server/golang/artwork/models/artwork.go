package models

type Owner struct {
	Address  string `json:"address"`
	IsSigned bool   `json:"isSigned"`
	Artworks []Artwork
}

type Artwork struct {
	ArtId              string  `json:"artId"`
	ChainId            string  `json:"chainId"`
	Title              string  `json:"title"`
	Price              float64 `json:"price"`
	IsAvailableForSale bool    `json:"is_available_for_sale"`
	Description        string  `json:"description"`
	IsUploaded         bool    `json:"isUploaded"`
	Data               []Data  `json:"data"`
	// Log [] Owner					`json:"log"`
}

type Data struct {
	Extra map[string]string `json:"extra"`
}
