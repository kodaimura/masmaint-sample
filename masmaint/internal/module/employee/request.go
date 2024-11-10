package employee


type PostBody struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName *string `json:"last_name"`
	Email *string `json:"email"`
	PhoneNumber *string `json:"phone_number"`
	Address *string `json:"address"`
	HireDate *string `json:"hire_date"`
	JobTitle *string `json:"job_title"`
	DepartmentCode *string `json:"department_code"`
	Salary *float64 `json:"salary"`
}

type PutBody struct {
	Id int `json:"id" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName *string `json:"last_name"`
	Email *string `json:"email"`
	PhoneNumber *string `json:"phone_number"`
	Address *string `json:"address"`
	HireDate *string `json:"hire_date"`
	JobTitle *string `json:"job_title"`
	DepartmentCode *string `json:"department_code"`
	Salary *float64 `json:"salary"`
}

type DeleteBody struct {
	Id int `json:"id" binding:"required"`
}