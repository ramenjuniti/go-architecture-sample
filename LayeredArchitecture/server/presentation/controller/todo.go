package controller

import (
	"encoding/json"
	"net/http"
	"server/application"
	"server/httputil"
	"strconv"

	"github.com/gorilla/mux"

	"server/presentation/view"

	"gopkg.in/gorp.v2"
)

type reqTodo struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type TodoController interface {
	Index(http.ResponseWriter, *http.Request)
	Show(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type todoController struct {
	todoApplication application.TodoApplication
	dbm             *gorp.DbMap
}

func NewTodoController(ta application.TodoApplication, dbm *gorp.DbMap) *todoController {
	return &todoController{
		todoApplication: ta,
		dbm:             dbm,
	}
}

func (tc *todoController) Index(w http.ResponseWriter, r *http.Request) {
	todos, err := tc.todoApplication.GetAll(tc.dbm)
	if err != nil {
		view.RenderJSON(w, http.StatusInternalServerError, &httputil.HTTPError{Message: err.Error()})
		return
	}

	view.RenderJSON(w, http.StatusOK, todos)
}

func (tc *todoController) Show(w http.ResponseWriter, r *http.Request) {
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

	todo, err := tc.todoApplication.GetByID(tc.dbm, id)
	if err != nil {
		view.RenderJSON(w, http.StatusInternalServerError, &httputil.HTTPError{Message: err.Error()})
		return
	}

	view.RenderJSON(w, http.StatusOK, todo)
}

func (tc *todoController) Create(w http.ResponseWriter, r *http.Request) {
	newTodo := &reqTodo{}
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		view.RenderJSON(w, http.StatusBadRequest, &httputil.HTTPError{Message: err.Error()})
		return
	}

	err := tc.todoApplication.Create(tc.dbm, newTodo.Title, newTodo.Body)
	if err != nil {
		view.RenderJSON(w, http.StatusInternalServerError, &httputil.HTTPError{Message: err.Error()})
		return
	}

	view.RenderJSON(w, http.StatusCreated, newTodo)
}

func (tc *todoController) Update(w http.ResponseWriter, r *http.Request) {
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

	updateTodo := &reqTodo{ID: id}
	if err := json.NewDecoder(r.Body).Decode(&updateTodo); err != nil {
		view.RenderJSON(w, http.StatusBadRequest, &httputil.HTTPError{Message: err.Error()})
		return
	}

	err = tc.todoApplication.Update(tc.dbm, updateTodo.ID, updateTodo.Title, updateTodo.Body)
	if err != nil {
		view.RenderJSON(w, http.StatusInternalServerError, &httputil.HTTPError{Message: err.Error()})
		return
	}

	view.RenderJSON(w, http.StatusOK, updateTodo)
}

func (tc *todoController) Delete(w http.ResponseWriter, r *http.Request) {
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

	err = tc.todoApplication.Delete(tc.dbm, id)
	if err != nil {
		view.RenderJSON(w, http.StatusInternalServerError, &httputil.HTTPError{Message: err.Error()})
		return
	}

	view.RenderJSON(w, http.StatusNoContent, nil)
}
