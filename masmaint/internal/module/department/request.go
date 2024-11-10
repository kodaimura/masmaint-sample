package department


type PostBody struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required`
	Description *string `json:"description"`
	ManagerId *int `json:"manager_id"`
	Location *string `json:"location"`
	Budget *float64 `json:"budget"`
}

type PutBody struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
	Description *string `json:"description"`
	ManagerId *int `json:"manager_id"`
	Location *string `json:"location"`
	Budget *float64 `json:"budget"`
}

type DeleteBody struct {
	Code string `json:"code" binding:"required"`
}