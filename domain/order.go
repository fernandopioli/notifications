package domain

type Order struct {
    ID       string   `json:"id"`
	Total    float64  `json:"total"`
	Customer Customer `json:"customer"`
}