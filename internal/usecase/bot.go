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
	Universities(context.Context) ([]entity.University, error)
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

// Universities returns all universities with their faculties and courses and groups
func (u *botUseCase) Universities(context.Context) ([]entity.University, error) {
	// TODO: implement
	return []entity.University{
		entity.University{
			Id:   1,
			Name: "МГУ 1",
			GroupTypes: []entity.GroupType{
				entity.GroupType{
					Id:   1,
					Name: "1.1 Очная группа 🧑‍🎓",
					Faculties: []entity.Faculty{
						entity.Faculty{
							Id:   1,
							Name: "1.1.1 Факультет информационных технологий",
							Courses: []entity.Course{
								entity.Course{
									Id:   1,
									Name: "1.1.1.1 Информационные технологии",
									Groups: []entity.Group{
										entity.Group{
											Id:   1,
											Name: "1.1.1.1.1 ИТ-1",
										},
										entity.Group{
											Id:   2,
											Name: "1.1.1.1.2 ИТ-2",
										},
									},
								},
							},
						},
						entity.Faculty{
							Id:   2,
							Name: "1.1.2 Факультет информационных технологий 1 1",
							Courses: []entity.Course{
								entity.Course{
									Id:   3,
									Name: "1.1.2.1 Информационные технологии 1 1",
									Groups: []entity.Group{
										entity.Group{
											Id:   3,
											Name: "1.1.2.1.1 ИТ-1 1 1",
										},
										entity.Group{
											Id:   4,
											Name: "1.1.2.1.2 ИТ-2 1 1 ",
										},
									},
								},
							},
						},
					},
				},
				entity.GroupType{
					Id:   2,
					Name: "1.2 Заочная группа 🧑‍🎓",
					Faculties: []entity.Faculty{
						entity.Faculty{
							Id:   1,
							Name: "1.2.1 Факультет информационных технологий",
							Courses: []entity.Course{
								entity.Course{
									Id:   1,
									Name: "1.2.1.1 Информационные технологии",
									Groups: []entity.Group{
										entity.Group{
											Id:   1,
											Name: "1.2.1.1.1 ИТ-1",
										},
										entity.Group{
											Id:   2,
											Name: "1.2.1.1.2 ИТ-2",
										},
									},
								},
							},
						},
						entity.Faculty{
							Id:   2,
							Name: "1.2.2 Факультет информационных технологий 1 1",
							Courses: []entity.Course{
								entity.Course{
									Id:   3,
									Name: "1.2.2.1 Информационные технологии 1 1",
									Groups: []entity.Group{
										entity.Group{
											Id:   3,
											Name: "1.2.2.1.1 ИТ-1 1 1",
										},
										entity.Group{
											Id:   4,
											Name: "1.2.2.1.2 ИТ-2 1 1 ",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		entity.University{
			Id:   2,
			Name: "МГУ 2",
			GroupTypes: []entity.GroupType{
				entity.GroupType{
					Id:   1,
					Name: "2.1 Очная группа 🧑‍🎓",
					Faculties: []entity.Faculty{
						entity.Faculty{
							Id:   5,
							Name: "Факультет информационных технологий 2",
							Courses: []entity.Course{
								entity.Course{
									Id:   5,
									Name: "Информационные технологии 2",
									Groups: []entity.Group{
										entity.Group{
											Id:   5,
											Name: "ИТ-1 2",
										},
										entity.Group{
											Id:   6,
											Name: "ИТ-2 2",
										},
									},
								},
							},
						},
						entity.Faculty{
							Id:   7,
							Name: "Факультет информационных технологий 2 2",
							Courses: []entity.Course{
								entity.Course{
									Id:   8,
									Name: "Информационные технологии 2 2",
									Groups: []entity.Group{
										entity.Group{
											Id:   8,
											Name: "ИТ-1 2 2",
										},
										entity.Group{
											Id:   9,
											Name: "ИТ-2 2 2",
										},
									},
								},
							},
						},
					},
				},
				entity.GroupType{
					Id:   2,
					Name: "2.2 Заочная группа 🧑‍🎓",
					Faculties: []entity.Faculty{
						entity.Faculty{
							Id:   5,
							Name: "Факультет информационных технологий 2",
							Courses: []entity.Course{
								entity.Course{
									Id:   5,
									Name: "Информационные технологии 2",
									Groups: []entity.Group{
										entity.Group{
											Id:   5,
											Name: "ИТ-1 2",
										},
										entity.Group{
											Id:   6,
											Name: "ИТ-2 2",
										},
									},
								},
							},
						},
						entity.Faculty{
							Id:   7,
							Name: "Факультет информационных технологий 2 2",
							Courses: []entity.Course{
								entity.Course{
									Id:   8,
									Name: "Информационные технологии 2 2",
									Groups: []entity.Group{
										entity.Group{
											Id:   8,
											Name: "ИТ-1 2 2",
										},
										entity.Group{
											Id:   9,
											Name: "ИТ-2 2 2",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}, nil
}
