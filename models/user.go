package models

// User model
type User struct {
	ID           int    `json:"id,omitempty"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	DisplayName  string `json:"display_name,omitempty"`
	Age          int    `json:"age,omitempty"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	PhoneNumber  string `json:"phone_number,omitempty"`
}
