package dto

type Image struct {
	Key string `json:"key"`
	// Other fields...
}
type CreateRequest struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Image     string `json:"image"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	DeletedBy string `json:"deleted_by"`
}

type SocmedResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Image     string `json:"image"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedBy string `json:"created_by"`
}
