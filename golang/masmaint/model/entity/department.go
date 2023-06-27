package entity

import (
	"strconv"
	"errors"
	"database/sql"

	"masmaint/dto"
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
	switch id.(type) {
	case int:
		d.Id = id.(int64)

	case string:
		x, err := strconv.Atoi(id.(string))
		if err != nil {
			return err
		}
		d.Id = int64(x)
	default:
		return errors.New("type error")
	}

	return nil
}

func (d *Department) SetName(name any) error {
	switch name.(type) {
	case string:
		d.Name = name.(string)
	default:
		return errors.New("type error")
	}

	return nil
}

func (d *Department) SetDescription(description any) error {
	if description == nil {
		d.Description.Valid = false
		return nil
	} 

	switch description.(type) {
	case string:
		d.Description.String = description.(string)
		d.Description.Valid = true
	default:
		return errors.New("type error")
	}

	return nil
}

func (d *Department) SetManagerId(managerId any) error {
	if managerId == nil {
		d.ManagerId.Valid = false
		return nil	
	} 

	switch managerId.(type) {
	case int:
		d.ManagerId.Int64 = managerId.(int64)
		d.ManagerId.Valid = true
	case string:
		if managerId == "" {
			d.ManagerId.Valid = false
		} else {
			x, err := strconv.Atoi(managerId.(string))
			if err != nil {
				return err
			}
			d.ManagerId.Int64 = int64(x)
			d.ManagerId.Valid = true
		}
	default:
		return errors.New("type error")
	}

	return nil
}

func (d *Department) SetLocation(location any) error {
	if location == nil {
		d.Location.Valid = false
		return nil
	} 

	switch location.(type) {
	case string:
		d.Location.String = location.(string)
		d.Location.Valid = true
	default:
		return errors.New("type error")
	}

	return nil
}

func (d *Department) SetBudget(budget any) error {
	if budget == nil {
		d.Budget.Valid = false
		return nil	
	}

	switch budget.(type) {
	case int32,int64,float32,float64:
		d.Budget.Float64 = budget.(float64)
		d.Budget.Valid = true
	case string:
		if budget == "" {
			d.Budget.Valid = false
		} else {
			x, err := strconv.ParseFloat(budget.(string), 64)
			if err != nil {
				return err
			}
			d.Budget.Float64 = float64(x)
			d.Budget.Valid = true
		}
	default:
		return errors.New("type error")
	}

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