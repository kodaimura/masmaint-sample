package service

import (
	"masmaint/core/logger"
	"masmaint/model/entity"
	"masmaint/model/dao"
	"masmaint/dto"
)


type EmployeeDao interface {
	Insert(e *entity.Employee) error
	Select(id int) (entity.Employee, error)
	Update(e *entity.Employee) error
	Delete(id int) error
	SelectAll() ([]entity.Employee, error)
}


type EmployeeService struct {
	eDao dao.EmployeeDao
}


func NewEmployeeService() *EmployeeService {
	dDao := dao.NewEmployeeDao()
	return &EmployeeService{eDao}
}


func (serv *EmployeeService) CreateEmployee(e *dto.EmployeeDto) error {
	var employee entity.Employee
	employee.FirstName = e.FirstName
	employee.LastName = e.LastName
	employee.Email = e.Email
	employee.PhoneNumber = e.PhoneNumber
	employee.Address = e.Address
	employee.HireDate = e.HireDate
	employee.JobTitle = e.JobTitle
	employee.DepartmentId = e.DepartmentId
	employee.Salary = e.Salary

	err = serv.eDao.Insert(&employee)

	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}


func (serv *EmployeeService) UpdateEmployee(e *dto.EmployeeDto) error {
	var employee entity.Employee
	employee.Id = e.Id
	employee.FirstName = e.FirstName
	employee.LastName = e.LastName
	employee.Email = e.Email
	employee.PhoneNumber = e.PhoneNumber
	employee.Address = e.Address
	employee.HireDate = e.HireDate
	employee.JobTitle = e.JobTitle
	employee.DepartmentId = e.DepartmentId
	employee.Salary = e.Salary

	err := serv.eDao.Update(&employee)

	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}


func (serv *userService) DeleteEmployee(id int) err {
	err := serv.eDao.Delete(id)

	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}


func (serv *userService) GetEmployees() ([]entity.Employee, error) {
	employee, err := serv.eDao.SelectAll()

	if err != nil {
		logger.LogError(err.Error())
	}

	return employee, err
}

