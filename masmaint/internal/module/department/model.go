package department


type Department struct {
	Code string `db:"code" json:"code"`
	Name string `db:"name" json:"name"`
	Description *string `db:"description" json:"description"`
	ManagerId *int `db:"manager_id" json:"manager_id"`
	Location *string `db:"location" json:"location"`
	Budget float64 `db:"budget" json:"budget"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}