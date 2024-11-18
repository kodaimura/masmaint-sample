
package product_category

type PostBody struct {
	Name string `json:"name" binding:"required"`
	ParentCategoryId *int `json:"parent_category_id"`
	Description *string `json:"description"`
	SortOrder int `json:"sort_order" binding:"required"`
	IsActive int `json:"is_active" binding:"required"`
}

type PutBody struct {
	Id int `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	ParentCategoryId *int `json:"parent_category_id"`
	Description *string `json:"description"`
	SortOrder int `json:"sort_order" binding:"required"`
	IsActive int `json:"is_active" binding:"required"`
}

type DeleteBody struct {
	Id int `json:"id" binding:"required"`
}
