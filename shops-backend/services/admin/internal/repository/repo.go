package repository

import (
	"github.com/student/shops/services/admin/internal"
	"github.com/student/shops/services/common"
)

type Repository struct {
	DB *common.Database
}

func NewRepository(db *common.Database) *Repository {
	return &Repository{DB: db}
}

func (ir *Repository) GetList() ([]internal.Admin, error) {
	query := "SELECT * FROM admin"
	rows, err := ir.DB.Connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var admins []internal.Admin
	for rows.Next() {
		var admin internal.Admin
		if err := rows.Scan(&admin.ID, &admin.Name, &admin.Itn, &admin.Passport, &admin.Snils, &admin.Phone, &admin.AccountId); err != nil {
			return nil, err
		}
		admins = append(admins, admin)
	}

	return admins, nil
}

func (ir *Repository) Authenticate(id int) (*internal.Admin, error) {
	query := "SELECT * FROM admin WHERE account_id = $1"
	row := ir.DB.Connection.QueryRow(query, id)

	var admin internal.Admin
	if err := row.Scan(&admin.ID, &admin.Name, &admin.Itn, &admin.Passport, &admin.Snils, &admin.Phone, &admin.AccountId); err != nil {
		return nil, err
	}

	return &admin, nil
}

func (ir *Repository) Get(id int) (*internal.Admin, error) {
	query := "SELECT * FROM admin WHERE id = $1"
	row := ir.DB.Connection.QueryRow(query, id)

	var admin internal.Admin
	if err := row.Scan(&admin.ID, &admin.Name, &admin.Itn, &admin.Passport, &admin.Snils, &admin.Phone, &admin.AccountId); err != nil {
		return nil, err
	}

	return &admin, nil
}

func (ir *Repository) Create(newItem *internal.Admin) error {
	query := "INSERT INTO admin (name, itn, passport, snils, phone, account_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	err := ir.DB.Connection.QueryRow(query, newItem.Name, newItem.Itn, newItem.Passport, newItem.Snils, newItem.Phone, newItem.AccountId).Scan(&newItem.ID)
	return err
}

func (ir *Repository) Update(id int, name string, itn string, passport string, snils string, phone string, accountId int) error {
	query := "UPDATE admin SET name = $1, itn = $2, passport = $3, snils = $4, phone = $5, account_id = $6 WHERE id = $7"
	_, err := ir.DB.Connection.Exec(query, name, itn, passport, snils, phone, accountId, id)
	return err
}

func (ir *Repository) Delete(id int) error {
	query := "DELETE FROM admin WHERE id = $1"
	_, err := ir.DB.Connection.Exec(query, id)
	return err
}
