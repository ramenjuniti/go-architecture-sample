package persistence

import (
	"server/domain/model"
	"server/domain/repository"

	"gopkg.in/gorp.v2"
)

type todoPersistence struct{}

func NewTodoPersistence() repository.TodoRepository {
	return &todoPersistence{}
}

func (tp todoPersistence) GetAll(dbm *gorp.DbMap) ([]model.Todo, error) {
	t := make([]model.Todo, 0)
	if _, err := dbm.Select(&t, `SELECT id, title, body FROM todo`); err != nil {
		return nil, err
	}
	return t, nil
}

func (tp todoPersistence) GetByID(dbm *gorp.DbMap, id int64) (*model.Todo, error) {
	t := model.Todo{}
	if err := dbm.SelectOne(&t, `SELECT id, title, body FROM todo WHERE id = ?`, id); err != nil {
		return nil, err
	}
	return &t, nil
}

func (tp todoPersistence) Create(trans *gorp.Transaction, t *model.Todo) error {
	return trans.Insert(t)
}

func (tp todoPersistence) Update(trans *gorp.Transaction, t *model.Todo) error {
	_, err := trans.Update(t)
	return err
}

func (tp todoPersistence) Delete(trans *gorp.Transaction, t *model.Todo) error {
	_, err := trans.Delete(t)
	return err
}
