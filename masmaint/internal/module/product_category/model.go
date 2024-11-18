
package product_category

type ProductCategory struct {
	Id int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	ParentCategoryId *int `db:"parent_category_id" json:"parent_category_id"`
	Description *string `db:"description" json:"description"`
	SortOrder int `db:"sort_order" json:"sort_order"`
	IsActive int `db:"is_active" json:"is_active"`
	CreatedAt *string `db:"created_at" json:"created_at"`
	UpdatedAt *string `db:"updated_at" json:"updated_at"`
}
