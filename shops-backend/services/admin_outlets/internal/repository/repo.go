package repository

import (
	"github.com/student/shops/services/admin_outlets/internal"
	"github.com/student/shops/services/common"
)

type Repository struct {
	DB *common.Database
}

func NewRepository(db *common.Database) *Repository {
	return &Repository{DB: db}
}

func (ir *Repository) GetList(id int) ([]internal.AdminOutlets, error) {
	query := "SELECT * FROM admin_outlets WHERE admin_id = $1"
	rows, err := ir.DB.Connection.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var adminOutlets []internal.AdminOutlets
	for rows.Next() {
		var adminOutlet internal.AdminOutlets
		if err := rows.Scan(&adminOutlet.ID, &adminOutlet.AdminId, &adminOutlet.OutletId); err != nil {
			return nil, err
		}
		adminOutlets = append(adminOutlets, adminOutlet)
	}

	return adminOutlets, nil
}

func (ir *Repository) Get(id int) (*internal.AdminOutlets, error) {
	query := "SELECT * FROM admin_outlets WHERE id = $1"
	row := ir.DB.Connection.QueryRow(query, id)

	var adminOutlets internal.AdminOutlets
	if err := row.Scan(&adminOutlets.ID, &adminOutlets.AdminId, &adminOutlets.OutletId); err != nil {
		return nil, err
	}

	return &adminOutlets, nil
}

func (ir *Repository) Create(newItem *internal.AdminOutlets) error {
	query := "INSERT INTO admin_outlets (admin_id, outlet_id) VALUES ($1, $2) RETURNING id"
	err := ir.DB.Connection.QueryRow(query, newItem.AdminId, newItem.OutletId).Scan(&newItem.ID)
	return err
}

func (ir *Repository) Update(id int, adminId int, outletId int) error {
	query := "UPDATE admin_outlets SET admin_id = $1, outlet_id = $2 WHERE id = $3"
	_, err := ir.DB.Connection.Exec(query, adminId, outletId, id)
	return err
}

func (ir *Repository) Delete(id int) error {
	query := "DELETE FROM admin_outlets WHERE id = $1"
	_, err := ir.DB.Connection.Exec(query, id)
	return err
}
