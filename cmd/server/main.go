package main

import (
	"database/sql"

	"github.com/SilentCanary/Account-service-Gofr-RBAC/internal/handler"
	"github.com/SilentCanary/Account-service-Gofr-RBAC/internal/service"
	"github.com/SilentCanary/Account-service-Gofr-RBAC/internal/store"
	"gofr.dev/pkg/gofr"

	_ "github.com/lib/pq"
)

func main() {
	app := gofr.New()

	db, err := sql.Open("postgres", "host=localhost port=5432 user=admin password=admin dbname=auth_service sslmode=disable")

	if err != nil {
		panic(err)
	}
	user_store := store.NewUserStore(db)
	user_handler := handler.NewUserHandler(service.NewUserService(user_store))
	auth_handler := handler.NewAuthHandler(service.NewAuthService(user_store))

	app.GET("/health", handler.Health)
	app.POST("/register", user_handler.RegisterUser)
	app.POST("/login", auth_handler.LoginUser)
	app.Run()
}
