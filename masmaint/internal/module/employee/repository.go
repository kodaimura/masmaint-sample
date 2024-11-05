package employee

import (
	"database/sql"

	"masmaint/internal/core/db"
	"masmaint/internal/model"
)


type EmployeeRepository interface {
	Get(e *model.Employee) ([]model.Employee, error)
	GetOne(e *model.Employee) (model.Employee, error)
	Insert(e *model.Employee, tx *sql.Tx) (int, error)
	Update(e *model.Employee, tx *sql.Tx) error
	Delete(e *model.Employee, tx *sql.Tx) error
}


type employeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository() EmployeeRepository {
	db := db.GetDB()
	return &employeeRepository{db}
}


func (rep *employeeRepository) Get(e *model.Employee) ([]model.Employee, error) {
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
		return []model.Employee{}, err
	}

	ret := []model.Employee{}
	for rows.Next() {
		e := model.Employee{}
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
			return []model.Employee{}, err
		}
		ret = append(ret, e)
	}

	return ret, nil
}


func (rep *employeeRepository) GetOne(e *model.Employee) (model.Employee, error) {
	var ret model.Employee
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


func (rep *employeeRepository) Insert(e *model.Employee, tx *sql.Tx) (int, error) {
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


func (rep *employeeRepository) Update(e *model.Employee, tx *sql.Tx) error {
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


func (rep *employeeRepository) Delete(e *model.Employee, tx *sql.Tx) error {
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