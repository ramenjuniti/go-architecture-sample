package controller

import (
	"encoding/json"
	"net/http"
	"server/httputil"
	"server/service"
	"strconv"

	"github.com/gorilla/mux"

	"server/model"

	"server/view"

	"gopkg.in/gorp.v2"
)

type Todo struct {
	dbm *gorp.DbMap
}

func NewTodo(dbm *gorp.DbMap) *Todo {
	dbm.AddTableWithName(model.Todo{}, "todo")

	return &Todo{dbm: dbm}
}

func (t *Todo) Index(w http.ResponseWriter, r *http.Request) {
	todos, err := model.AllTodos(t.dbm)
	if err != nil {
		view.RenderJSON(w, http.StatusInternalServerError, &httputil.HTTPError{Message: err.Error()})
		return
	}

	view.RenderJSON(w, http.StatusOK, todos)
}

func (t *Todo) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	idStr, ok := vars["id"]
	if !ok {
		view.RenderJSON(w, http.StatusBadRequest, &httputil.HTTPError{Message: "invalid path parameter"})
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		view.RenderJSON(w, http.StatusBadRequest, &httputil.HTTPError{Message: err.Error()})
		return
	}

	todo, err := model.FindTodo(t.dbm, id)
	if err != nil {
		view.RenderJSON(w, http.StatusInternalServerError, &httputil.HTTPError{Message: err.Error()})
		return
	}

	view.RenderJSON(w, http.StatusOK, todo)
}

func (t *Todo) Create(w http.ResponseWriter, r *http.Request) {
	newTodo := &model.Todo{}
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		view.RenderJSON(w, http.StatusBadRequest, &httputil.HTTPError{Message: err.Error()})
		return
	}

	err := service.CreateTodo(t.dbm, newTodo)
	if err != nil {
		view.RenderJSON(w, http.StatusInternalServerError, &httputil.HTTPError{Message: err.Error()})
		return
	}

	view.RenderJSON(w, http.StatusCreated, newTodo)
}

func (t *Todo) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	idStr, ok := vars["id"]
	if !ok {
		view.RenderJSON(w, http.StatusBadRequest, &httputil.HTTPError{Message: "invalid path parameter"})
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		view.RenderJSON(w, http.StatusBadRequest, &httputil.HTTPError{Message: err.Error()})
		return
	}

	reqTodo := &model.Todo{ID: id}
	if err := json.NewDecoder(r.Body).Decode(&reqTodo); err != nil {
		view.RenderJSON(w, http.StatusBadRequest, &httputil.HTTPError{Message: err.Error()})
		return
	}

	err = service.UpdateTodo(t.dbm, reqTodo)
	if err != nil {
		view.RenderJSON(w, http.StatusBadRequest, &httputil.HTTPError{Message: err.Error()})
		return
	}

	view.RenderJSON(w, http.StatusOK, reqTodo)
}

func (t *Todo) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	idStr, ok := vars["id"]
	if !ok {
		view.RenderJSON(w, http.StatusBadRequest, &httputil.HTTPError{Message: "invalid path parameter"})
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		view.RenderJSON(w, http.StatusBadRequest, &httputil.HTTPError{Message: err.Error()})
		return
	}

	reqTodo := &model.Todo{ID: id}

	err = service.DeleteTodo(t.dbm, reqTodo)
	if err != nil {
		view.RenderJSON(w, http.StatusBadRequest, &httputil.HTTPError{Message: err.Error()})
		return
	}

	view.RenderJSON(w, http.StatusNoContent, nil)
}
