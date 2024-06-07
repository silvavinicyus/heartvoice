package models

import (
	"errors"
	"heartvoice/src/security"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Password  string    `json:"password,omitempty"`
	Email     string    `json:"email,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (u *User) Prepare(step string) error {
	if erro := u.validate(step); erro != nil {
		return erro
	}

	erro := u.format(step)
	if erro != nil {
		return erro
	}

	return nil
}

func (u *User) validate(step string) error {
	if u.Name == "" {
		return errors.New("name is a mandatory parameter and should not be empty")
	}

	if u.Email == "" {
		return errors.New("email is a mandatory parameter and should not be empty")
	}

	if erro := checkmail.ValidateFormat(u.Email); erro != nil {
		return errors.New("this email is not valid")
	}

	if u.Nickname == "" {
		return errors.New("nickname is a mandatory parameter and should not be empty")
	}

	if step == "signup" && u.Password == "" {
		return errors.New("password is a mandatory parameter and should not be empty")
	}

	return nil
}

func (u *User) format(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.Nickname = strings.TrimSpace(u.Nickname)
	u.Name = strings.TrimSpace(u.Name)
	u.Name = strings.TrimSpace(u.Name)

	if step == "signup" {
		hashedPassword, hashError := security.Hash(u.Password)

		if hashError != nil {
			return hashError
		}

		u.Password = string(hashedPassword)
	}

	return nil
}
