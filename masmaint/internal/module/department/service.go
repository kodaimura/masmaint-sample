package department

import (
	"masmaint/internal/core/logger"
	"masmaint/internal/core/errs"
	"masmaint/internal/core/utils"
)

type Service interface {
	Get() ([]Department, error)
	Create(input PostBody) (Department, error)
	Update(input PutBody) (Department, error)
	Delete(input DeleteBody) error
}

type service struct {
	repository Repository
}

func NewService() Service {
	return &service{
		repository: NewRepository(),
	}
}


func (srv *service) Get() ([]Department, error) {
	rows, err := srv.repository.Get(&Department{})
	if err != nil {
		logger.Error(err.Error())
		return []Department{}, errs.NewError(err)
	}
	return rows, nil
}


func (srv *service) Create(input PostBody) (Department, error) {
	var model Department
	utils.MapFields(&model, input)

	err := srv.repository.Insert(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return Department{}, errs.NewError(err)
	}

	return srv.repository.GetOne(&Department{ Code: input.Code })
}


func (srv *service) Update(input PutBody) (Department, error) {
	var model Department
	utils.MapFields(&model, input)

	err := srv.repository.Delete(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return Department{}, errs.NewError(err)
	}

	return srv.repository.GetOne(&Department{ Code: input.Code })
}


func (srv *service) Delete(input DeleteBody) error {
	var model Department
	utils.MapFields(&model, input)

	err := srv.repository.Delete(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return errs.NewError(err)
	}
	return nil
}