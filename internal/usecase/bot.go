package usecase

import (
	"context"
	"fmt"
	"github.com/Nekr0bz/timetable_bot/internal/entity"
	"github.com/Nekr0bz/timetable_bot/internal/usecase/repo"
	tele "gopkg.in/telebot.v3"
)

type BotUseCase interface {
	SignUpUser(context.Context, tele.Context) (bool, error)
}

type botUseCase struct {
	usrRepo repo.UserRepo
}

func NewBotUseCase(usrRepo repo.UserRepo) BotUseCase {
	return &botUseCase{
		usrRepo: usrRepo,
	}
}

// SignUpUser returns true if user was created
// and false if user already exists
func (u *botUseCase) SignUpUser(ctx context.Context, c tele.Context) (bool, error) {
	teleUser := c.Sender()
	if teleUser.IsBot {
		// TODO: custom errors!
		return false, fmt.Errorf("user is bot")
	}

	user := entity.MarshalTeleUser(teleUser)
	return u.usrRepo.GetOrCreateUser(ctx, user)
}
