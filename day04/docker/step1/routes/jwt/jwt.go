package jwt

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = map[string]User{}

// Useful response types
type AuthResponse struct {
	AccessToken string `json:"accessToken"`
	User        `json:"user"`
	Message     string `json:"message"`
}

type MeResponse struct {
	User    `json:"user"`
	Message string `json:"message"`
}
