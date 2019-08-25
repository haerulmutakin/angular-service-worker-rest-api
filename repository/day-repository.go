package repository

func GetTodo() string {
	return "SELECT id, name, completed FROM go_test.todo ORDER BY completed ASC ,date_update DESC"
}

func AddTodo() string {
	return "INSERT INTO todo (name, completed, date_update) VALUES (?, ?, ?)"
}

func DeleteTodo() string {
	return "UPDATE todo SET completed=1, date_update=? WHERE id = ?"
}