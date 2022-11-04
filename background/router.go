package main

import (
	"net/http"

	"github.com/YuminosukeSato/slopegram/background/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func newRouter() *echo.Echo {
	e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000/", "http://localhost:3000/"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
    e.POST("/signup", handler.Signup)
    e.POST("/login", handler.Login)

    api := e.Group("/api")
    api.Use(middleware.JWTWithConfig(handler.Config))
    api.GET("/todos", handler.GetTodos)
    api.POST("/todos", handler.AddTodo)
    api.DELETE("/todos/:id", handler.DeleteTodo)
    api.PUT("/todos/:id/completed", handler.UpdateTodo)

	return e
}