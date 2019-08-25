package main
import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"./delivery"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.GET("/todo", delivery.GetTodoData)
	e.POST("/todo", delivery.CreateTodo)
	e.PUT("/todo", delivery.UpdateTodo)
	e.Logger.Fatal(e.Start(":1323"))
}