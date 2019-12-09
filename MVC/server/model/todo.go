package model

import (
	"gopkg.in/gorp.v2"
)

type Todo struct {
	ID    int64  `json:"id" db:"id, primarykey, autoincrement"`
	Title string `json:"title" db:"title"`
	Body  string `json:"body" db:"body"`
}

func AllTodos(dbm *gorp.DbMap) ([]Todo, error) {
	t := make([]Todo, 0)
	if _, err := dbm.Select(&t, `SELECT id, title, body FROM todo`); err != nil {
		return nil, err
	}
	return t, nil
}

func FindTodo(dbm *gorp.DbMap, id int64) (*Todo, error) {
	t := Todo{}
	if err := dbm.SelectOne(&t, `SELECT id, title, body FROM todo WHERE id = ?`, id); err != nil {
		return nil, err
	}
	return &t, nil
}

func CreateTodo(trans *gorp.Transaction, t *Todo) error {
	return trans.Insert(t)
}

func UpdateTodo(trans *gorp.Transaction, t *Todo) error {
	_, err := trans.Update(t)
	return err
}

func DeleteTodo(trans *gorp.Transaction, t *Todo) error {
	_, err := trans.Delete(t)
	return err
}
