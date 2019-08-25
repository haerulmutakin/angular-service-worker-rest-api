package delivery

import (
	"net/http"
	"github.com/labstack/echo"
	"../models"
	"../service"
)

func GetTodoData(c echo.Context) error {
	dayData, _ := service.GetTodoData()

	if dayData == nil {
		result := echo.Map{
			"count": 0,
			"data": []int{},

		}
		return c.JSON(http.StatusOK, result)
	}
	result := echo.Map{
		"count": len(dayData),
		"data":  dayData,
	}
	return c.JSON(http.StatusOK, result)
}

func CreateTodo(c echo.Context) error {
	datadiri := new(model.TodoCreate)
	if err := c.Bind(datadiri); err != nil {
		result := echo.Map{"message ": "Data kosong!"}
		return c.JSON(http.StatusBadRequest, result)
	}
	create := service.CreateTodo(datadiri.NAME, datadiri.COMPLETED)
	if create != nil {
		result := echo.Map{"message" : "Gagal tambah data"}
		return c.JSON(http.StatusOK, result)
	}
	result := echo.Map{"message": "Berhasil tambah data"}

	return c.JSON(http.StatusOK, result)
}

func UpdateTodo(c echo.Context) error {
	datadiri := new(model.Todo)
	if err := c.Bind(datadiri); err != nil {
		result := echo.Map{"message ": "Data kosong!"}
		return c.JSON(http.StatusOK, result)
	}
	delete := service.DeleteTodo(datadiri.ID)

	if delete != nil {
		result := echo.Map{"message ": " Gagal ubah data"}
		return c.JSON(http.StatusOK, result)
	}
	result := echo.Map{"message": "Berhasil ubah data"}

	return c.JSON(http.StatusCreated, result)
}