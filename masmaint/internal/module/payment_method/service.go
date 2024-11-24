package payment_method

import (
	"masmaint/internal/module"
	"masmaint/internal/core/logger"
	"masmaint/internal/core/utils"
	"masmaint/internal/core/errs"
)

type Service interface {
	Get() ([]PaymentMethod, error)
	Create(input PostBody) (PaymentMethod, error)
	Update(input PutBody) (PaymentMethod, error)
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


func (srv *service) Get() ([]PaymentMethod, error) {
	rows, err := srv.repository.Get(&PaymentMethod{})
	if err != nil {
		logger.Error(err.Error())
		return []PaymentMethod{}, errs.NewUnexpectedError(err.Error())
	}
	return rows, nil
}


func (srv *service) Create(input PostBody) (PaymentMethod, error) {
	var model PaymentMethod
	utils.MapFields(&model, input)

	err := srv.repository.Insert(&model, nil)
	if err != nil {
		if column, ok := module.GetConflictColumn(err); ok {
			return PaymentMethod{}, errs.NewConflictError(column)
		}
		logger.Error(err.Error())
		return PaymentMethod{}, errs.NewUnexpectedError(err.Error())
	}

	row, err := srv.repository.GetOne(&PaymentMethod{ Code: input.Code })
	if err != nil {
		logger.Error(err.Error())
		return PaymentMethod{}, errs.NewUnexpectedError(err.Error())
	}
	return row, nil
}


func (srv *service) Update(input PutBody) (PaymentMethod, error) {
	var model PaymentMethod
	utils.MapFields(&model, input)

	err := srv.repository.Update(&model, nil)
	if err != nil {
		if column, ok := module.GetConflictColumn(err); ok {
			return PaymentMethod{}, errs.NewConflictError(column)
		}
		logger.Error(err.Error())
		return PaymentMethod{}, errs.NewUnexpectedError(err.Error())
	}

	row, err := srv.repository.GetOne(&PaymentMethod{ Code: input.Code })
	if err != nil {
		logger.Error(err.Error())
		return PaymentMethod{}, errs.NewUnexpectedError(err.Error())
	}
	return row, nil
}


func (srv *service) Delete(input DeleteBody) error {
	var model PaymentMethod
	utils.MapFields(&model, input)

	err := srv.repository.Delete(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return errs.NewUnexpectedError(err.Error())
	}
	return nil
}