package request

type PostCreateRequest struct {
	Title       string `validate:"required min=1, max=100" json:"title"`
	Description string `json:"description"`
	Published   bool   `json:"published"`
}
