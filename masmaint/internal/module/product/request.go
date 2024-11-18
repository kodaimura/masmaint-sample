
package product

type PostBody struct {
	Name string `json:"name" binding:"required"`
	Description *string `json:"description"`
	Price float64 `json:"price" binding:"required"`
	StockQuantity int `json:"stock_quantity" binding:"required"`
	CategoryId int `json:"category_id" binding:"required"`
}

type PutBody struct {
	Id int `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Description *string `json:"description"`
	Price float64 `json:"price" binding:"required"`
	StockQuantity int `json:"stock_quantity" binding:"required"`
	CategoryId int `json:"category_id" binding:"required"`
}

type DeleteBody struct {
	Id int `json:"id" binding:"required"`
}
