package usecase

import (
	"context"
	"fmt"
	"github.com/Nekr0bz/timetable_bot/internal/entity"
	"github.com/Nekr0bz/timetable_bot/internal/usecase/repo"
	tele "gopkg.in/telebot.v3"
)

type BotUseCase interface {
	GetHello() string
	SignUpUser(tele.Context) (bool, error)
}

type botUseCase struct {
	usrRepo repo.UserRepo
}

func NewBotUseCase(usrRepo repo.UserRepo) BotUseCase {
	return &botUseCase{
		usrRepo: usrRepo,
	}
}

func (u *botUseCase) GetHello() string {
	return "HELLO WORLD!"
}

// SignUpUser returns true if user was created
// and false if user already exists
func (u *botUseCase) SignUpUser(c tele.Context) (bool, error) {
	teleUser := c.Sender()
	if teleUser.IsBot {
		return false, fmt.Errorf("user is bot")
	}

	user := entity.MarshalTeleUser(teleUser)
	return u.usrRepo.GetOrCreateUser(context.TODO(), user)
}
