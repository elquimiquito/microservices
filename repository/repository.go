package repository

import "context"

type Repository interface {
	Insert(ctx context.Context, e *Employee)
	GetByID(ctx context.Context, id int)
	GetBirthdays(ctx context.Context)
}

type Employee struct {
	ID        int
	Name      string
	Birthdate string
}
