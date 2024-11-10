package employee

import (
	"masmaint/internal/core/logger"
	"masmaint/internal/core/utils"
)

type Service interface {
	Get() ([]Employee, error)
	Create(input PostBody) (Employee, error)
	Update(input PutBody) (Employee, error)
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


func (srv *service) Get() ([]Employee, error) {
	rows, err := srv.repository.Get(&Employee{})
	if err != nil {
		logger.Error(err.Error())
		return []Employee{}, err
	}
	return rows, nil
}


func (srv *service) Create(input PostBody) (Employee, error) {
	var model Employee
	utils.MapFields(&model, input)

	id, err := srv.repository.Insert(&model, nil)
	if err != nil {
		logger.Error(err.Error())
		return Employee{}, err
	}

	return srv.repository.GetOne(&Employee{ Id: id })
}


func (srv *service) Update(input PutBody) (Employee, error) {
	var model Employee
	utils.MapFields(&model, input)

	if err := srv.repository.Update(&model, nil); err != nil {
		logger.Error(err.Error())
		return Employee{}, err
	}

	return srv.repository.GetOne(&Employee{ Id: input.Id })
}


func (srv *service) Delete(input DeleteBody) error {
	var model Employee
	utils.MapFields(&model, input)

	if err := srv.repository.Delete(&model, nil); err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}