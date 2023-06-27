package dto


type DepartmentDto struct {
	Id any `json:"id"`
	Name any `json:"name"`
	Description any `json:"description"`
	ManagerId any `json:"manager_id"`
	Location any `json:"location"`
	Budget any `json:"budget"`
	CreatedAt any `json:"created_at"`
	UpdatedAt any `json:"updated_at"`
}