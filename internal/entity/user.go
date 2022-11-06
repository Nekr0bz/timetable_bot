package entity

import (
	tele "gopkg.in/telebot.v3"
	"time"
)

type User struct {
	ID           int64     `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Username     string    `json:"username"`
	LanguageCode string    `json:"language_code"`
	CreatedAt    time.Time `json:"created_at"`
}

func MarshalTeleUser(teleUser *tele.User) *User {
	return &User{
		ID:           teleUser.ID,
		FirstName:    teleUser.FirstName,
		LastName:     teleUser.LastName,
		Username:     teleUser.Username,
		LanguageCode: teleUser.LanguageCode,
		CreatedAt:    time.Now(),
	}
}
