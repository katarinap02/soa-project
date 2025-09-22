package dto

type UserDTO struct {
	Id       string `json:"id"` // PROMENIO sa uuid.UUID na string
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"` // PROMENIO sa model.Role na string
	AccountStatus string        `json:"account_status"`
}
