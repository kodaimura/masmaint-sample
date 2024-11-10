package department

import (
	"database/sql"
	"masmaint/internal/core/db"
)

type Repository interface {
	Get(d *Department) ([]Department, error)
	GetOne(d *Department) (Department, error)
	Insert(d *Department, tx *sql.Tx) error
	Update(d *Department, tx *sql.Tx) error
	Delete(d *Department, tx *sql.Tx) error
}

type repository struct {
	db *sql.DB
}

func NewRepository() Repository {
	db := db.GetDB()
	return &repository{db}
}


func (rep *repository) Get(d *Department) ([]Department, error) {
	where, binds := db.BuildWhereClause(d)
	query := 
	`SELECT
		code
		,name
		,description
		,manager_id
		,location
		,budget
		,created_at
		,updated_at
	 FROM department ` + where
	rows, err := rep.db.Query(query, binds...)
	defer rows.Close()

	if err != nil {
		return []Department{}, err
	}

	ret := []Department{}
	for rows.Next() {
		d := Department{}
		err = rows.Scan(
			&d.Code,
			&d.Name,
			&d.Description,
			&d.ManagerId,
			&d.Location,
			&d.Budget,
			&d.CreatedAt,
			&d.UpdatedAt,
		)
		if err != nil {
			return []Department{}, err
		}
		ret = append(ret, d)
	}

	return ret, nil
}


func (rep *repository) GetOne(d *Department) (Department, error) {
	var ret Department
	where, binds := db.BuildWhereClause(d)
	query := 
	`SELECT
		code
		,name
		,description
		,manager_id
		,location
		,budget
		,created_at
		,updated_at
	 FROM department ` + where

	err := rep.db.QueryRow(query, binds...).Scan(
		&ret.Code,
		&ret.Name,
		&ret.Description,
		&ret.ManagerId,
		&ret.Location,
		&ret.Budget,
		&ret.CreatedAt,
		&ret.UpdatedAt,
	)

	return ret, err
}


func (rep *repository) Insert(d *Department, tx *sql.Tx) error {
	cmd := 
	`INSERT INTO department (
		code
		,name
		,description
		,manager_id
		,location
		,budget
	 ) VALUES(?,?,?,?,?,?)`

	binds := []interface{}{
		d.Code,
		d.Name,
		d.Description,
		d.ManagerId,
		d.Location,
		d.Budget,
	}

	var err error
	if tx != nil {
		_, err = tx.Exec(cmd, binds...)
	} else {
		_, err = rep.db.Exec(cmd, binds...)
	}

	return err
}


func (rep *repository) Update(d *Department, tx *sql.Tx) error {
	cmd := 
	`UPDATE department
	 SET code = ?
		,name = ?
		,description = ?
		,manager_id = ?
		,location = ?
		,budget = ?
	 WHERE code = ?`
	binds := []interface{}{
		d.Code,
		d.Name,
		d.Description,
		d.ManagerId,
		d.Location,
		d.Budget,
		d.Code,
	}

	var err error
	if tx != nil {
        _, err = tx.Exec(cmd, binds...)
    } else {
        _, err = rep.db.Exec(cmd, binds...)
    }

	return err
}


func (rep *repository) Delete(d *Department, tx *sql.Tx) error {
	where, binds := db.BuildWhereClause(d)
	cmd := "DELETE FROM department " + where

	var err error
	if tx != nil {
        _, err = tx.Exec(cmd, binds...)
    } else {
        _, err = rep.db.Exec(cmd, binds...)
    }

	return err
}