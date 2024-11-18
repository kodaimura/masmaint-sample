
package product

type Product struct {
	Id int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Description *string `db:"description" json:"description"`
	Price float64 `db:"price" json:"price"`
	StockQuantity int `db:"stock_quantity" json:"stock_quantity"`
	CategoryId int `db:"category_id" json:"category_id"`
	CreatedAt *string `db:"created_at" json:"created_at"`
	UpdatedAt *string `db:"updated_at" json:"updated_at"`
}
