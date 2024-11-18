
package supplier

type Supplier struct {
	Id int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	ContactPerson *string `db:"contact_person" json:"contact_person"`
	Phone *string `db:"phone" json:"phone"`
	Email *string `db:"email" json:"email"`
	Address *string `db:"address" json:"address"`
	IsActive int `db:"is_active" json:"is_active"`
	CreatedAt *string `db:"created_at" json:"created_at"`
	UpdatedAt *string `db:"updated_at" json:"updated_at"`
}
