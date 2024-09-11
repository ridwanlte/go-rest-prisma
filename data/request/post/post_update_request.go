package request

type PostUpdateRequest struct {
	Id          string
	Title       string `validate:"required min=1, max=100" json:"title"`
	Description string `json:"description"`
	Published   bool   `json:"published"`
}
