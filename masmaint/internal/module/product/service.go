package product

import (
	"masmaint/internal/module"
	"masmaint/internal/core/logger"
	"masmaint/internal/core/utils"
	"masmaint/internal/core/errs"
)

type Service interface {
	Get() ([]Product, error)
	Create(input PostBody) (Product, error)
	Update(input PutBody) (Product, error)
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


func (srv *service) Get() ([]Product, error) {
	rows, err := srv.repository.Get(&Product{})
	if err != nil {
		logger.Error(err.Error())
		return []Product{},  errs.NewUnexpectedError(err.Error())
	}
	return rows, nil
}


func (srv *service) Create(input PostBody) (Product, error) {
	var model Product
	utils.MapFields(&model, input)

	id, err := srv.repository.Insert(&model, nil)
	if err != nil {
		if column, ok := module.GetConflictColumn(err); ok {
			return Product{}, errs.NewConflictError(column)
		}
		logger.Error(err.Error())
		return Product{}, errs.NewUnexpectedError(err.Error())
	}

	row, err := srv.repository.GetOne(&Product{ Id: id })
	if err != nil {
		logger.Error(err.Error())
		return Product{}, errs.NewUnexpectedError(err.Error())
	}
	return row, nil
}


func (srv *service) Update(input PutBody) (Product, error) {
	var model Product
	utils.MapFields(&model, input)

	err := srv.repository.Update(&model, nil)
	if err != nil {
		if column, ok := module.GetConflictColumn(err); ok {
			return Product{}, errs.NewConflictError(column)
		}
		logger.Error(err.Error())
		return Product{}, errs.NewUnexpectedError(err.Error())
	}

	row, err := srv.repository.GetOne(&Product{ Id: input.Id })
	if err != nil {
		logger.Error(err.Error())
		return Product{}, errs.NewUnexpectedError(err.Error())
	}
	return row, nil
}


func (srv *service) Delete(input DeleteBody) error {
	var model Product
	utils.MapFields(&model, input)

	err := srv.repository.Delete(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return errs.NewUnexpectedError(err.Error())
	}
	return nil
}