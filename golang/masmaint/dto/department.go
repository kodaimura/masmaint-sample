package dto


type DepartmentDto struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	ManagerId int `json:"manager_id"`
	Location string `json:"location"`
	Budget float32 `json:"budget"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}