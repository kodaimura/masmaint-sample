
package payment_method

type PaymentMethod struct {
	Code string `db:"code" json:"code"`
	Name string `db:"name" json:"name"`
	Description *string `db:"description" json:"description"`
	IsActive int `db:"is_active" json:"is_active"`
	CreatedAt *string `db:"created_at" json:"created_at"`
	UpdatedAt *string `db:"updated_at" json:"updated_at"`
}
