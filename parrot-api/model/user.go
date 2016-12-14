package model

import "github.com/anthonynsimon/parrot/parrot-api/errors"
import "strings"

var (
	ErrInvalidEmail = &errors.Error{
		Type:    "InvalidEmail",
		Message: "invalid email"}
	ErrInvalidName = &errors.Error{
		Type:    "InvalidName",
		Message: "invalid name"}
	ErrInvalidPassword = &errors.Error{
		Type:    "InvalidPassword",
		Message: "invalid password"}
)

type UserStorer interface {
	GetUserByID(string) (*User, error)
	GetUserByEmail(string) (*User, error)
	CreateUser(User) (*User, error)
	UpdateUserPassword(User) (*User, error)
	UpdateUserName(User) (*User, error)
	UpdateUserEmail(User) (*User, error)
}

type User struct {
	ID       string `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password,omitempty"`
}

type Validatable interface {
	Validate() error
}

func (u *User) Validate() error {
	var errs []errors.Error
	if !ValidEmail(u.Email) {
		errs = append(errs, *ErrInvalidEmail)
	}
	if !HasMinLength(strings.Trim(u.Name, " "), 1) {
		errs = append(errs, *ErrInvalidName)
	}
	if !HasMinLength(u.Password, 8) {
		errs = append(errs, *ErrInvalidPassword)
	}
	if errs != nil {
		return NewValidationError(errs)
	}
	return nil
}