package product_category

import (
	"masmaint/internal/module"
	"masmaint/internal/core/logger"
	"masmaint/internal/core/utils"
	"masmaint/internal/core/errs"
)

type Service interface {
	Get() ([]ProductCategory, error)
	Create(input PostBody) (ProductCategory, error)
	Update(input PutBody) (ProductCategory, error)
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


func (srv *service) Get() ([]ProductCategory, error) {
	rows, err := srv.repository.Get(&ProductCategory{})
	if err != nil {
		logger.Error(err.Error())
		return []ProductCategory{}, errs.NewUnexpectedError(err.Error())
	}
	return rows, nil
}


func (srv *service) Create(input PostBody) (ProductCategory, error) {
	var model ProductCategory
	utils.MapFields(&model, input)

	id, err := srv.repository.Insert(&model, nil)
	if err != nil {
		if column, ok := module.GetConflictColumn(err); ok {
			return ProductCategory{}, errs.NewConflictError(column)
		}
		logger.Error(err.Error())
		return ProductCategory{}, errs.NewUnexpectedError(err.Error())
	}

	row, err := srv.repository.GetOne(&ProductCategory{ Id: id })
	if err != nil {
		logger.Error(err.Error())
		return ProductCategory{}, errs.NewUnexpectedError(err.Error())
	}
	return row, nil
}


func (srv *service) Update(input PutBody) (ProductCategory, error) {
	var model ProductCategory
	utils.MapFields(&model, input)

	err := srv.repository.Update(&model, nil)
	if err != nil {
		if column, ok := module.GetConflictColumn(err); ok {
			return ProductCategory{}, errs.NewConflictError(column)
		}
		logger.Error(err.Error())
		return ProductCategory{}, errs.NewUnexpectedError(err.Error())
	}

	row, err := srv.repository.GetOne(&ProductCategory{ Id: input.Id })
	if err != nil {
		logger.Error(err.Error())
		return ProductCategory{}, errs.NewUnexpectedError(err.Error())
	}
	return row, nil
}


func (srv *service) Delete(input DeleteBody) error {
	var model ProductCategory
	utils.MapFields(&model, input)

	err := srv.repository.Delete(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return errs.NewUnexpectedError(err.Error())
	}
	return nil
}