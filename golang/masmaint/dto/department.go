package dto


type DepartmentDto struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	ManagerId int `json:"manager_id"`
	Location string `json:"location"`
	Budget float32 `json:"budget"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}