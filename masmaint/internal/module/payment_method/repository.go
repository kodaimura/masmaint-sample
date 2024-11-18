package payment_method

import (
	"database/sql"
	"masmaint/internal/core/db"
)


type Repository interface {
	Get(pm *PaymentMethod) ([]PaymentMethod, error)
	GetOne(pm *PaymentMethod) (PaymentMethod, error)
	Insert(pm *PaymentMethod, tx *sql.Tx) error
	Update(pm *PaymentMethod, tx *sql.Tx) error
	Delete(pm *PaymentMethod, tx *sql.Tx) error
}


type repository struct {
	db *sql.DB
}

func NewRepository() Repository {
	db := db.GetDB()
	return &repository{db}
}


func (rep *repository) Get(pm *PaymentMethod) ([]PaymentMethod, error) {
	where, binds := db.BuildWhereClause(pm)
	query := 
	`SELECT
		code
		,name
		,description
		,is_active
		,created_at
		,updated_at
	 FROM payment_method ` + where
	rows, err := rep.db.Query(query, binds...)
	defer rows.Close()

	if err != nil {
		return []PaymentMethod{}, err
	}

	ret := []PaymentMethod{}
	for rows.Next() {
		pm := PaymentMethod{}
		err = rows.Scan(
			&pm.Code,
			&pm.Name,
			&pm.Description,
			&pm.IsActive,
			&pm.CreatedAt,
			&pm.UpdatedAt,
		)
		if err != nil {
			return []PaymentMethod{}, err
		}
		ret = append(ret, pm)
	}

	return ret, nil
}


func (rep *repository) GetOne(pm *PaymentMethod) (PaymentMethod, error) {
	var ret PaymentMethod
	where, binds := db.BuildWhereClause(pm)
	query := 
	`SELECT
		code
		,name
		,description
		,is_active
		,created_at
		,updated_at
	 FROM payment_method ` + where

	err := rep.db.QueryRow(query, binds...).Scan(
		&ret.Code,
		&ret.Name,
		&ret.Description,
		&ret.IsActive,
		&ret.CreatedAt,
		&ret.UpdatedAt,
	)

	return ret, err
}


func (rep *repository) Insert(pm *PaymentMethod, tx *sql.Tx) error {
	cmd := 
	`INSERT INTO payment_method (
		code
		,name
		,description
		,is_active
	 ) VALUES(?,?,?,?)`

	binds := []interface{}{
		pm.Code,
		pm.Name,
		pm.Description,
		pm.IsActive,
	}

	var err error
	if tx != nil {
		_, err = tx.Exec(cmd, binds...)
	} else {
		_, err = rep.db.Exec(cmd, binds...)
	}

	return err
}


func (rep *repository) Update(pm *PaymentMethod, tx *sql.Tx) error {
	cmd := 
	`UPDATE payment_method
	 SET name = ?
		,description = ?
		,is_active = ?
	 WHERE code = ?`
	binds := []interface{}{
		pm.Name,
		pm.Description,
		pm.IsActive,
		pm.Code,
	}

	var err error
	if tx != nil {
		_, err = tx.Exec(cmd, binds...)
	} else {
		_, err = rep.db.Exec(cmd, binds...)
	}

	return err
}


func (rep *repository) Delete(pm *PaymentMethod, tx *sql.Tx) error {
	where, binds := db.BuildWhereClause(pm)
	cmd := "DELETE FROM payment_method " + where

	var err error
	if tx != nil {
		_, err = tx.Exec(cmd, binds...)
	} else {
		_, err = rep.db.Exec(cmd, binds...)
	}

	return err
}