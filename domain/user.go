package domain

// User will hold the information of people using this applcation
type User struct {
	id       string `json:"id,omitempty"`
	name     string `json:"name"`
	email    string `json:"email"`
	password string `json:"password"`
}
