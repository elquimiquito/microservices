package proxy

import (
	"Wildberries_services/repository"
	"context"
)

func Insert(e *repository.Employee) {
	ctx := context.Background()
	repo := repository.NewPostgresRepository()
	repo.Insert(ctx, e)
}

func GetByID(id int) *repository.Employee {
	ctx := context.Background()
	repo := repository.NewPostgresRepository()
	return repo.GetById(ctx, id)
}

func GetBirthdays() *[]repository.Employee {
	ctx := context.Background()
	repo := repository.NewPostgresRepository()
	return repo.GetBirthdays(ctx)
}
