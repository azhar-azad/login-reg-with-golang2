package main

import (
	"awesomeProject/login-reg-cookie-demo/common"
	"github.com/gorilla/mux"
	"net/http"
)

var myRouter = mux.NewRouter()

func main() {

	myRouter.HandleFunc("/", common.LoginPageHandler).Methods("GET")

	myRouter.HandleFunc("/index", common.IndexPageHandler).Methods("GET")
	myRouter.HandleFunc("/login", common.LoginHandler).Methods("POST")

	myRouter.HandleFunc("/register", common.RegisterPageHandler).Methods("GET")
	myRouter.HandleFunc("/register", common.RegisterHandler).Methods("POST")

	myRouter.HandleFunc("/logout", common.LogoutHandler).Methods("POST")

	//http.Handle("/", myRouter)
	http.ListenAndServe(":8010", myRouter)
}
