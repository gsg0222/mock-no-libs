package usecase

import (
	"context"
	"fmt"

	"tsuchiya.blog/mock-no-libs/pkg/repository"
)

type Usecase struct {
	repo repository.Repository
}

func NewUsecase(repo repository.Repository) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) Run(id string) error {
	ctx := context.Background()
	output, err := u.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	fmt.Println(output)
	return nil
}
