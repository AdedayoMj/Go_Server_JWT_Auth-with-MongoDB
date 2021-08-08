package main

import (
	"fmt"
	"net/http"
	config "servergo-jwt/config"
	"servergo-jwt/driver"
	"servergo-jwt/handler"

	"github.com/gorilla/mux"
)

func main() {
	driver.ConnectMongoDB(config.DB_USER, config.DB_PASS)
	router := mux.NewRouter()

	router.HandleFunc("/login", handler.Login)
	router.HandleFunc("/register", handler.Register)
	router.HandleFunc("/user", handler.GetUser)

	fmt.Println("Server running [:8000]")
	http.ListenAndServe(":8000", router)
}
