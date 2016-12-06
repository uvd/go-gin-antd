package main

import (
	"github.com/gorilla/mux"
)

func initRouter() *mux.Router {
	mux := mux.NewRouter()
	return mux
}
