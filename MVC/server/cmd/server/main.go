package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"server/controller"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"

	"gopkg.in/gorp.v2"
)

func main() {
	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(os.Stdout)

	datasource := os.Getenv("DB_DATASOURCE")

	db, err := sql.Open("mysql", datasource)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	dbm := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	defer dbm.Db.Close()

	r := mux.NewRouter()
	todoCtrl := controller.NewTodo(dbm)
	r.Methods(http.MethodGet).Path("/todos").HandlerFunc(todoCtrl.Index)
	r.Methods(http.MethodGet).Path("/todo/{id}").HandlerFunc(todoCtrl.Show)
	r.Methods(http.MethodPost).Path("/todo").HandlerFunc(todoCtrl.Create)
	r.Methods(http.MethodPut).Path("/todo/{id}").HandlerFunc(todoCtrl.Update)
	r.Methods(http.MethodDelete).Path("/todo/{id}").HandlerFunc(todoCtrl.Delete)

	addr := os.Getenv("PORT")
	log.Printf("Listening on port %s", addr)
	err = http.ListenAndServe(fmt.Sprintf(":%s", addr), handlers.CombinedLoggingHandler(os.Stdout, r))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
