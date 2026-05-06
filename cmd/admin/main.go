package main

import (
	"github.com/abg216te/tgbt/internal/app/admin/handler"
	"github.com/abg216te/tgbt/internal/infrastructure/http"
	"github.com/abg216te/tgbt/internal/repository/memory"
	"github.com/abg216te/tgbt/internal/service"
)

func main() {
	repo := memory.NewUserRepo()
	service := service.NewUserService(repo)

	server := http.NewServer()

	userHandler := handler.NewUserHandler(service)

	server.Engine().GET("/users", userHandler.GetUsers)

	server.Run(":8080")
}
