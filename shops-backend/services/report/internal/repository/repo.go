package repository

import (
	"github.com/student/shops/services/common"
	"github.com/student/shops/services/report/internal"
	"time"
)

type Repository struct {
	DB *common.Database
}

func NewRepository(db *common.Database) *Repository {
	return &Repository{DB: db}
}

func (ir *Repository) GetList(id int) ([]internal.Report, error) {
	query := "SELECT * FROM report WHERE outlet_id = $1"
	rows, err := ir.DB.Connection.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []internal.Report
	for rows.Next() {
		var report internal.Report
		if err := rows.Scan(&report.ID, &report.Income, &report.Expenses, &report.Coefficient, &report.Date, &report.OutletId); err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}

	return reports, nil
}

func (ir *Repository) Get(id int) (*internal.Report, error) {
	query := "SELECT * FROM report WHERE id = $1"
	row := ir.DB.Connection.QueryRow(query, id)

	var report internal.Report
	if err := row.Scan(&report.ID, &report.Income, &report.Expenses, &report.Coefficient, &report.Date, &report.OutletId); err != nil {
		return nil, err
	}

	return &report, nil
}

func (ir *Repository) Create(newItem *internal.Report) error {
	query := "INSERT INTO report (income, expenses, coefficient, date, outlet_id) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := ir.DB.Connection.QueryRow(query, newItem.Income, newItem.Expenses, newItem.Coefficient, newItem.Date, newItem.OutletId).Scan(&newItem.ID)
	return err
}

func (ir *Repository) Update(id int, income float32, expenses float32, coefficient float32, date time.Time, outletId int) error {
	query := "UPDATE report SET income = $1, expenses = $2, coefficient = $3, date = $4, outlet_id = $5 WHERE id = $6"
	_, err := ir.DB.Connection.Exec(query, income, expenses, coefficient, date, outletId, id)
	return err
}

func (ir *Repository) Delete(id int) error {
	query := "DELETE FROM report WHERE id = $1"
	_, err := ir.DB.Connection.Exec(query, id)
	return err
}
