package service

import (
	cerror "masmaint/core/error"
	"masmaint/core/logger"
	"masmaint/model/entity"
	"masmaint/model/dao"
	"masmaint/dto"
)


type EmployeeDao interface {
	SelectAll() ([]entity.Employee, error)
	Select(e *entity.Employee) (entity.Employee, error)
	Insert(e *entity.Employee) (entity.Employee, error)
	Update(e *entity.Employee) (entity.Employee, error)
	Delete(e *entity.Employee) error
}

type employeeService struct {
	eDao EmployeeDao
}

func NewEmployeeService() *employeeService {
	eDao := dao.NewEmployeeDao()
	return &employeeService{eDao}
}


func (serv *employeeService) GetAll() ([]dto.EmployeeDto, error) {
	rows, err := serv.eDao.SelectAll()
	if err != nil {
		logger.LogError(err.Error())
		return []dto.EmployeeDto{}, cerror.NewDaoError("取得に失敗しました。")
	}

	var ret []dto.EmployeeDto
	for _, row := range rows {
		ret = append(ret, row.ToEmployeeDto())
	}

	return ret, nil
}


func (serv *employeeService) GetOne(eDto *dto.EmployeeDto) (dto.EmployeeDto, error) {
	var e *entity.Employee = entity.NewEmployee()

	if e.SetId(eDto.Id) != nil {
		return dto.EmployeeDto{}, cerror.NewInvalidArgumentError("不正な値があります。")
	}

	row, err := serv.eDao.Select(e)
	if err != nil {
		logger.LogError(err.Error())
		return dto.EmployeeDto{}, cerror.NewDaoError("取得に失敗しました。")
	}

	return row.ToEmployeeDto(), nil
}


func (serv *employeeService) Create(eDto *dto.EmployeeDto) (dto.EmployeeDto, error) {
	var e *entity.Employee = entity.NewEmployee()

	if e.SetFirstName(eDto.FirstName) != nil ||
	e.SetLastName(eDto.LastName) != nil ||
	e.SetEmail(eDto.Email) != nil ||
	e.SetPhoneNumber(eDto.PhoneNumber) != nil ||
	e.SetAddress(eDto.Address) != nil ||
	e.SetHireDate(eDto.HireDate) != nil ||
	e.SetJobTitle(eDto.JobTitle) != nil ||
	e.SetDepartmentCode(eDto.DepartmentCode) != nil ||
	e.SetSalary(eDto.Salary) != nil {
		return dto.EmployeeDto{}, cerror.NewInvalidArgumentError("不正な値があります。")
	}

	row, err := serv.eDao.Insert(e)
	if err != nil {
		logger.LogError(err.Error())
		return dto.EmployeeDto{}, cerror.NewDaoError("登録に失敗しました。")
	}

	return row.ToEmployeeDto(), nil
}


func (serv *employeeService) Update(eDto *dto.EmployeeDto) (dto.EmployeeDto, error) {
	var e *entity.Employee = entity.NewEmployee()

	if e.SetId(eDto.Id) != nil ||
	e.SetFirstName(eDto.FirstName) != nil ||
	e.SetLastName(eDto.LastName) != nil ||
	e.SetEmail(eDto.Email) != nil ||
	e.SetPhoneNumber(eDto.PhoneNumber) != nil ||
	e.SetAddress(eDto.Address) != nil ||
	e.SetHireDate(eDto.HireDate) != nil ||
	e.SetJobTitle(eDto.JobTitle) != nil ||
	e.SetDepartmentCode(eDto.DepartmentCode) != nil ||
	e.SetSalary(eDto.Salary) != nil {
		return dto.EmployeeDto{}, cerror.NewInvalidArgumentError("不正な値があります。")
	}

	row, err := serv.eDao.Update(e)
	if err != nil {
		logger.LogError(err.Error())
		return dto.EmployeeDto{}, cerror.NewDaoError("更新に失敗しました。")
	}

	return row.ToEmployeeDto(), nil
}


func (serv *employeeService) Delete(eDto *dto.EmployeeDto) error {
	var e *entity.Employee = entity.NewEmployee()

	if e.SetId(eDto.Id) != nil {
		return cerror.NewInvalidArgumentError("不正な値があります。")
	}

	err := serv.eDao.Delete(e)
	if err != nil {
		logger.LogError(err.Error())
		return cerror.NewDaoError("削除に失敗しました。")
	}

	return nil
}
