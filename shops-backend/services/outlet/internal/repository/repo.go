package repository

import (
	"github.com/student/shops/services/common"
	"github.com/student/shops/services/outlet/internal"
)

type Repository struct {
	DB *common.Database
}

func NewRepository(db *common.Database) *Repository {
	return &Repository{DB: db}
}

func (ir *Repository) GetList(id int) ([]internal.Outlet, error) {
	query := "SELECT * FROM outlet WHERE owner_id = $1"
	rows, err := ir.DB.Connection.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var outlets []internal.Outlet
	for rows.Next() {
		var outlet internal.Outlet
		if err := rows.Scan(&outlet.ID, &outlet.Name, &outlet.Address, &outlet.PlanedProfit, &outlet.OwnerId); err != nil {
			return nil, err
		}
		outlets = append(outlets, outlet)
	}

	return outlets, nil
}

func (ir *Repository) Get(id int) (*internal.Outlet, error) {
	query := "SELECT * FROM outlet WHERE id = $1"
	row := ir.DB.Connection.QueryRow(query, id)

	var outlet internal.Outlet
	if err := row.Scan(&outlet.ID, &outlet.Name, &outlet.Address, &outlet.PlanedProfit, &outlet.OwnerId); err != nil {
		return nil, err
	}

	return &outlet, nil
}

func (ir *Repository) Create(newItem *internal.Outlet) error {
	query := "INSERT INTO outlet (name, address, planned_profit, owner_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := ir.DB.Connection.QueryRow(query, newItem.Name, newItem.Address, newItem.PlanedProfit, newItem.OwnerId).Scan(&newItem.ID)
	return err
}

func (ir *Repository) Update(id int, name string, address string, planedProfit float32, ownerId int) error {
	query := "UPDATE outlet SET name = $1, address = $2, planned_profit = $3, owner_id = $4 WHERE id = $5"
	_, err := ir.DB.Connection.Exec(query, name, address, planedProfit, ownerId, id)
	return err
}

func (ir *Repository) Delete(id int) error {
	query := "DELETE FROM outlet WHERE id = $1"
	_, err := ir.DB.Connection.Exec(query, id)
	return err
}
