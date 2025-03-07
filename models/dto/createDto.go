package dto

type CreateDto struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
}
