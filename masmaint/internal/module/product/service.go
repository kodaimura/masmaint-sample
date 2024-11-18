package product

import (
	"masmaint/internal/core/logger"
	"masmaint/internal/core/utils"
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
		return []Product{}, err
	}
	return rows, nil
}


func (srv *service) Create(input PostBody) (Product, error) {
	var model Product
	utils.MapFields(&model, input)

	id, err := srv.repository.Insert(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return Product{}, err
	}

	return srv.repository.GetOne(&Product{ Id: id })
}


func (srv *service) Update(input PutBody) (Product, error) {
	var model Product
	utils.MapFields(&model, input)

	err := srv.repository.Update(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return Product{}, err
	}

	return srv.repository.GetOne(&Product{ Id: input.Id })
}


func (srv *service) Delete(input DeleteBody) error {
	var model Product
	utils.MapFields(&model, input)

	err := srv.repository.Delete(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}