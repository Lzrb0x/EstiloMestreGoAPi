package models

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model        //id, created_at, updated_at, deleted_at
	Name       string `gorm:"not null"`
	Email      string `gorm:"not null;unique"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewUser(name, email string) (*User, error) {
	user := &User{
		Name:  name,
		Email: email,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() error {

	var errorMessages []string

	if u.Name == "" {
		errorMessages = append(errorMessages, "Name is required")
	}
	if u.Email == "" {
		errorMessages = append(errorMessages, "Email is required")
	}
	if len(errorMessages) > 0 {
		return fmt.Errorf("%s", strings.Join(errorMessages, ", "))
	}

	return nil
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
		UpdatedAt: u.UpdatedAt.Format(time.RFC3339),
	}
}
