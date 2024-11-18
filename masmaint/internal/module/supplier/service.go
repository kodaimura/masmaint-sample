package supplier

import (
	"masmaint/internal/core/logger"
	"masmaint/internal/core/utils"
)

type Service interface {
	Get() ([]Supplier, error)
	Create(input PostBody) (Supplier, error)
	Update(input PutBody) (Supplier, error)
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


func (srv *service) Get() ([]Supplier, error) {
	rows, err := srv.repository.Get(&Supplier{})
	if err != nil {
		logger.Error(err.Error())
		return []Supplier{}, err
	}
	return rows, nil
}


func (srv *service) Create(input PostBody) (Supplier, error) {
	var model Supplier
	utils.MapFields(&model, input)

	id, err := srv.repository.Insert(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return Supplier{}, err
	}

	return srv.repository.GetOne(&Supplier{ Id: id })
}


func (srv *service) Update(input PutBody) (Supplier, error) {
	var model Supplier
	utils.MapFields(&model, input)

	err := srv.repository.Update(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return Supplier{}, err
	}

	return srv.repository.GetOne(&Supplier{ Id: input.Id })
}


func (srv *service) Delete(input DeleteBody) error {
	var model Supplier
	utils.MapFields(&model, input)

	err := srv.repository.Delete(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}