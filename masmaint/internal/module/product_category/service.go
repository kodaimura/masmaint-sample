package product_category

import (
	"masmaint/internal/core/logger"
	"masmaint/internal/core/utils"
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
		return []ProductCategory{}, err
	}
	return rows, nil
}


func (srv *service) Create(input PostBody) (ProductCategory, error) {
	var model ProductCategory
	utils.MapFields(&model, input)

	id, err := srv.repository.Insert(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return ProductCategory{}, err
	}

	return srv.repository.GetOne(&ProductCategory{ Id: id })
}


func (srv *service) Update(input PutBody) (ProductCategory, error) {
	var model ProductCategory
	utils.MapFields(&model, input)

	err := srv.repository.Update(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return ProductCategory{}, err
	}

	return srv.repository.GetOne(&ProductCategory{ Id: input.Id })
}


func (srv *service) Delete(input DeleteBody) error {
	var model ProductCategory
	utils.MapFields(&model, input)

	err := srv.repository.Delete(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}