package employee

import (
	"database/sql"
	"masmaint/internal/core/db"
)

type Repository interface {
	Get(e *Employee) ([]Employee, error)
	GetOne(e *Employee) (Employee, error)
	Insert(e *Employee, tx *sql.Tx) (int, error)
	Update(e *Employee, tx *sql.Tx) error
	Delete(e *Employee, tx *sql.Tx) error
}

type repository struct {
	db *sql.DB
}

func NewRepository() Repository {
	db := db.GetDB()
	return &repository{db}
}


func (rep *repository) Get(e *Employee) ([]Employee, error) {
	where, binds := db.BuildWhereClause(e)
	query := 
	`SELECT
		id
		,first_name
		,last_name
		,email
		,phone_number
		,address
		,hire_date
		,job_title
		,department_code
		,salary
		,created_at
		,updated_at
	 FROM employee ` + where
	rows, err := rep.db.Query(query, binds...)
	defer rows.Close()

	if err != nil {
		return []Employee{}, err
	}

	ret := []Employee{}
	for rows.Next() {
		e := Employee{}
		err = rows.Scan(
			&e.Id,
			&e.FirstName,
			&e.LastName,
			&e.Email,
			&e.PhoneNumber,
			&e.Address,
			&e.HireDate,
			&e.JobTitle,
			&e.DepartmentCode,
			&e.Salary,
			&e.CreatedAt,
			&e.UpdatedAt,
		)
		if err != nil {
			return []Employee{}, err
		}
		ret = append(ret, e)
	}

	return ret, nil
}


func (rep *repository) GetOne(e *Employee) (Employee, error) {
	var ret Employee
	where, binds := db.BuildWhereClause(e)
	query := 
	`SELECT
		id
		,first_name
		,last_name
		,email
		,phone_number
		,address
		,hire_date
		,job_title
		,department_code
		,salary
		,created_at
		,updated_at
	 FROM employee ` + where

	err := rep.db.QueryRow(query, binds...).Scan(
		&ret.Id,
		&ret.FirstName,
		&ret.LastName,
		&ret.Email,
		&ret.PhoneNumber,
		&ret.Address,
		&ret.HireDate,
		&ret.JobTitle,
		&ret.DepartmentCode,
		&ret.Salary,
		&ret.CreatedAt,
		&ret.UpdatedAt,
	)

	return ret, err
}


func (rep *repository) Insert(e *Employee, tx *sql.Tx) (int, error) {
	cmd := 
	`INSERT INTO employee (
		first_name
		,last_name
		,email
		,phone_number
		,address
		,hire_date
		,job_title
		,department_code
		,salary
	 ) VALUES(?,?,?,?,?,?,?,?,?)
	 RETURNING id`

	binds := []interface{}{
		e.FirstName,
		e.LastName,
		e.Email,
		e.PhoneNumber,
		e.Address,
		e.HireDate,
		e.JobTitle,
		e.DepartmentCode,
		e.Salary,
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


func (rep *repository) Update(e *Employee, tx *sql.Tx) error {
	cmd := 
	`UPDATE employee
	 SET first_name = ?
		,last_name = ?
		,email = ?
		,phone_number = ?
		,address = ?
		,hire_date = ?
		,job_title = ?
		,department_code = ?
		,salary = ?
	 WHERE id = ?`
	binds := []interface{}{
		e.FirstName,
		e.LastName,
		e.Email,
		e.PhoneNumber,
		e.Address,
		e.HireDate,
		e.JobTitle,
		e.DepartmentCode,
		e.Salary,
		e.Id,
	}

	var err error
	if tx != nil {
        _, err = tx.Exec(cmd, binds...)
    } else {
        _, err = rep.db.Exec(cmd, binds...)
    }

	return err
}


func (rep *repository) Delete(e *Employee, tx *sql.Tx) error {
	where, binds := db.BuildWhereClause(e)
	cmd := "DELETE FROM employee " + where

	var err error
	if tx != nil {
        _, err = tx.Exec(cmd, binds...)
    } else {
        _, err = rep.db.Exec(cmd, binds...)
    }

	return err
}