package department

import (
	"goat/internal/core/logger"
	"goat/internal/core/utils"
)

type Service interface {
	Get() ([]Department, error)
	Create(input *PostBody) (Department, error)
	Update(input *PutBody) (Department, error)
	Delete(input *DeleteBody) error
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
	rows, err := srv.repository.Get()
	if err != nil {
		logger.Error(err.Error())
		return []Department{}, err
	}
	return rows, nil
}


func (srv *service) Create(input PostBody) (Department, error) {
	var model Department
	utils.MapFields(&model, input)

	id, err := srv.repository.Insert(model, nil)
	if err != nil {
		logger.Error(err.Error())
		return Department{}, err
	}

	return srv.repository.GetOne(&Department{ Id:id })
}


func (srv *service) Update(input PutBody) (Department, error) {
	var model Department
	utils.MapFields(&model, input)

	if err := srv.repository.Update(model, nil); err != nil {
		logger.Error(err.Error())
		return Department{}, err
	}

	return srv.repository.GetOne(&Department{ Id:input.Id })
}


func (srv *service) Delete(input DeleteBody) error {
	var model Department
	utils.MapFields(&model, input)

	if err := srv.repository.Delete(model, nil); err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}