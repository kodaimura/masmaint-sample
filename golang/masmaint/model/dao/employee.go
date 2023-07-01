package dao

import (
	"database/sql"

	"masmaint/core/db"
	"masmaint/model/entity"
)


type EmployeeDao struct {
	db *sql.DB
}


func NewEmployeeDao() *EmployeeDao {
	db := db.GetDB()
	return &EmployeeDao{db}
}


func (rep *EmployeeDao) SelectAll() ([]entity.Employee, error) {
	var ret []entity.Employee

	rows, err := rep.db.Query(
		`SELECT
			id,
			first_name,
			last_name,
			email,
			phone_number,
			address,
			hire_date,
			job_title,
			department_id,
			salary
		 FROM employee`,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		e := entity.Employee{}
		err = rows.Scan(
			&e.Id,
			&e.FirstName,
			&e.LastName,
			&e.Email,
			&e.PhoneNumber,
			&e.Address,
			&e.HireDate,
			&e.JobTitle,
			&e.DepartmentId,
			&e.Salary,
		)
		if err != nil {
			break
		}
		ret = append(ret, e)
	}

	return ret, err
}


func (rep *EmployeeDao) Select(e *entity.Employee) (entity.Employee, error) {
	var ret entity.Employee

	err := rep.db.QueryRow(
		`SELECT
			id,
			first_name,
			last_name,
			email,
			phone_number,
			address,
			hire_date,
			job_title,
			department_id,
			salary
		 FROM employee
		 WHERE id = $1`,
		e.Id,
	).Scan(
		&ret.Id,
		&ret.FirstName,
		&ret.LastName,
		&ret.Email,
		&ret.PhoneNumber,
		&ret.Address,
		&ret.HireDate,
		&ret.JobTitle,
		&ret.DepartmentId,
		&ret.Salary,
	)

	return ret, err
}


func (rep *EmployeeDao) Insert(e *entity.Employee) (entity.Employee, error) {
	var ret entity.Employee

	err := rep.db.QueryRow(
		`INSERT INTO employee (
			first_name,
			last_name,
			email,
			phone_number,
			address,
			hire_date,
			job_title,
			department_id,
			salary
		 ) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9) 
		 RETURNING
		 	id,
			first_name,
			last_name,
			email,
			phone_number,
			address,
			hire_date,
			job_title,
			department_id,
			salary`,
		e.FirstName,
		e.LastName,
		e.Email,
		e.PhoneNumber,
		e.Address,
		e.HireDate,
		e.JobTitle,
		e.DepartmentId,
		e.Salary,
	).Scan(
		&ret.Id,
		&ret.FirstName,
		&ret.LastName,
		&ret.Email,
		&ret.PhoneNumber,
		&ret.Address,
		&ret.HireDate,
		&ret.JobTitle,
		&ret.DepartmentId,
		&ret.Salary,
	)

	return ret, err
}


func (rep *EmployeeDao) Update(e *entity.Employee) (entity.Employee, error) {
	var ret entity.Employee
	
	err := rep.db.QueryRow(
		`UPDATE employee
		 SET
			first_name = $1,
			last_name = $2,
			email = $3,
			phone_number = $4,
			address = $5,
			hire_date = $6,
			job_title = $7,
			department_id = $8,
			salary = $9
		 WHERE id = $10
		 RETURNING
		 	id,
			first_name,
			last_name,
			email,
			phone_number,
			address,
			hire_date,
			job_title,
			department_id,
			salary`,
		e.FirstName,
		e.LastName,
		e.Email,
		e.PhoneNumber,
		e.Address,
		e.HireDate,
		e.JobTitle,
		e.DepartmentId,
		e.Salary,
		e.Id,
	).Scan(
		&ret.Id,
		&ret.FirstName,
		&ret.LastName,
		&ret.Email,
		&ret.PhoneNumber,
		&ret.Address,
		&ret.HireDate,
		&ret.JobTitle,
		&ret.DepartmentId,
		&ret.Salary,
	)

	return ret, err
}


func (rep *EmployeeDao) Delete(e *entity.Employee) error {
	_, err := rep.db.Exec(
		`DELETE FROM employee
		 WHERE id = $1`,
		e.Id,
	)

	return err
}
