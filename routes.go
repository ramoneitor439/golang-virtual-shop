package main

import (
	"github.com/gorilla/mux"
	authController "mystore.com/controllers/authController"
	"mystore.com/middlewares"
)

func mapAuthController(mux *mux.Router) {
	mux.HandleFunc("auth/sign-in", authController.SignIn).Methods("POST")
	mux.HandleFunc("auth/sign-up", authController.SignUp).Methods("POST")

	authRoute := mux.PathPrefix("/auth").Subrouter()

	authRoute.Use(middlewares.AuthorizeMiddleware)

	authRoute.HandleFunc("/delete/{userId:[0-9]+}", authController.Delete).Methods("DELETE")
	authRoute.HandleFunc("/update/{userId:[0-9]+}", authController.Update).Methods("PUT")
}
