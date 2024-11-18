
package customer

type Customer struct {
	Id int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
	Phone *string `db:"phone" json:"phone"`
	Address *string `db:"address" json:"address"`
	IsActive int `db:"is_active" json:"is_active"`
	CreatedAt *string `db:"created_at" json:"created_at"`
	UpdatedAt *string `db:"updated_at" json:"updated_at"`
}
