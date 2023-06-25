package service

import (
	"masmaint/core/logger"
	"masmaint/model/entity"
	"masmaint/model/dao"
	"masmaint/dto"
)


type DepartmentDao interface {
	Insert(e *entity.Department) error
	Select(id int) (entity.Department, error)
	Update(e *entity.Department) error
	Delete(id int) error
	SelectAll() ([]entity.Department, error)
}


type DepartmentService struct {
	dDao dao.DepartmentDao
}


func NewDepartmentService() *DepartmentService {
	dDao := dao.NewDepartmentDao()
	return &DepartmentService{dDao}
}


func (serv *DepartmentService) CreateDepartment(d *dto.DepartmentDto) error {
	var department entity.Department
	department.Name = d.Name
	department.Description = d.Description
	department.ManagerId = d.ManagerId
	department.Location = d.Location
	department.Budget = d.Budget

	err = serv.dDao.Insert(&department)

	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}


func (serv *DepartmentService) UpdateDepartment(d *dto.DepartmentDto) error {
	var department entity.Department
	department.Id = d.Id
	department.Name = d.Name
	department.Description = d.Description
	department.ManagerId = d.ManagerId
	department.Location = d.Location
	department.Budget = d.Budget

	err := serv.dDao.Update(&department)

	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}


func (serv *DepartmentService) DeleteDepartment(id int) err {
	err := serv.dDao.Delete(id)

	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}


func (serv *DepartmentService) GetDepartments() ([]entity.Department, error) {
	departments, err := serv.dDao.SelectAll()

	if err != nil {
		logger.LogError(err.Error())
	}

	return departments, err
}

