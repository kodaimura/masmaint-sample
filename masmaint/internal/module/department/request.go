package department


type PostBody struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Description *string `json:"description"`
	ManagerId *int `json:"manager_id"`
	Location *string `db:"json:"location"`
	Budget float64 `db:"json:"budget"`
}

type PutBody struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Description *string `json:"description"`
	ManagerId *int `json:"manager_id"`
	Location *string `db:"json:"location"`
	Budget float64 `db:"json:"budget"`
}

type DeleteBody struct {
	Code string `json:"code"`
}