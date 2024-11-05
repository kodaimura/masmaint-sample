package repository

import (
	"database/sql"

	"masmaint/internal/core/db"
	"masmaint/internal/model"
)


type DepartmentRepository interface {
	Get(d *model.Department) ([]model.Department, error)
	GetOne(d *model.Department) (model.Department, error)
	Insert(d *model.Department, tx *sql.Tx) error
	Update(d *model.Department, tx *sql.Tx) error
	Delete(d *model.Department, tx *sql.Tx) error
}


type departmentRepository struct {
	db *sql.DB
}

func NewDepartmentRepository() DepartmentRepository {
	db := db.GetDB()
	return &departmentRepository{db}
}


func (rep *departmentRepository) Get(d *model.Department) ([]model.Department, error) {
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
		return []model.Department{}, err
	}

	ret := []model.Department{}
	for rows.Next() {
		d := model.Department{}
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
			return []model.Department{}, err
		}
		ret = append(ret, d)
	}

	return ret, nil
}


func (rep *departmentRepository) GetOne(d *model.Department) (model.Department, error) {
	var ret model.Department
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


func (rep *departmentRepository) Insert(d *model.Department, tx *sql.Tx) error {
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


func (rep *departmentRepository) Update(d *model.Department, tx *sql.Tx) error {
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


func (rep *departmentRepository) Delete(d *model.Department, tx *sql.Tx) error {
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