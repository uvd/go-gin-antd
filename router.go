package main

import (
	"antd-admin-golang/controllers"

	"github.com/gorilla/mux"
)

func initRouter() *mux.Router {
	mux := mux.NewRouter()

	//user begin router
	userCtr := new(controllers.UserController)
	mux.HandleFunc("/api/user/login", userCtr.Login)
	mux.HandleFunc("/api/user/regist", userCtr.Regist)
	mux.HandleFunc("/api/user/list", userCtr.List)
	mux.HandleFunc("/api/user/update", userCtr.Update)
	return mux
}
