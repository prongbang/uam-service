package auth

type Login struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Credential struct {
	Token string   `json:"token"`
	Role  []string `json:"roles"`
}
