package dao


import (
	"database/sql"

	"masmaint/core/db"
	"masmaint/model/entity"
)


type DepartmentDao interface {
	Insert(d *entity.Department) error
	Select(id int) (entity.Department, error)
	Update(d *entity.Department) error
	Delete(id int) error
	SelectAll() ([]entity.Department, error)
}


type departmentDao struct {
	db *sql.DB
}


func NewDepartmentDao() DepartmentDao {
	db := db.GetDB()
	return &departmentDao{db}
}


func (rep *departmentDao) Insert(d *entity.Department) error {
	_, err := rep.db.Exec(
		`INSERT INTO department (
			name,
			description,
			manager_id,
			location,
			budget
		 ) VALUES($1,$2,$3,$4,$5)`,
		d.Name,
		d.Description,
		d.ManagerId,
		d.Location,
		d.Budget,
	)

	return err
}


func (rep *departmentDao) Select(id int) (entity.Department, error) {
	var ret entity.Department

	err := rep.db.QueryRow(
		`SELECT
			id,
			name,
			description,
			manager_id,
			location,
			budget,
			create_at,
			update_at
		 FROM department
		 WHERE id = $1`,
		id,
	).Scan(
		&ret.Id,
		&ret.Name,
		&ret.Description,
		&ret.ManagerId,
		&ret.Location,
		&ret.Budget,
		&ret.CreateAt,
		&ret.UpdateAt,
	)

	return ret, err
}


func (rep *departmentDao) Update(d *entity.Department) error {
	_, err := rep.db.Exec(
		`UPDATE department
		 SET
			name = $1,
			description = $2,
			manager_id = $3,
			location = $4,
			budget = $5
		 WHERE id = $6`,
		d.Name,
		d.Description,
		d.ManagerId,
		d.Location,
		d.Budget,
		d.Id,
	)

	return err
}


func (rep *departmentDao) Delete(id int) error {
	_, err := rep.db.Exec(
		`DELETE FROM department
		 WHERE id = $1`,
		id,
	)

	return err
}


func (rep *departmentDao) SelectAll() ([]entity.Department, error) {
	var ret []entity.Department

	rows, err := rep.db.Query(
		`SELECT
			id,
			name,
			description,
			manager_id,
			location,
			budget,
			create_at,
			update_at
		 FROM department`,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		d := entity.Department{}
		err = rows.Scan(
			&d.Id,
			&d.Name,
			&d.Description,
			&d.ManagerId,
			&d.Location,
			&d.Budget,
			&d.CreateAt,
			&d.UpdateAt,
		)
		if err != nil {
			break
		}
		ret = append(ret, d)
	}

	return ret, err
}