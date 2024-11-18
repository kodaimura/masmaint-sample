package payment_method

import (
	"masmaint/internal/core/logger"
	"masmaint/internal/core/utils"
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
		return []PaymentMethod{}, err
	}
	return rows, nil
}


func (srv *service) Create(input PostBody) (PaymentMethod, error) {
	var model PaymentMethod
	utils.MapFields(&model, input)

	err := srv.repository.Insert(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return PaymentMethod{}, err
	}

	return srv.repository.GetOne(&PaymentMethod{ Code: input.Code })
}


func (srv *service) Update(input PutBody) (PaymentMethod, error) {
	var model PaymentMethod
	utils.MapFields(&model, input)

	err := srv.repository.Update(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return PaymentMethod{}, err
	}

	return srv.repository.GetOne(&PaymentMethod{ Code: input.Code })
}


func (srv *service) Delete(input DeleteBody) error {
	var model PaymentMethod
	utils.MapFields(&model, input)

	err := srv.repository.Delete(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}