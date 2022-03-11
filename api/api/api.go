package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

type Store interface {
	ToDoStore
	UserStore
}

func StartServer(s Store, port string) {
	router := mux.NewRouter()
	auth := NewAuthenticator(s)
	amw := NewAuthenticationMiddleware(auth)

	authRouter := router.NewRoute().Subrouter()
	authRouter.Use(amw.Middleware)

	todoHandler := NewTodoHandler(s, auth)
	userHandler := NewUserHandler(s, auth)

	router.HandleFunc("/login", userHandler.login).Methods("POST")
	authRouter.HandleFunc("/user", userHandler.showUser).Methods("GET")

	authRouter.HandleFunc("/todo", todoHandler.listTodos).Methods("GET")
	authRouter.HandleFunc("/todo/{id:[\\d]+}", todoHandler.showTodo).Methods("GET")
	authRouter.HandleFunc("/todo", todoHandler.createTodo).Methods("POST")
	authRouter.HandleFunc("/todo/{id:[\\d]+}", todoHandler.updateTodo).Methods("PUT")
	authRouter.HandleFunc("/todo/{id:[\\d]+}", todoHandler.deleteTodo).Methods("DELETE")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost,http.MethodPut, http.MethodDelete},
		AllowedHeaders: []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	fmt.Println(" Server Starting on Port "+port)
	if err := http.ListenAndServe(":"+port, handler ); err != nil {
		panic(err)
	}
}

func writeJson(res http.ResponseWriter, content interface{}, status int) {
	jsonBytes, _ := json.MarshalIndent(content, "", "  ")
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	res.Write(jsonBytes)
}

func jsonErr(res http.ResponseWriter, error string, status int) {
	writeJson(res, struct {
		Error string `json:"error"`
	}{
		Error: error,
	}, status)
}