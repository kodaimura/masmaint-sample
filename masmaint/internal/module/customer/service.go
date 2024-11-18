package customer

import (
	"masmaint/internal/core/logger"
	"masmaint/internal/core/utils"
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
		return []Customer{}, err
	}
	return rows, nil
}


func (srv *service) Create(input PostBody) (Customer, error) {
	var model Customer
	utils.MapFields(&model, input)

	id, err := srv.repository.Insert(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return Customer{}, err
	}

	return srv.repository.GetOne(&Customer{ Id: id })
}


func (srv *service) Update(input PutBody) (Customer, error) {
	var model Customer
	utils.MapFields(&model, input)

	err := srv.repository.Update(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return Customer{}, err
	}

	return srv.repository.GetOne(&Customer{ Id: input.Id })
}


func (srv *service) Delete(input DeleteBody) error {
	var model Customer
	utils.MapFields(&model, input)

	err := srv.repository.Delete(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}