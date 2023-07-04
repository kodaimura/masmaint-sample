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
		 FROM employee
		 ORDER BY id ASC`,
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
			&e.DepartmentCode,
			&e.Salary,
			&e.CreatedAt,
			&e.UpdatedAt,
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
		 FROM employee
		 WHERE id = ?`,
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
		&ret.DepartmentCode,
		&ret.Salary,
		&ret.CreatedAt,
		&ret.UpdatedAt,
	)

	return ret, err
}


func (rep *EmployeeDao) Insert(e *entity.Employee) (entity.Employee, error) {
	var ret entity.Employee

	err := rep.db.QueryRow(
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
		 ) VALUES (?,?,?,?,?,?,?,?,?)
		 RETURNING
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
			,updated_at`,
		e.FirstName,
		e.LastName,
		e.Email,
		e.PhoneNumber,
		e.Address,
		e.HireDate,
		e.JobTitle,
		e.DepartmentCode,
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
		&ret.DepartmentCode,
		&ret.Salary,
		&ret.CreatedAt,
		&ret.UpdatedAt,
	)

	return ret, err
}


func (rep *EmployeeDao) Update(e *entity.Employee) (entity.Employee, error) {
	var ret entity.Employee

	err := rep.db.QueryRow(
		`UPDATE employee
		 SET
			first_name = ?
			,last_name = ?
			,email = ?
			,phone_number = ?
			,address = ?
			,hire_date = ?
			,job_title = ?
			,department_code = ?
			,salary = ?
		 WHERE id = ?
		 RETURNING 
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
			,updated_at`,
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
	).Scan(
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


func (rep *EmployeeDao) Delete(e *entity.Employee) error {
	_, err := rep.db.Exec(
		`DELETE FROM employee
		 WHERE id = ?`,
		e.Id,
	)

	return err
}
