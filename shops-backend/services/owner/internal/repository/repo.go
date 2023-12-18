package repository

import (
	"github.com/student/shops/services/common"
	"github.com/student/shops/services/owner/internal"
)

type Repository struct {
	DB *common.Database
}

func NewRepository(db *common.Database) *Repository {
	return &Repository{DB: db}
}

func (ir *Repository) GetList() ([]internal.Owner, error) {
	query := "SELECT * FROM owner"
	rows, err := ir.DB.Connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var owners []internal.Owner
	for rows.Next() {
		var owner internal.Owner
		if err := rows.Scan(&owner.ID, &owner.Name, &owner.Itn, &owner.Passport, &owner.Snils, &owner.Phone, &owner.AccountId); err != nil {
			return nil, err
		}
		owners = append(owners, owner)
	}

	return owners, nil
}

func (ir *Repository) Authenticate(id int) (*internal.Owner, error) {
	query := "SELECT * FROM owner WHERE account_id = $1"
	row := ir.DB.Connection.QueryRow(query, id)

	var owner internal.Owner
	if err := row.Scan(&owner.ID, &owner.Name, &owner.Itn, &owner.Passport, &owner.Snils, &owner.Phone, &owner.AccountId); err != nil {
		return nil, err
	}

	return &owner, nil
}

func (ir *Repository) Get(id int) (*internal.Owner, error) {
	query := "SELECT * FROM owner WHERE id = $1"
	row := ir.DB.Connection.QueryRow(query, id)

	var owner internal.Owner
	if err := row.Scan(&owner.ID, &owner.Name, &owner.Itn, &owner.Passport, &owner.Snils, &owner.Phone, &owner.AccountId); err != nil {
		return nil, err
	}

	return &owner, nil
}

func (ir *Repository) Create(newItem *internal.Owner) error {
	query := "INSERT INTO owner (name, itn, passport, snils, phone, account_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	err := ir.DB.Connection.QueryRow(query, newItem.Name, newItem.Itn, newItem.Passport, newItem.Snils, newItem.Phone, newItem.AccountId).Scan(&newItem.ID)
	return err
}

func (ir *Repository) Update(id int, name string, itn string, passport string, snils string, phone string, accountId int) error {
	query := "UPDATE owner SET name = $1, itn = $2, passport = $3, snils = $4, phone = $5, account_id = $6 WHERE id = $7"
	_, err := ir.DB.Connection.Exec(query, name, itn, passport, snils, phone, accountId, id)
	return err
}

func (ir *Repository) Delete(id int) error {
	query := "DELETE FROM owner WHERE id = $1"
	_, err := ir.DB.Connection.Exec(query, id)
	return err
}
