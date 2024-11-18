
package customer

type PostBody struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Phone *string `json:"phone"`
	Address *string `json:"address"`
	IsActive int `json:"is_active" binding:"required"`
}

type PutBody struct {
	Id int `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Phone *string `json:"phone"`
	Address *string `json:"address"`
	IsActive int `json:"is_active" binding:"required"`
}

type DeleteBody struct {
	Id int `json:"id" binding:"required"`
}
