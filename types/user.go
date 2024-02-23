package types

import "golang.org/x/crypto/bcrypt"

type UserInputParams struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID           string `json:"id,omitempty"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

func NewUserFromParams(params UserInputParams) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), 12)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:    params.FirstName,
		LastName:     params.LastName,
		Email:        params.Email,
		PasswordHash: string(hash),
	}, nil
}
