package employee


type Employee struct {
	Id int `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName *string `db:"last_name" json:"last_name"`
	Email *string `db:"email" json:"email"`
	PhoneNumber *string `db:"phone_number" json:"phone_number"`
	Address *string `db:"address" json:"address"`
	HireDate *string `db:"hire_date" json:"hire_date"`
	JobTitle *string `db:"job_title" json:"job_title"`
	DepartmentCode *string `db:"department_code" json:"department_code"`
	Salary *float64 `db:"salary" json:"salary"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}