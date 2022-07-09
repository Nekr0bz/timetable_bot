package usecase

type botUseCase struct{}

func NewBotUseCase() *botUseCase {
	return &botUseCase{}
}

func (u *botUseCase) GetHello() string {
	return "HELLO WORLD!"
}
