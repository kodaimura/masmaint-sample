package service

import (
	"errors"

	"masmaint/core/logger"
	"masmaint/model/entity"
	"masmaint/model/dao"
	"masmaint/dto"
)


type EmployeeDao interface {
	Insert(e *entity.Employee) (entity.Employee, error)
	Update(e *entity.Employee) (entity.Employee, error)
	Delete(e *entity.Employee) error
	SelectAll() ([]entity.Employee, error)
	Select(e *entity.Employee) (entity.Employee, error)
}

type EmployeeService struct {
	eDao *dao.EmployeeDao
}


func NewEmployeeService() *EmployeeService {
	eDao := dao.NewEmployeeDao()
	return &EmployeeService{eDao}
}


func (serv *EmployeeService) Create(eDto *dto.EmployeeDto) (dto.EmployeeDto, error) {
	var e *entity.Employee = entity.NewEmployee()

	if e.SetFirstName(eDto.FirstName) != nil || 
	e.SetLastName(eDto.LastName) != nil || 
	e.SetEmail(eDto.Email) != nil || 
	e.SetPhoneNumber(eDto.PhoneNumber) != nil || 
	e.SetAddress(eDto.Address) != nil || 
	e.SetHireDate(eDto.HireDate) != nil || 
	e.SetJobTitle(eDto.JobTitle) != nil ||
	e.SetDepartmentId(eDto.DepartmentId) != nil || 
	e.SetSalary(eDto.Salary) != nil {
		return dto.EmployeeDto{}, errors.New("不正な値があります。")
	}

	ret, err := serv.eDao.Insert(e)

	if err != nil {
		logger.LogError(err.Error())
	}

	return ret.ToEmployeeDto(), err
}


func (serv *EmployeeService) Update(eDto *dto.EmployeeDto) (dto.EmployeeDto, error) {
	var e *entity.Employee = entity.NewEmployee()

	if e.SetId(eDto.Id) != nil || 
	e.SetFirstName(eDto.FirstName) != nil || 
	e.SetLastName(eDto.LastName) != nil || 
	e.SetEmail(eDto.Email) != nil || 
	e.SetPhoneNumber(eDto.PhoneNumber) != nil || 
	e.SetAddress(eDto.Address) != nil || 
	e.SetHireDate(eDto.HireDate) != nil || 
	e.SetJobTitle(eDto.JobTitle) != nil ||
	e.SetDepartmentId(eDto.DepartmentId) != nil || 
	e.SetSalary(eDto.Salary) != nil {
		return dto.EmployeeDto{}, errors.New("不正な値があります。")
	}

	ret, err := serv.eDao.Update(e)

	if err != nil {
		logger.LogError(err.Error())
	}

	return ret.ToEmployeeDto(), err
}


func (serv *EmployeeService) Delete(eDto *dto.EmployeeDto) error {
	var e *entity.Employee = entity.NewEmployee()

	if e.SetId(eDto.Id) != nil {
		return errors.New("不正な値があります。")
	}

	err := serv.eDao.Delete(e)

	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}


func (serv *EmployeeService) GetAll() ([]dto.EmployeeDto, error) {
	es, err := serv.eDao.SelectAll()

	if err != nil {
		logger.LogError(err.Error())
	}

	var ret []dto.EmployeeDto
	for _, e := range es {
		ret = append(ret, e.ToEmployeeDto())
	}

	return ret, err
}


func (serv *EmployeeService) GetOne(eDto *dto.EmployeeDto) (dto.EmployeeDto, error) {
	var e *entity.Employee = entity.NewEmployee()
	e.SetId(eDto.Id)
	ret, err := serv.eDao.Select(e)

	if err != nil {
		logger.LogError(err.Error())
	}

	return ret.ToEmployeeDto(), err
}

