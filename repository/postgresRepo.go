package repository

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository() *PostgresRepository {
	connStr := "user=elquimiquito password=emgyrvaremreys dbname=wb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Insert(ctx context.Context, e *Employee) {
	var id int
	t, err := time.Parse(time.DateOnly, e.Birthdate)
	if err != nil {
		log.Printf("Error %s", err)
	}
	err = r.db.QueryRowContext(ctx, "insert into employees(id, name, birth_date) values ($1, $2, $3) returning id", e.ID, e.Name, t).Scan(&id)
	if err != nil {
		log.Printf("Error %s", err)
	}
}

func (r *PostgresRepository) GetById(ctx context.Context, id int) *Employee {
	row := r.db.QueryRowContext(ctx, "select * from employees where id = $1", id)
	var employee Employee
	err := row.Scan(&employee.ID, &employee.Name, &employee.Birthdate)
	if err != nil {
		log.Printf("Ошибка %s", err)
	}
	return &employee
}

func (r *PostgresRepository) GetBirthdays(ctx context.Context) *[]Employee {
	rows, err := r.db.QueryContext(ctx, "select birthday_list() from employees")
	if err != nil {
		log.Fatal(err)
	}
	var employees []Employee
	var e Employee
	for rows.Next() {
		err := rows.Scan(&e.Name)
		if err != nil {
			log.Fatal(err)
		}
		employees = append(employees, e)
	}
	return &employees
}
