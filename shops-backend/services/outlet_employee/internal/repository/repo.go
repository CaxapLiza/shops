package repository

import (
	"github.com/student/shops/services/common"
	"github.com/student/shops/services/outlet_employee/internal"
)

type Repository struct {
	DB *common.Database
}

func NewRepository(db *common.Database) *Repository {
	return &Repository{DB: db}
}

func (ir *Repository) GetList(id int) ([]internal.OutletEmployee, error) {
	query := "SELECT * FROM outlet_employee WHERE outlet_id = $1"
	rows, err := ir.DB.Connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var outletEmployees []internal.OutletEmployee
	for rows.Next() {
		var outletEmployee internal.OutletEmployee
		if err := rows.Scan(&outletEmployee.ID, &outletEmployee.OutletId, &outletEmployee.EmployeeId); err != nil {
			return nil, err
		}
		outletEmployees = append(outletEmployees, outletEmployee)
	}

	return outletEmployees, nil
}

func (ir *Repository) Get(id int) (*internal.OutletEmployee, error) {
	query := "SELECT * FROM outlet_employee WHERE id = $1"
	row := ir.DB.Connection.QueryRow(query, id)

	var outletEmployee internal.OutletEmployee
	if err := row.Scan(&outletEmployee.ID, &outletEmployee.OutletId, &outletEmployee.EmployeeId); err != nil {
		return nil, err
	}

	return &outletEmployee, nil
}

func (ir *Repository) Create(newItem *internal.OutletEmployee) error {
	query := "INSERT INTO outlet_employee (outlet_id, employee_id) VALUES ($1, $2) RETURNING id"
	err := ir.DB.Connection.QueryRow(query, newItem.OutletId, newItem.EmployeeId).Scan(&newItem.ID)
	return err
}

func (ir *Repository) Update(id int, outletId int, employeeId int) error {
	query := "UPDATE outlet_employee SET outlet_id = $1, employee_id = $2 WHERE id = $3"
	_, err := ir.DB.Connection.Exec(query, outletId, employeeId, id)
	return err
}

func (ir *Repository) Delete(id int) error {
	query := "DELETE FROM outlet_employee WHERE id = $1"
	_, err := ir.DB.Connection.Exec(query, id)
	return err
}
