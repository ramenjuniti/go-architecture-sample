package service

import (
	"server/model"

	"gopkg.in/gorp.v2"
)

func CreateTodo(dbm *gorp.DbMap, t *model.Todo) error {
	trans, err := dbm.Begin()
	if err != nil {
		return err
	}

	if err = model.CreateTodo(trans, t); err != nil {
		return err
	}

	if err = trans.Commit(); err != nil {
		return err
	}

	return nil
}

func UpdateTodo(dbm *gorp.DbMap, t *model.Todo) error {
	trans, err := dbm.Begin()
	if err != nil {
		return err
	}

	if err = model.UpdateTodo(trans, t); err != nil {
		return err
	}

	if err = trans.Commit(); err != nil {
		return err
	}

	return nil
}

func DeleteTodo(dbm *gorp.DbMap, t *model.Todo) error {
	trans, err := dbm.Begin()
	if err != nil {
		return err
	}

	if err = model.DeleteTodo(trans, t); err != nil {
		return err
	}

	if err = trans.Commit(); err != nil {
		return err
	}

	return nil
}
