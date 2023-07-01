package dao

import (
	"database/sql"

	"masmaint/core/db"
	"masmaint/model/entity"
)


type DepartmentDao struct {
	db *sql.DB
}

func NewDepartmentDao() *DepartmentDao {
	db := db.GetDB()
	return &DepartmentDao{db}
}


func (rep *DepartmentDao) SelectAll() ([]entity.Department, error) {
	var ret []entity.Department

	rows, err := rep.db.Query(
		`SELECT
			code,
			name,
			description,
			manager_id,
			location,
			budget,
			created_at,
			updated_at
		 FROM department
		 ORDER BY code ASC `,
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


func (rep *DepartmentDao) Select(d *entity.Department) (entity.Department, error) {
	var ret entity.Department

	err := rep.db.QueryRow(
		`SELECT
			code,
			name,
			description,
			manager_id,
			location,
			budget,
			created_at,
			updated_at
		 FROM department
		 WHERE code = $1`,
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


func (rep *DepartmentDao) Insert(d *entity.Department) (entity.Department, error) {
	var ret entity.Department

	err := rep.db.QueryRow(
		`INSERT INTO department (
			code,
			name,
			description,
			manager_id,
			location,
			budget
		 ) VALUES($1,$2,$3,$4,$5,$6) 
		 RETURNING
		 	code,
			name,
			description,
			manager_id,
			location,
			budget,
			created_at,
			updated_at`,
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


func (rep *DepartmentDao) Update(d *entity.Department) (entity.Department, error) {
	var ret entity.Department

	err := rep.db.QueryRow(
		`UPDATE department
		 SET
			name = $1,
			description = $2,
			manager_id = $3,
			location = $4,
			budget = $5
		 WHERE code = $6
		 RETURNING 
		 	code,
			name,
			description,
			manager_id,
			location,
			budget,
			created_at,
			updated_at`,
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


func (rep *DepartmentDao) Delete(d *entity.Department) error {
	_, err := rep.db.Exec(
		`DELETE FROM department
		 WHERE code = $1`,
		d.Code,
	)

	return err
}
