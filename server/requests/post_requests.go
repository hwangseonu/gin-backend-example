package requests

type CreatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type AddCommentRequest struct {
	Content string `json:"content"`
}