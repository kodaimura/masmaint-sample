package entity

import (
	"database/sql"

	"masmaint/dto"
	"masmaint/core/utils"
)

type Department struct {
	Id int64 `db:"id"`
	Name string `db:"name"`
	Description sql.NullString `db:"description"`
	ManagerId sql.NullInt64 `db:"manager_id"`
	Location sql.NullString `db:"location"`
	Budget sql.NullFloat64 `db:"budget"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func NewDepartment() *Department {
	return &Department{}
}

func (e *Department) SetId(id any) error {
	x, err := utils.ToInt64(id)
	if err != nil {
		return err
	}
	e.Id = x
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
	if managerId == nil {
		e.ManagerId.Valid = false
		return nil
	} 
	if managerId == "" {
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
	if budget == nil {
		e.Budget.Valid = false
		return nil	
	}
	if budget == "" {
		e.Budget.Valid = false
		return nil
	}

	x, err := utils.ToFloat64(budget)
	if err != nil {
		return err
	}
	e.Budget.Float64 = x
	e.Budget.Valid = true
	return nil
}


func (e *Department) ToDepartmentDto() dto.DepartmentDto {
	var dDto dto.DepartmentDto

	dDto.Id = e.Id
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
	if e.Budget.Valid != false {
		dDto.Budget = e.Budget.Float64
	}
	dDto.CreatedAt = e.CreatedAt
	dDto.UpdatedAt = e.UpdatedAt

	return dDto
}