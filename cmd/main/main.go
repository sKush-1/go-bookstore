package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sKush-1/bookstore/pkg/routes"
	"github.com/sKush-1/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
