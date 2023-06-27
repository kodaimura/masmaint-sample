package service

import (
	"errors"

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
	Select(d *entity.Department) (entity.Department, error)
}

type DepartmentService struct {
	dDao *dao.DepartmentDao
}


func NewDepartmentService() *DepartmentService {
	dDao := dao.NewDepartmentDao()
	return &DepartmentService{dDao}
}


func (serv *DepartmentService) Create(dDto *dto.DepartmentDto) error {
	var d *entity.Department = entity.NewDepartment()
	d.SetName(dDto.Name)
	d.SetDescription(dDto.Description)
	d.SetManagerId(dDto.ManagerId)
	d.SetLocation(dDto.Location)
	d.SetBudget(dDto.Budget)

	err := serv.dDao.Insert(d)

	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}


func (serv *DepartmentService) Update(dDto *dto.DepartmentDto) error {
	var d *entity.Department = entity.NewDepartment()

	if d.SetId(dDto.Id) != nil || d.SetName(dDto.Name) != nil || 
	d.SetDescription(dDto.Description) != nil || 
	d.SetManagerId(dDto.ManagerId) != nil || 
	d.SetLocation(dDto.Location) != nil || 
	d.SetBudget(dDto.Budget) != nil {
		return errors.New("不正な値があります。")
	}

	err := serv.dDao.Update(d)

	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}


func (serv *DepartmentService) Delete(dDto *dto.DepartmentDto) error {
	var d *entity.Department = entity.NewDepartment()
	d.SetId(dDto.Id)

	err := serv.dDao.Delete(d)

	if err != nil {
		logger.LogError(err.Error())
	}

	return err
}


func (serv *DepartmentService) GetAll() ([]dto.DepartmentDto, error) {
	ds, err := serv.dDao.SelectAll()

	if err != nil {
		logger.LogError(err.Error())
	}

	var ret []dto.DepartmentDto
	for _, d := range ds {
		ret = append(ret, d.ToDepartmentDto())
	}

	return ret, err
}


func (serv *DepartmentService) GetOne(dDto *dto.DepartmentDto) (dto.DepartmentDto, error) {
	var d *entity.Department = entity.NewDepartment()
	d.SetId(dDto.Id)
	ret, err := serv.dDao.Select(d)

	if err != nil {
		logger.LogError(err.Error())
	}

	return ret.ToDepartmentDto(), err
}

