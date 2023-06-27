package entity

import (
	"database/sql"

	"masmaint/dto"
	"masmaint/core/utils"
)

type Department struct {
	Id int64 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Description sql.NullString `db:"description" json:"description"`
	ManagerId sql.NullInt64 `db:"manager_id" json:"manager_id"`
	Location sql.NullString `db:"location" json:"location"`
	Budget sql.NullFloat64 `db:"budget" json:"budget"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

func NewDepartment() *Department {
	return &Department{}
}

func (d *Department) SetId(id any) error {
	x, err := utils.ToInt64(id)
	if err != nil {
		return err
	}
	d.Id = x
	return nil
}

func (d *Department) SetName(name any) error {
	d.Name = utils.ToString(name)
	return nil
}

func (d *Department) SetDescription(description any) error {
	if description == nil {
		d.Description.Valid = false
		return nil
	} 

	d.Description.String = utils.ToString(description)
	d.Description.Valid = true
	return nil
}

func (d *Department) SetManagerId(managerId any) error {
	if managerId == nil {
		d.ManagerId.Valid = false
		return nil
	} 
	if managerId == "" {
		d.ManagerId.Valid = false
		return nil
	} 

	x, err := utils.ToInt64(managerId)
	if err != nil {
		return err
	}
	d.ManagerId.Int64 = x
	d.ManagerId.Valid = true
	return nil
}

func (d *Department) SetLocation(location any) error {
	if location == nil {
		d.Location.Valid = false
		return nil
	} 

	d.Location.String = utils.ToString(location)
	d.Location.Valid = true
	return nil
}

func (d *Department) SetBudget(budget any) error {
	if budget == nil {
		d.Budget.Valid = false
		return nil	
	}
	if budget == "" {
		d.Budget.Valid = false
		return nil
	}

	x, err := utils.ToFloat64(budget)
	if err != nil {
		return err
	}
	d.Budget.Float64 = x
	d.Budget.Valid = true
	return nil
}


func (d *Department) ToDepartmentDto() dto.DepartmentDto {
	var dDto dto.DepartmentDto

	dDto.Id = d.Id
	dDto.Name = d.Name
	if d.Description.Valid != false {
		dDto.Description = d.Description.String
	}
	if d.ManagerId.Valid != false {
		dDto.ManagerId = d.ManagerId.Int64
	}
	if d.Location.Valid != false {
		dDto.Location = d.Location.String
	}
	if d.Budget.Valid != false {
		dDto.Budget = d.Budget.Float64
	}
	dDto.CreatedAt = d.CreatedAt
	dDto.UpdatedAt = d.UpdatedAt

	return dDto
}