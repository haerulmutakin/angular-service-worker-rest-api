package service

import (
	"log"
	"fmt"
	"time"
	"../models"
	"../conn"
	"../repository"
)

func GetTodoData() ([]model.Todo, error) {
	dbConn := connection.DBConnect()

	preparedQuery := repository.GetTodo()
	rows, err := dbConn.Query(preparedQuery)
	if err != nil {
		panic(err)
	}
	var todos []model.Todo
	var todo model.Todo

	for rows.Next() {
		rows.Scan(&todo.ID, &todo.NAME, &todo.COMPLETED)
		todos = append(todos, todo)
	}

	defer dbConn.Close()
	return todos, nil
}

func CreateTodo(todoName string, todoCompleted int) error {
	conn := connection.DBConnect()

	time := time.Now()
	fmt.Println(time.Format("2006-01-02 15:04:05"))

	sqlQuery := repository.AddTodo()
	sqlStmt, errorDB := conn.Prepare(sqlQuery)
	if errorDB != nil {
		log.Fatal(errorDB)
	}
	_, errorExecute := sqlStmt.Exec(todoName, todoCompleted, time.Format("2006-01-02 15:04:05"))

	defer conn.Close()
	return errorExecute
}

func DeleteTodo(id int) error {
	conn := connection.DBConnect()

	time := time.Now()
	fmt.Println(time.Format("2006-01-02 15:04:05"))

	sqlQuery := repository.DeleteTodo()
	sqlStmt, errorDB := conn.Prepare(sqlQuery)
	if errorDB != nil {
		log.Fatal(errorDB)
	}
	_, errorExecute := sqlStmt.Exec(time.Format("2006-01-02 15:04:05"), id)

	defer conn.Close()
	return errorExecute
}