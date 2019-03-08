package responses

type AuthResponse struct {
	Access string `json:"access"`
	Refresh string `json:"refresh"`
}
