package product_category

import (
	"database/sql"
	"masmaint/internal/core/db"
)


type Repository interface {
	Get(pc *ProductCategory) ([]ProductCategory, error)
	GetOne(pc *ProductCategory) (ProductCategory, error)
	Insert(pc *ProductCategory, tx *sql.Tx) (int, error)
	Update(pc *ProductCategory, tx *sql.Tx) error
	Delete(pc *ProductCategory, tx *sql.Tx) error
}


type repository struct {
	db *sql.DB
}

func NewRepository() Repository {
	db := db.GetDB()
	return &repository{db}
}


func (rep *repository) Get(pc *ProductCategory) ([]ProductCategory, error) {
	where, binds := db.BuildWhereClause(pc)
	query := 
	`SELECT
		id
		,name
		,parent_category_id
		,description
		,sort_order
		,is_active
		,created_at
		,updated_at
	 FROM product_category ` + where
	rows, err := rep.db.Query(query, binds...)
	defer rows.Close()

	if err != nil {
		return []ProductCategory{}, err
	}

	ret := []ProductCategory{}
	for rows.Next() {
		pc := ProductCategory{}
		err = rows.Scan(
			&pc.Id,
			&pc.Name,
			&pc.ParentCategoryId,
			&pc.Description,
			&pc.SortOrder,
			&pc.IsActive,
			&pc.CreatedAt,
			&pc.UpdatedAt,
		)
		if err != nil {
			return []ProductCategory{}, err
		}
		ret = append(ret, pc)
	}

	return ret, nil
}


func (rep *repository) GetOne(pc *ProductCategory) (ProductCategory, error) {
	var ret ProductCategory
	where, binds := db.BuildWhereClause(pc)
	query := 
	`SELECT
		id
		,name
		,parent_category_id
		,description
		,sort_order
		,is_active
		,created_at
		,updated_at
	 FROM product_category ` + where

	err := rep.db.QueryRow(query, binds...).Scan(
		&ret.Id,
		&ret.Name,
		&ret.ParentCategoryId,
		&ret.Description,
		&ret.SortOrder,
		&ret.IsActive,
		&ret.CreatedAt,
		&ret.UpdatedAt,
	)

	return ret, err
}


func (rep *repository) Insert(pc *ProductCategory, tx *sql.Tx) (int, error) {
	cmd := 
	`INSERT INTO product_category (
		name
		,parent_category_id
		,description
		,sort_order
		,is_active
	 ) VALUES(?,?,?,?,?)
	 RETURNING id`

	binds := []interface{}{
		pc.Name,
		pc.ParentCategoryId,
		pc.Description,
		pc.SortOrder,
		pc.IsActive,
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


func (rep *repository) Update(pc *ProductCategory, tx *sql.Tx) error {
	cmd := 
	`UPDATE product_category
	 SET name = ?
		,parent_category_id = ?
		,description = ?
		,sort_order = ?
		,is_active = ?
	 WHERE id = ?`
	binds := []interface{}{
		pc.Name,
		pc.ParentCategoryId,
		pc.Description,
		pc.SortOrder,
		pc.IsActive,
		pc.Id,
	}

	var err error
	if tx != nil {
		_, err = tx.Exec(cmd, binds...)
	} else {
		_, err = rep.db.Exec(cmd, binds...)
	}

	return err
}


func (rep *repository) Delete(pc *ProductCategory, tx *sql.Tx) error {
	where, binds := db.BuildWhereClause(pc)
	cmd := "DELETE FROM product_category " + where

	var err error
	if tx != nil {
		_, err = tx.Exec(cmd, binds...)
	} else {
		_, err = rep.db.Exec(cmd, binds...)
	}

	return err
}