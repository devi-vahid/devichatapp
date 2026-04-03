package domain

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Family   string `json:"family"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}