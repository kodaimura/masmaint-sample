package product

import (
	"database/sql"
	"masmaint/internal/core/db"
)


type Repository interface {
	Get(p *Product) ([]Product, error)
	GetOne(p *Product) (Product, error)
	Insert(p *Product, tx *sql.Tx) (int, error)
	Update(p *Product, tx *sql.Tx) error
	Delete(p *Product, tx *sql.Tx) error
}


type repository struct {
	db *sql.DB
}

func NewRepository() Repository {
	db := db.GetDB()
	return &repository{db}
}


func (rep *repository) Get(p *Product) ([]Product, error) {
	where, binds := db.BuildWhereClause(p)
	query := 
	`SELECT
		id
		,name
		,description
		,price
		,stock_quantity
		,category_id
		,created_at
		,updated_at
	 FROM product ` + where
	rows, err := rep.db.Query(query, binds...)
	defer rows.Close()

	if err != nil {
		return []Product{}, err
	}

	ret := []Product{}
	for rows.Next() {
		p := Product{}
		err = rows.Scan(
			&p.Id,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.StockQuantity,
			&p.CategoryId,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return []Product{}, err
		}
		ret = append(ret, p)
	}

	return ret, nil
}


func (rep *repository) GetOne(p *Product) (Product, error) {
	var ret Product
	where, binds := db.BuildWhereClause(p)
	query := 
	`SELECT
		id
		,name
		,description
		,price
		,stock_quantity
		,category_id
		,created_at
		,updated_at
	 FROM product ` + where

	err := rep.db.QueryRow(query, binds...).Scan(
		&ret.Id,
		&ret.Name,
		&ret.Description,
		&ret.Price,
		&ret.StockQuantity,
		&ret.CategoryId,
		&ret.CreatedAt,
		&ret.UpdatedAt,
	)

	return ret, err
}


func (rep *repository) Insert(p *Product, tx *sql.Tx) (int, error) {
	cmd := 
	`INSERT INTO product (
		name
		,description
		,price
		,stock_quantity
		,category_id
	 ) VALUES(?,?,?,?,?)
	 RETURNING id`

	binds := []interface{}{
		p.Name,
		p.Description,
		p.Price,
		p.StockQuantity,
		p.CategoryId,
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


func (rep *repository) Update(p *Product, tx *sql.Tx) error {
	cmd := 
	`UPDATE product
	 SET name = ?
		,description = ?
		,price = ?
		,stock_quantity = ?
		,category_id = ?
	 WHERE id = ?`
	binds := []interface{}{
		p.Name,
		p.Description,
		p.Price,
		p.StockQuantity,
		p.CategoryId,
		p.Id,
	}

	var err error
	if tx != nil {
		_, err = tx.Exec(cmd, binds...)
	} else {
		_, err = rep.db.Exec(cmd, binds...)
	}

	return err
}


func (rep *repository) Delete(p *Product, tx *sql.Tx) error {
	where, binds := db.BuildWhereClause(p)
	cmd := "DELETE FROM product " + where

	var err error
	if tx != nil {
		_, err = tx.Exec(cmd, binds...)
	} else {
		_, err = rep.db.Exec(cmd, binds...)
	}

	return err
}