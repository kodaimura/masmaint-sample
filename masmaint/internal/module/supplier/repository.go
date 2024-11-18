package supplier

import (
	"database/sql"
	"masmaint/internal/core/db"
)


type Repository interface {
	Get(s *Supplier) ([]Supplier, error)
	GetOne(s *Supplier) (Supplier, error)
	Insert(s *Supplier, tx *sql.Tx) (int, error)
	Update(s *Supplier, tx *sql.Tx) error
	Delete(s *Supplier, tx *sql.Tx) error
}


type repository struct {
	db *sql.DB
}

func NewRepository() Repository {
	db := db.GetDB()
	return &repository{db}
}


func (rep *repository) Get(s *Supplier) ([]Supplier, error) {
	where, binds := db.BuildWhereClause(s)
	query := 
	`SELECT
		id
		,name
		,contact_person
		,phone
		,email
		,address
		,is_active
		,created_at
		,updated_at
	 FROM supplier ` + where
	rows, err := rep.db.Query(query, binds...)
	defer rows.Close()

	if err != nil {
		return []Supplier{}, err
	}

	ret := []Supplier{}
	for rows.Next() {
		s := Supplier{}
		err = rows.Scan(
			&s.Id,
			&s.Name,
			&s.ContactPerson,
			&s.Phone,
			&s.Email,
			&s.Address,
			&s.IsActive,
			&s.CreatedAt,
			&s.UpdatedAt,
		)
		if err != nil {
			return []Supplier{}, err
		}
		ret = append(ret, s)
	}

	return ret, nil
}


func (rep *repository) GetOne(s *Supplier) (Supplier, error) {
	var ret Supplier
	where, binds := db.BuildWhereClause(s)
	query := 
	`SELECT
		id
		,name
		,contact_person
		,phone
		,email
		,address
		,is_active
		,created_at
		,updated_at
	 FROM supplier ` + where

	err := rep.db.QueryRow(query, binds...).Scan(
		&ret.Id,
		&ret.Name,
		&ret.ContactPerson,
		&ret.Phone,
		&ret.Email,
		&ret.Address,
		&ret.IsActive,
		&ret.CreatedAt,
		&ret.UpdatedAt,
	)

	return ret, err
}


func (rep *repository) Insert(s *Supplier, tx *sql.Tx) (int, error) {
	cmd := 
	`INSERT INTO supplier (
		name
		,contact_person
		,phone
		,email
		,address
		,is_active
	 ) VALUES(?,?,?,?,?,?)
	 RETURNING id`

	binds := []interface{}{
		s.Name,
		s.ContactPerson,
		s.Phone,
		s.Email,
		s.Address,
		s.IsActive,
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


func (rep *repository) Update(s *Supplier, tx *sql.Tx) error {
	cmd := 
	`UPDATE supplier
	 SET name = ?
		,contact_person = ?
		,phone = ?
		,email = ?
		,address = ?
		,is_active = ?
	 WHERE id = ?`
	binds := []interface{}{
		s.Name,
		s.ContactPerson,
		s.Phone,
		s.Email,
		s.Address,
		s.IsActive,
		s.Id,
	}

	var err error
	if tx != nil {
		_, err = tx.Exec(cmd, binds...)
	} else {
		_, err = rep.db.Exec(cmd, binds...)
	}

	return err
}


func (rep *repository) Delete(s *Supplier, tx *sql.Tx) error {
	where, binds := db.BuildWhereClause(s)
	cmd := "DELETE FROM supplier " + where

	var err error
	if tx != nil {
		_, err = tx.Exec(cmd, binds...)
	} else {
		_, err = rep.db.Exec(cmd, binds...)
	}

	return err
}