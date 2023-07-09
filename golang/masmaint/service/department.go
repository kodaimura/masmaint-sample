package service

import (
	"errors"

	"masmaint/core/logger"
	"masmaint/model/entity"
	"masmaint/model/dao"
	"masmaint/dto"
)


type DepartmentDao interface {
	SelectAll() ([]entity.Department, error)
	Select(d *entity.Department) (entity.Department, error)
	Insert(d *entity.Department) (entity.Department, error)
	Update(d *entity.Department) (entity.Department, error)
	Delete(d *entity.Department) error
}

type departmentService struct {
	dDao DepartmentDao
}

func NewDepartmentService() *departmentService {
	dDao := dao.NewDepartmentDao()
	return &departmentService{dDao}
}


func (serv *departmentService) GetAll() ([]dto.DepartmentDto, error) {
	rows, err := serv.dDao.SelectAll()
	if err != nil {
		logger.LogError(err.Error())
		return []dto.DepartmentDto{}, errors.New("取得に失敗しました。")
	}

	var ret []dto.DepartmentDto
	for _, row := range rows {
		ret = append(ret, row.ToDepartmentDto())
	}

	return ret, nil
}


func (serv *departmentService) GetOne(dDto *dto.DepartmentDto) (dto.DepartmentDto, error) {
	var d *entity.Department = entity.NewDepartment()

	if d.SetCode(dDto.Code) != nil {
		return dto.DepartmentDto{}, errors.New("不正な値があります。")
	}

	row, err := serv.dDao.Select(d)
	if err != nil {
		logger.LogError(err.Error())
		return dto.DepartmentDto{}, errors.New("取得に失敗しました。")
	}

	return row.ToDepartmentDto(), nil
}


func (serv *departmentService) Create(dDto *dto.DepartmentDto) (dto.DepartmentDto, error) {
	var d *entity.Department = entity.NewDepartment()

	if d.SetCode(dDto.Code) != nil ||
	d.SetName(dDto.Name) != nil ||
	d.SetDescription(dDto.Description) != nil ||
	d.SetManagerId(dDto.ManagerId) != nil ||
	d.SetLocation(dDto.Location) != nil ||
	d.SetBudget(dDto.Budget) != nil {
		return dto.DepartmentDto{}, errors.New("不正な値があります。")
	}

	row, err := serv.dDao.Insert(d)
	if err != nil {
		logger.LogError(err.Error())
		return dto.DepartmentDto{}, errors.New("登録に失敗しました。")
	}

	return row.ToDepartmentDto(), nil
}


func (serv *departmentService) Update(dDto *dto.DepartmentDto) (dto.DepartmentDto, error) {
	var d *entity.Department = entity.NewDepartment()

	if d.SetCode(dDto.Code) != nil ||
	d.SetName(dDto.Name) != nil ||
	d.SetDescription(dDto.Description) != nil ||
	d.SetManagerId(dDto.ManagerId) != nil ||
	d.SetLocation(dDto.Location) != nil ||
	d.SetBudget(dDto.Budget) != nil {
		return dto.DepartmentDto{}, errors.New("不正な値があります。")
	}

	row, err := serv.dDao.Update(d)
	if err != nil {
		logger.LogError(err.Error())
		return dto.DepartmentDto{}, errors.New("更新に失敗しました。")
	}

	return row.ToDepartmentDto(), nil
}


func (serv *departmentService) Delete(dDto *dto.DepartmentDto) error {
	var d *entity.Department = entity.NewDepartment()

	if d.SetCode(dDto.Code) != nil {
		return errors.New("不正な値があります。")
	}

	err := serv.dDao.Delete(d)
	if err != nil {
		logger.LogError(err.Error())
		return errors.New("削除に失敗しました。")
	}

	return nil
}
