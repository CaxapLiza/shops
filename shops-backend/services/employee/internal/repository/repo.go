package repository

import (
	"github.com/student/shops/services/common"
	"github.com/student/shops/services/employee/internal"
)

type Repository struct {
	DB *common.Database
}

func NewRepository(db *common.Database) *Repository {
	return &Repository{DB: db}
}

func (ir *Repository) GetList() ([]internal.Employee, error) {
	query := "SELECT * FROM employee"
	rows, err := ir.DB.Connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []internal.Employee
	for rows.Next() {
		var employee internal.Employee
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.Itn, &employee.Passport, &employee.Snils, &employee.Phone, &employee.AccountId); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
}

func (ir *Repository) Authenticate(id int) (*internal.Employee, error) {
	query := "SELECT * FROM employee WHERE account_id = $1"
	row := ir.DB.Connection.QueryRow(query, id)

	var employee internal.Employee
	if err := row.Scan(&employee.ID, &employee.Name, &employee.Itn, &employee.Passport, &employee.Snils, &employee.Phone, &employee.AccountId); err != nil {
		return nil, err
	}

	return &employee, nil
}

func (ir *Repository) Get(id int) (*internal.Employee, error) {
	query := "SELECT * FROM employee WHERE id = $1"
	row := ir.DB.Connection.QueryRow(query, id)

	var employee internal.Employee
	if err := row.Scan(&employee.ID, &employee.Name, &employee.Itn, &employee.Passport, &employee.Snils, &employee.Phone, &employee.AccountId); err != nil {
		return nil, err
	}

	return &employee, nil
}

func (ir *Repository) Create(newItem *internal.Employee) error {
	query := "INSERT INTO employee (name, itn, passport, snils, phone, account_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	err := ir.DB.Connection.QueryRow(query, newItem.Name, newItem.Itn, newItem.Passport, newItem.Snils, newItem.Phone, newItem.AccountId).Scan(&newItem.ID)
	return err
}

func (ir *Repository) Update(id int, name string, itn string, passport string, snils string, phone string, accountId int) error {
	query := "UPDATE employee SET name = $1, itn = $2, passport = $3, snils = $4, phone = $5, account_id = $6 WHERE id = $7"
	_, err := ir.DB.Connection.Exec(query, name, itn, passport, snils, phone, accountId, id)
	return err
}

func (ir *Repository) Delete(id int) error {
	query := "DELETE FROM employee WHERE id = $1"
	_, err := ir.DB.Connection.Exec(query, id)
	return err
}
