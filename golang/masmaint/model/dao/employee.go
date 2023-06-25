package dao


import (
	"database/sql"

	"masmaint/core/db"
	"masmaint/model/entity"
)


type EmployeeDao struct {
	db *sql.DB
}


func NewEmployeeDao() EmployeeDao {
	db := db.GetDB()
	return &EmployeeDao{db}
}


func (rep *employeeDao) Insert(e *entity.Employee) error {
	_, err := rep.db.Exec(
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
		 ) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)`,
		e.FirstName,
		e.LastName,
		e.Email,
		e.PhoneNumber,
		e.Address,
		e.HireDate,
		e.JobTitle,
		e.DepartmentId,
		e.Salary,
	)

	return err
}


func (rep *employeeDao) Select(id int) (entity.Employee, error) {
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
		id,
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


func (rep *employeeDao) Update(e *entity.Employee) error {
	_, err := rep.db.Exec(
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
		 WHERE id = $10`,
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
	)

	return err
}


func (rep *employeeDao) Delete(id int) error {
	_, err := rep.db.Exec(
		`DELETE FROM employee
		 WHERE id = $1`,
		id,
	)

	return err
}


func (rep *employeeDao) SelectAll() ([]entity.Employee, error) {
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
			&e.Salary
		)
		if err != nil {
			break
		}
		ret = append(ret, e)
	}

	return ret, err
}