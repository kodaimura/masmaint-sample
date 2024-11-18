package customer

import (
	"database/sql"
	"masmaint/internal/core/db"
)


type Repository interface {
	Get(c *Customer) ([]Customer, error)
	GetOne(c *Customer) (Customer, error)
	Insert(c *Customer, tx *sql.Tx) (int, error)
	Update(c *Customer, tx *sql.Tx) error
	Delete(c *Customer, tx *sql.Tx) error
}


type repository struct {
	db *sql.DB
}

func NewRepository() Repository {
	db := db.GetDB()
	return &repository{db}
}


func (rep *repository) Get(c *Customer) ([]Customer, error) {
	where, binds := db.BuildWhereClause(c)
	query := 
	`SELECT
		id
		,name
		,email
		,phone
		,address
		,is_active
		,created_at
		,updated_at
	 FROM customer ` + where
	rows, err := rep.db.Query(query, binds...)
	defer rows.Close()

	if err != nil {
		return []Customer{}, err
	}

	ret := []Customer{}
	for rows.Next() {
		c := Customer{}
		err = rows.Scan(
			&c.Id,
			&c.Name,
			&c.Email,
			&c.Phone,
			&c.Address,
			&c.IsActive,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return []Customer{}, err
		}
		ret = append(ret, c)
	}

	return ret, nil
}


func (rep *repository) GetOne(c *Customer) (Customer, error) {
	var ret Customer
	where, binds := db.BuildWhereClause(c)
	query := 
	`SELECT
		id
		,name
		,email
		,phone
		,address
		,is_active
		,created_at
		,updated_at
	 FROM customer ` + where

	err := rep.db.QueryRow(query, binds...).Scan(
		&ret.Id,
		&ret.Name,
		&ret.Email,
		&ret.Phone,
		&ret.Address,
		&ret.IsActive,
		&ret.CreatedAt,
		&ret.UpdatedAt,
	)

	return ret, err
}


func (rep *repository) Insert(c *Customer, tx *sql.Tx) (int, error) {
	cmd := 
	`INSERT INTO customer (
		name
		,email
		,phone
		,address
		,is_active
	 ) VALUES(?,?,?,?,?)
	 RETURNING id`

	binds := []interface{}{
		c.Name,
		c.Email,
		c.Phone,
		c.Address,
		c.IsActive,
	}

	var id int
	var err error
	if tx != nil {
		err = tx.QueryRow(cmd, binds...).Scan(&id)
	} else {
		err = rep.db.QueryRow(cmd, binds...).Scan(&id)
	}

	return id, err
}


func (rep *repository) Update(c *Customer, tx *sql.Tx) error {
	cmd := 
	`UPDATE customer
	 SET name = ?
		,email = ?
		,phone = ?
		,address = ?
		,is_active = ?
	 WHERE id = ?`
	binds := []interface{}{
		c.Name,
		c.Email,
		c.Phone,
		c.Address,
		c.IsActive,
		c.Id,
	}

	var err error
	if tx != nil {
		_, err = tx.Exec(cmd, binds...)
	} else {
		_, err = rep.db.Exec(cmd, binds...)
	}

	return err
}


func (rep *repository) Delete(c *Customer, tx *sql.Tx) error {
	where, binds := db.BuildWhereClause(c)
	cmd := "DELETE FROM customer " + where

	var err error
	if tx != nil {
		_, err = tx.Exec(cmd, binds...)
	} else {
		_, err = rep.db.Exec(cmd, binds...)
	}

	return err
}