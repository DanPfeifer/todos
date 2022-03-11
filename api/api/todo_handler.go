package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"todos/resources"
)

type DeleteResponse struct {
	Message string `json:"message"`
}

type TodoRequest struct {
	Title string `json:"title"`
	Message string `json:"message"`
	Priority int `json:"priority"`
	Status int `json:"status"`
	DueDate string `json:"due_date"`
}

type TodoResponse struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Message string `json:"message"`
	Priority int `json:"priority"`
	Status int `json:"status"`
	UserID int `json:"user_id"`
	InsertDate string `json:"insert_date"`
	DueDate string `json:"due_date"`
}

type TodosResponses struct {
	Todos []TodoResponse `json:"todos"`
}

type ToDoStore interface {
	GetAllTodos(user resources.User) (results []resources.ToDo, err error)
	GetOneTodo(id int) (todo resources.ToDo, err error)
	CreateTodo(t resources.ToDo) (todo resources.ToDo, err error)
	UpdateTodo(t resources.ToDo, id int) (todo resources.ToDo, err error)
	DeleteTodo(id int) ( err error)
}


type todoHandler struct {
	d ToDoStore
	auth Authenticator
}

func NewTodoHandler(db ToDoStore, auth Authenticator) *todoHandler {
	return &todoHandler{
		d: db,
		auth: auth,
	}
}

func (th *todoHandler) listTodos(res http.ResponseWriter, req *http.Request) {

	user, err := th.auth.GetUser(req)
	if err != nil {
		fmt.Print(err)
		jsonErr(res, "Something went wrong", http.StatusInternalServerError)
		return
	}
	todos, err := th.d.GetAllTodos(user)

	if err != nil {
		fmt.Print(err)
		jsonErr(res, "Something went wrong", http.StatusInternalServerError)
		return
	}
	var tdr []TodoResponse

	for _, t := range todos {
		tdr = append(tdr, mapTodoResponse(t))
	}

	resBody := TodosResponses {
		tdr,
	}

	writeJson(res, resBody, http.StatusOK)
}

func (th *todoHandler) showTodo(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		jsonErr(res, "Something went wrong", http.StatusInternalServerError)
		return
	}
	t, err := th.d.GetOneTodo(id)

	if err != nil {
		jsonErr(res, "Could not find todo", http.StatusBadRequest)
		return
	}

	resBody := mapTodoResponse(t)

	writeJson(res, resBody, http.StatusOK)
}

func (th *todoHandler) createTodo(res http.ResponseWriter, req *http.Request) {
	user, err := th.auth.GetUser(req)
	if err != nil {
		jsonErr(res, "Something went wrong", http.StatusInternalServerError)
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		jsonErr(res, "invalid json request", http.StatusBadRequest)
		return
	}

	var tr TodoRequest
	if err := json.Unmarshal(body, &tr); err != nil {
		fmt.Print(err)
		jsonErr(res, "invalid json request", http.StatusBadRequest)
		return
	}

	t := resources.ToDo {
		Title: tr.Title,
		Message: tr.Message,
		Priority: tr.Priority,
		Status: tr.Status,
		UserID: user.ID,
		DueDate: tr.DueDate,
	}

	results, err := th.d.CreateTodo(t)

	if err != nil {
		fmt.Print(err)
		jsonErr(res, "Could not Create todo", http.StatusInternalServerError)
		return
	}

	resBody :=  mapTodoResponse(results)

	writeJson(res, resBody, http.StatusOK)
	return
}

func (th *todoHandler) updateTodo(res http.ResponseWriter, req *http.Request) {

	user, err := th.auth.GetUser(req)
	if err != nil {
		jsonErr(res, "Something went wrong", http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Print(err)
		jsonErr(res, "invalid json request", http.StatusBadRequest)
		return
	}

	var tr TodoRequest
	if err := json.Unmarshal(body, &tr); err != nil {
		fmt.Print(err)
		jsonErr(res, "invalid json request", http.StatusBadRequest)
		return
	}

	t := resources.ToDo {
		Title: tr.Title,
		Message: tr.Message,
		Priority: tr.Priority,
		Status: tr.Status,
		UserID: user.ID,
		DueDate: tr.DueDate,
	}

	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		jsonErr(res, "Something went wrong", http.StatusInternalServerError)
		return
	}

	results, err := th.d.UpdateTodo(t, id)

	if err != nil {
		jsonErr(res, "Could not Update todo", http.StatusInternalServerError)
		return
	}

	resBody :=  mapTodoResponse(results)

	writeJson(res, resBody, http.StatusOK)
}

func (th *todoHandler) deleteTodo(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		jsonErr(res, "Something went wrong", http.StatusInternalServerError)
	}

	err = th.d.DeleteTodo(id)

	if err != nil {
		jsonErr(res, "Could not Delete todo", http.StatusInternalServerError)
		return
	}

	resBody := DeleteResponse{
		Message: "ToDo has been deleted",
	}

	writeJson(res, resBody, http.StatusOK)
}

func mapTodoResponse(t resources.ToDo) TodoResponse {
	return TodoResponse{
		ID: t.ID,
		Title: t.Title,
		Message: t.Message,
		Priority: t.Priority,
		Status: t.Status,
		UserID: t.UserID,
		InsertDate: t.InsertDate,
		DueDate: t.DueDate ,
	}
}