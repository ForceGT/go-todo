package main

import (
	"context"
	"go_todo/config"
	"go_todo/dao"
	"go_todo/db"
	"go_todo/server/controller"
	"go_todo/server/middleware"
	"go_todo/server/routes"
	"go_todo/server/validator"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {

	startServer()
}

func startServer() {
	db, err := db.ConnectToDB()
	if err != nil {
		log.Printf("error connecting to database: %v", err)
	}

	e := echo.New()
	e.Validator = validator.NewEchoRequestValidator()

	roleDao := dao.NewRoleDao(db)
	userDao := dao.NewUserDao(db)
	todoDao := dao.NewTodoDao(db)

	roleController := controller.NewRoleController(roleDao)
	userController := controller.NewUserController(userDao)
	todoController := controller.NewTodoController(todoDao)

	authController := controller.NewAuthController()
	routes.Auth(e, userController, authController)

	jwtMiddleware := middleware.GetJWTMiddleware(userController)

	api := e.Group("/api", jwtMiddleware)

	routes.Role(api, roleController)
	routes.User(api, userController)
	routes.Todo(api, todoController)

	httpServer := &http.Server{
		Addr: ":8080",
	}

	e.Debug = true

	go func() {
		log.Printf("Starting server  at %v", config.PortNumber)
		if err := e.StartServer(httpServer); err != nil {
			log.Panicf("Shutting down the server due to %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Println("Error shutting down server gracefully")
	}

}
