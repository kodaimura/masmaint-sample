package service

import (
	"masmaint/core/logger"
	"masmaint/model/entity"
	"masmaint/model/dao"
	"masmaint/dto"
)


type EmployeeDao interface {
	Insert(e *entity.Employee) error
	Update(e *entity.Employee) error
	Delete(e *entity.Employee) error
	SelectAll() ([]entity.Employee, error)
	Select(id int) (entity.Employee, error)
}

type EmployeeService struct {
	eDao *dao.EmployeeDao
}


func NewEmployeeService() *EmployeeService {
	eDao := dao.NewEmployeeDao()
	return &EmployeeService{eDao}
}


func (serv *EmployeeService) Create(eDto *dto.EmployeeDto) error {
	var e entity.Employee
	e.FirstName = eDto.FirstName
	e.LastName = eDto.LastName
	e.Email = eDto.Email
	e.PhoneNumber = eDto.PhoneNumber
	e.Address = eDto.Address
	e.HireDate = eDto.HireDate
	e.JobTitle = eDto.JobTitle
	e.DepartmentId = eDto.DepartmentId
	e.Salary = eDto.Salary

	err := serv.eDao.Insert(&e)

	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}


func (serv *EmployeeService) Update(eDto *dto.EmployeeDto) error {
	var e entity.Employee
	e.Id = eDto.Id
	e.FirstName = eDto.FirstName
	e.LastName = eDto.LastName
	e.Email = eDto.Email
	e.PhoneNumber = eDto.PhoneNumber
	e.Address = eDto.Address
	e.HireDate = eDto.HireDate
	e.JobTitle = eDto.JobTitle
	e.DepartmentId = eDto.DepartmentId
	e.Salary = eDto.Salary

	err := serv.eDao.Update(&e)

	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}


func (serv *EmployeeService) Delete(eDto *dto.EmployeeDto) error {
	var e entity.Employee
	e.Id = eDto.Id

	err := serv.eDao.Delete(&e)

	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}


func (serv *EmployeeService) GetAll() ([]entity.Employee, error) {
	es, err := serv.eDao.SelectAll()

	if err != nil {
		logger.LogError(err.Error())
	}

	return es, err
}


func (serv *EmployeeService) GetOne(id int) (entity.Employee, error) {
	e, err := serv.eDao.Select(id)

	if err != nil {
		logger.LogError(err.Error())
	}

	return e, err
}

