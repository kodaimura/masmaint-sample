
package supplier

type PostBody struct {
	Name string `json:"name" binding:"required"`
	ContactPerson *string `json:"contact_person"`
	Phone *string `json:"phone"`
	Email *string `json:"email"`
	Address *string `json:"address"`
	IsActive int `json:"is_active" binding:"required"`
}

type PutBody struct {
	Id int `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	ContactPerson *string `json:"contact_person"`
	Phone *string `json:"phone"`
	Email *string `json:"email"`
	Address *string `json:"address"`
	IsActive int `json:"is_active" binding:"required"`
}

type DeleteBody struct {
	Id int `json:"id" binding:"required"`
}
