package models

import (
	"DevBookAPI/src/security"
	"errors"
	"regexp"
	"strings"
	"time"
)

type Users struct {
	Id        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"-"`
}

func (users *Users) Prepare(step string) error {
	err := users.ValidateFields(step)
	if err != nil {
		return err
	}

	err = users.Format(step)
	if err != nil {
		return err
	}

	return nil
}

func (users *Users) Format(step string) error {
	users.Name = strings.TrimSpace(users.Name)
	users.Nick = strings.TrimSpace(users.Nick)
	users.Email = strings.TrimSpace(users.Email)

	if step == "registration" {
		passwordHash, err := security.Hash(users.Password)
		if err != nil {
			return err
		}

		users.Password = string(passwordHash)
	}

	return nil
}

func (users *Users) ValidateFields(step string) error {
	if users.Name == "" {
		return errors.New("name is required")
	}
	if users.Nick == "" {
		return errors.New("nick is required")
	}
	if users.Email == "" {
		return errors.New("email is required")
	}
	if users.Password == "" && step == "registration" {
		return errors.New("password is required")
	}

	err := users.IsValidEmail()
	if err != nil {
		return err
	}

	return nil
}

func (users *Users) IsValidEmail() error {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	emailIsValid := emailRegex.MatchString(users.Email)

	if !emailIsValid {
		return errors.New("email is invalid")
	}

	return nil
}
