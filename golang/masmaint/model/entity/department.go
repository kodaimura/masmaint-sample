package entity

import (
	"database/sql"

	"masmaint/dto"
	"masmaint/core/utils"
)


type Department struct {
	Code string `db:"code"`
	Name string `db:"name"`
	Description sql.NullString `db:"description"`
	ManagerId sql.NullInt64 `db:"manager_id"`
	Location sql.NullString `db:"location"`
	Budget float64 `db:"budget"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func NewDepartment() *Department {
	return &Department{}
}

func (e *Department) SetCode(code any) error {
	e.Code = utils.ToString(code)
	return nil
}

func (e *Department) SetName(name any) error {
	e.Name = utils.ToString(name)
	return nil
}

func (e *Department) SetDescription(description any) error {
	if description == nil {
		e.Description.Valid = false
		return nil
	}

	e.Description.String = utils.ToString(description)
	e.Description.Valid = true
	return nil
}

func (e *Department) SetManagerId(managerId any) error {
	if managerId == nil || managerId == "" {
		e.ManagerId.Valid = false
		return nil
	}

	x, err := utils.ToInt64(managerId)
	if err != nil {
		return err
	}
	e.ManagerId.Int64 = x
	e.ManagerId.Valid = true
	return nil
}

func (e *Department) SetLocation(location any) error {
	if location == nil {
		e.Location.Valid = false
		return nil
	}

	e.Location.String = utils.ToString(location)
	e.Location.Valid = true
	return nil
}

func (e *Department) SetBudget(budget any) error {
	x, err := utils.ToFloat64(budget)
	if err != nil {
		return err
	}
	e.Budget = x
	return nil
}

func (e *Department) SetCreatedAt(createdAt any) error {
	e.CreatedAt = utils.ToString(createdAt)
	return nil
}

func (e *Department) SetUpdatedAt(updatedAt any) error {
	e.UpdatedAt = utils.ToString(updatedAt)
	return nil
}


func (e *Department) ToDepartmentDto() dto.DepartmentDto {
	var dDto dto.DepartmentDto

	dDto.Code = e.Code
	dDto.Name = e.Name
	if e.Description.Valid != false {
		dDto.Description = e.Description.String
	}
	if e.ManagerId.Valid != false {
		dDto.ManagerId = e.ManagerId.Int64
	}
	if e.Location.Valid != false {
		dDto.Location = e.Location.String
	}
	dDto.Budget = e.Budget
	dDto.CreatedAt = e.CreatedAt
	dDto.UpdatedAt = e.UpdatedAt

	return dDto
}
