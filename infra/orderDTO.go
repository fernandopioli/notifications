package infra

type OrderDTO struct {
	ID       string  `json:"id"`
	Total    float64 `json:"total"`
	Customer struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"customer"`
}