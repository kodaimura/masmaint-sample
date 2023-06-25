package dto


type Employee struct {
	Id int `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName string `db:"last_name" json:"last_name"`
	Email string `db:"email" json:"email"`
	PhoneNumber string `db:"phone_number" json:"phone_number"`
	Address string `db:"address" json:"address"`
	HireDate string `db:"hire_date" json:"hire_date"`
	JobTitle string `db:"job_title" json:"job_title"`
	DepartmentId int `db:"department_id" json:"department_id"`
	Salary int `db:"salary" json:"salary"`
}