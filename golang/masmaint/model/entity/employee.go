package entity

import (
	"database/sql"

	"masmaint/dto"
	"masmaint/core/utils"
)


type Employee struct {
	Id int64 `db:"id"`
	FirstName string `db:"first_name"`
	LastName sql.NullString `db:"last_name"`
	Email sql.NullString `db:"email"`
	PhoneNumber sql.NullString `db:"phone_number"`
	Address sql.NullString `db:"address"`
	HireDate sql.NullString `db:"hire_date"`
	JobTitle sql.NullString `db:"job_title"`
	DepartmentCode sql.NullString `db:"department_code"`
	Salary sql.NullFloat64 `db:"salary"`
}

func NewEmployee() *Employee {
	return &Employee{}
}


func (e *Employee) SetId(id any) error {
	x, err := utils.ToInt64(id)
	if err != nil {
		return err
	}
	e.Id = x
	return nil
}

func (e *Employee) SetFirstName(firstName any) error {
	e.FirstName = utils.ToString(firstName)
	return nil
}

func (e *Employee) SetLastName(lastName any) error {
	if lastName == nil {
		e.LastName.Valid = false
		return nil
	} 

	e.LastName.String = utils.ToString(lastName)
	e.LastName.Valid = true
	return nil
}

func (e *Employee) SetEmail(email any) error {
	if email == nil {
		e.Email.Valid = false
		return nil
	} 

	e.Email.String = utils.ToString(email)
	e.Email.Valid = true
	return nil
}

func (e *Employee) SetPhoneNumber(phoneNumber any) error {
	if phoneNumber == nil {
		e.PhoneNumber.Valid = false
		return nil
	} 

	e.PhoneNumber.String = utils.ToString(phoneNumber)
	e.PhoneNumber.Valid = true
	return nil
}

func (e *Employee) SetAddress(address any) error {
	if address == nil {
		e.Address.Valid = false
		return nil
	} 

	e.Address.String = utils.ToString(address)
	e.Address.Valid = true
	return nil
}

func (e *Employee) SetHireDate(hireDate any) error {
	//日付型は "" の時は null 扱いとする。
	if hireDate == nil || hireDate == "" {
		e.HireDate.Valid = false
		return nil
	} 

	e.HireDate.String = utils.ToString(hireDate)
	e.HireDate.Valid = true
	return nil
}

func (e *Employee) SetJobTitle(jobTitle any) error {
	if jobTitle == nil {
		e.JobTitle.Valid = false
		return nil
	} 

	e.JobTitle.String = utils.ToString(jobTitle)
	e.JobTitle.Valid = true
	return nil
}

func (e *Employee) SetDepartmentCode(departmentCode any) error {
	if departmentCode == nil {
		e.DepartmentCode.Valid = false
		return nil
	}

	e.DepartmentCode.String = utils.ToString(departmentCode)
	e.DepartmentCode.Valid = true
	return nil
}

func (e *Employee) SetSalary(salary any) error {
	if salary == nil || salary == "" {
		e.Salary.Valid = false
		return nil
	}

	x, err := utils.ToFloat64(salary)
	if err != nil {
		return err
	}
	e.Salary.Float64 = x
	e.Salary.Valid = true
	return nil
}


func (e *Employee) ToEmployeeDto() dto.EmployeeDto {
	var eDto dto.EmployeeDto

	eDto.Id = e.Id
	eDto.FirstName = e.FirstName
	if e.LastName.Valid != false {
		eDto.LastName = e.LastName.String
	}
	if e.Email.Valid != false {
		eDto.Email = e.Email.String
	}
	if e.PhoneNumber.Valid != false {
		eDto.PhoneNumber = e.PhoneNumber.String
	}
	if e.Address.Valid != false {
		eDto.Address = e.Address.String
	}
	if e.HireDate.Valid != false {
		eDto.HireDate = e.HireDate.String
	}
	if e.JobTitle.Valid != false {
		eDto.JobTitle = e.JobTitle.String
	}
	if e.DepartmentCode.Valid != false {
		eDto.DepartmentCode = e.DepartmentCode.String
	}
	if e.Salary.Valid != false {
		eDto.Salary = e.Salary.Float64
	}

	return eDto
}
