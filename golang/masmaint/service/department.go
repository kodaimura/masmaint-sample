package service

import (
	"masmaint/core/logger"
	"masmaint/model/entity"
	"masmaint/model/dao"
	"masmaint/dto"
)


type DepartmentDao interface {
	Insert(d *entity.Department) error
	Update(d *entity.Department) error
	Delete(d *entity.Department) error
	SelectAll() ([]entity.Department, error)
	Select(id int) (entity.Department, error)
}

type DepartmentService struct {
	dDao *dao.DepartmentDao
}


func NewDepartmentService() *DepartmentService {
	dDao := dao.NewDepartmentDao()
	return &DepartmentService{dDao}
}


func (serv *DepartmentService) Create(dDto *dto.DepartmentDto) error {
	var d entity.Department
	d.Name = dDto.Name
	d.Description = dDto.Description
	d.ManagerId = dDto.ManagerId
	d.Location = dDto.Location
	d.Budget = dDto.Budget

	err := serv.dDao.Insert(&d)

	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}


func (serv *DepartmentService) Update(dDto *dto.DepartmentDto) error {
	var d entity.Department
	d.Id = dDto.Id
	d.Name = dDto.Name
	d.Description = dDto.Description
	d.ManagerId = dDto.ManagerId
	d.Location = dDto.Location
	d.Budget = dDto.Budget

	err := serv.dDao.Update(&d)

	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}


func (serv *DepartmentService) Delete(dDto *dto.DepartmentDto) error {
	var d entity.Department
	d.Id = dDto.Id

	err := serv.dDao.Delete(&d)

	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}


func (serv *DepartmentService) GetAll() ([]entity.Department, error) {
	ds, err := serv.dDao.SelectAll()

	if err != nil {
		logger.LogError(err.Error())
	}

	return ds, err
}


func (serv *DepartmentService) GetOne(id int) (entity.Department, error) {
	d, err := serv.dDao.Select(id)

	if err != nil {
		logger.LogError(err.Error())
	}

	return d, err
}

