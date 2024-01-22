package dto

type UserGetDTO struct {
	ID string `json:"id"`
}

type UserCreateDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserUpdateDTO struct {
	Email     string `json:"email,omitempty"`
	Username  string `json:"username,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

type UserResponseDTO struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}
