package requests

type SignUpRequest struct {
	Username string `json:"username" required:"true"`
	Password string `json:"password" required:"true"`
	Nickname string `json:"nickname" required:"true"`
	Email    string `json:"email" required:"true"`
}
