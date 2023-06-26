package dto


type EmployeeDto struct {
	Id int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address string `json:"address"`
	HireDate string `json:"hire_date"`
	JobTitle string `json:"job_title"`
	DepartmentId int `json:"department_id"`
	Salary int `json:"salary"`
}