package dto


type EmployeeDto struct {
	Id any `json:"id"`
	FirstName any `json:"first_name"`
	LastName any `json:"last_name"`
	Email any `json:"email"`
	PhoneNumber any `json:"phone_number"`
	Address any `json:"address"`
	HireDate any `json:"hire_date"`
	JobTitle any `json:"job_title"`
	DepartmentCode any `json:"department_code"`
	Salary any `json:"salary"`
}