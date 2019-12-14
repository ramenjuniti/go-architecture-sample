package application

import (
	"server/domain/model"
	"server/domain/repository"

	"gopkg.in/gorp.v2"
)

type TodoApplication interface {
	GetAll(*gorp.DbMap) ([]model.Todo, error)
	GetByID(*gorp.DbMap, int64) (*model.Todo, error)
	Create(*gorp.DbMap, string, string) error
	Update(*gorp.DbMap, int64, string, string) error
	Delete(*gorp.DbMap, int64) error
}

type todoApplication struct {
	todoRepository repository.TodoRepository
}

func NewTodoApplication(tr repository.TodoRepository) TodoApplication {
	return &todoApplication{todoRepository: tr}
}

func (ta todoApplication) GetAll(dbm *gorp.DbMap) ([]model.Todo, error) {
	todos, err := ta.todoRepository.GetAll(dbm)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (ta todoApplication) GetByID(dbm *gorp.DbMap, id int64) (*model.Todo, error) {
	todo, err := ta.todoRepository.GetByID(dbm, id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (ta todoApplication) Create(dbm *gorp.DbMap, title, body string) error {
	todo := &model.Todo{Title: title, Body: body}

	trans, err := dbm.Begin()
	if err != nil {
		return err
	}

	if err = ta.todoRepository.Create(trans, todo); err != nil {
		return err
	}

	if err = trans.Commit(); err != nil {
		return err
	}

	return nil
}

func (ta todoApplication) Update(dbm *gorp.DbMap, id int64, title, body string) error {
	todo := &model.Todo{ID: id, Title: title, Body: body}

	trans, err := dbm.Begin()
	if err != nil {
		return err
	}

	if err = ta.todoRepository.Update(trans, todo); err != nil {
		return err
	}

	if err = trans.Commit(); err != nil {
		return err
	}

	return nil
}

func (ta todoApplication) Delete(dbm *gorp.DbMap, id int64) error {
	todo := &model.Todo{ID: id}

	trans, err := dbm.Begin()
	if err != nil {
		return err
	}

	if err = ta.todoRepository.Delete(trans, todo); err != nil {
		return err
	}

	if err = trans.Commit(); err != nil {
		return err
	}

	return nil
}
