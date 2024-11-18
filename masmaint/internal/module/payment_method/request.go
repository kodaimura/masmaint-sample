
package payment_method

type PostBody struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
	Description *string `json:"description"`
	IsActive int `json:"is_active" binding:"required"`
}

type PutBody struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
	Description *string `json:"description"`
	IsActive int `json:"is_active" binding:"required"`
}

type DeleteBody struct {
	Code string `json:"code" binding:"required"`
}
