package container

import (
	"assignment-2/app/adapter/database"
	order "assignment-2/app/respository/order"
	"assignment-2/app/usecase"
)

type Container struct {
	Usecase usecase.Usecase
}

func SetupContainer() Container {
	db := database.SetupDatabase().DB

	repo := order.NewOrderRepository(db)
	usecase := usecase.NewUsecase(repo)

	return Container{
		Usecase: usecase,
	}

}
