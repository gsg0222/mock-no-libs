package main

import (
	"fmt"

	"tsuchiya.blog/mock-no-libs/pkg/gateway"
	"tsuchiya.blog/mock-no-libs/pkg/usecase"
)

func main() {
	fmt.Println("Input ID:")
	var id string
	fmt.Scan(&id)

	// RepositoryをUsecaseにDI
	usecase := usecase.NewUsecase(&gateway.SampelRepository{})
	usecase.Run(id)
}
