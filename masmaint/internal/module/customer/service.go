package customer

import (
	"masmaint/internal/module"
	"masmaint/internal/core/logger"
	"masmaint/internal/core/utils"
	"masmaint/internal/core/errs"
)

type Service interface {
	Get() ([]Customer, error)
	Create(input PostBody) (Customer, error)
	Update(input PutBody) (Customer, error)
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


func (srv *service) Get() ([]Customer, error) {
	rows, err := srv.repository.Get(&Customer{})
	if err != nil {
		logger.Error(err.Error())
		return []Customer{}, errs.NewUnexpectedError(err.Error())
	}
	return rows, nil
}


func (srv *service) Create(input PostBody) (Customer, error) {
	var model Customer
	utils.MapFields(&model, input)

	id, err := srv.repository.Insert(&model, nil)
	if err != nil {
		if column, ok := module.GetConflictColumn(err); ok {
			return Customer{}, errs.NewConflictError(column)
		}
		logger.Error(err.Error())
		return Customer{}, errs.NewUnexpectedError(err.Error())
	}

	row, err := srv.repository.GetOne(&Customer{ Id: id })
	if err != nil {
		logger.Error(err.Error())
		return Customer{}, errs.NewUnexpectedError(err.Error())
	}
	return row, nil
}


func (srv *service) Update(input PutBody) (Customer, error) {
	var model Customer
	utils.MapFields(&model, input)

	err := srv.repository.Update(&model, nil)
	if err != nil {
		if column, ok := module.GetConflictColumn(err); ok {
			return Customer{}, errs.NewConflictError(column)
		}
		logger.Error(err.Error())
		return Customer{}, errs.NewUnexpectedError(err.Error())
	}

	row, err := srv.repository.GetOne(&Customer{ Id: input.Id })
	if err != nil {
		logger.Error(err.Error())
		return Customer{}, errs.NewUnexpectedError(err.Error())
	}
	return row, nil
}


func (srv *service) Delete(input DeleteBody) error {
	var model Customer
	utils.MapFields(&model, input)

	err := srv.repository.Delete(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return errs.NewUnexpectedError(err.Error())
	}
	return nil
}