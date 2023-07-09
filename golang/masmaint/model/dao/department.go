package dao

import (
	"database/sql"

	"masmaint/core/db"
	"masmaint/model/entity"
)


type departmentDao struct {
	db *sql.DB
}

func NewDepartmentDao() *departmentDao {
	db := db.GetDB()
	return &departmentDao{db}
}


func (rep *departmentDao) SelectAll() ([]entity.Department, error) {
	var ret []entity.Department

	rows, err := rep.db.Query(
		`SELECT
			code
			,name
			,description
			,manager_id
			,location
			,budget
			,created_at
			,updated_at
		 FROM department
		 ORDER BY code ASC`,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		d := entity.Department{}
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
			break
		}
		ret = append(ret, d)
	}

	return ret, err
}


func (rep *departmentDao) Select(d *entity.Department) (entity.Department, error) {
	var ret entity.Department

	err := rep.db.QueryRow(
		`SELECT
			code
			,name
			,description
			,manager_id
			,location
			,budget
			,created_at
			,updated_at
		 FROM department
		 WHERE code = ?`,
		d.Code,
	).Scan(
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


func (rep *departmentDao) Insert(d *entity.Department) (entity.Department, error) {
	var ret entity.Department

	err := rep.db.QueryRow(
		`INSERT INTO department (
			code
			,name
			,description
			,manager_id
			,location
			,budget
		 ) VALUES (?,?,?,?,?,?)
		 RETURNING
			code
			,name
			,description
			,manager_id
			,location
			,budget
			,created_at
			,updated_at`,
		d.Code,
		d.Name,
		d.Description,
		d.ManagerId,
		d.Location,
		d.Budget,
	).Scan(
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


func (rep *departmentDao) Update(d *entity.Department) (entity.Department, error) {
	var ret entity.Department

	err := rep.db.QueryRow(
		`UPDATE department
		 SET
			name = ?
			,description = ?
			,manager_id = ?
			,location = ?
			,budget = ?
		 WHERE code = ?
		 RETURNING 
			code
			,name
			,description
			,manager_id
			,location
			,budget
			,created_at
			,updated_at`,
		d.Name,
		d.Description,
		d.ManagerId,
		d.Location,
		d.Budget,
		d.Code,
	).Scan(
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


func (rep *departmentDao) Delete(d *entity.Department) error {
	_, err := rep.db.Exec(
		`DELETE FROM department
		 WHERE code = ?`,
		d.Code,
	)

	return err
}
