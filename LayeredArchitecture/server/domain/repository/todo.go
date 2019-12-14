package repository

import (
	"server/domain/model"

	"gopkg.in/gorp.v2"
)

type TodoRepository interface {
	GetAll(dbm *gorp.DbMap) ([]model.Todo, error)
	GetByID(dbm *gorp.DbMap, id int64) (*model.Todo, error)
	Create(dbm *gorp.Transaction, t *model.Todo) error
	Update(dbm *gorp.Transaction, t *model.Todo) error
	Delete(dbm *gorp.Transaction, t *model.Todo) error
}
