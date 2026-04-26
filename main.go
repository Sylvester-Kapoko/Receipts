package main

import (
	"github.com/Sylvester-Kapoko/Receipts/service"
	"github.com/Sylvester-Kapoko/Receipts/handler"
	"github.com/Sylvester-Kapoko/Receipts/repository"
                "fmt"
	"net/http"

)



func main() {
	repo := repository.NewUserRepository()  //lowest layer
	service := service.NewUserService(repo) // depends on repo
	handler := handler.NewHandler(service)  //depends on service

	http.Handle("/user", handler)
	fmt.Println("Server is running. Open your browser at http://localhost:8080/user")
                err := http.ListenAndServe(":8080", nil)
                if err != nil {
                               fmt.Println("Error starting server:", err)
                }

}
