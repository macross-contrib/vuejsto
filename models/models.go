package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DB   *sql.DB
	Path string
)

func init() {
	var err error
	Path = "storage.db"
	DB, err = sql.Open("sqlite3", Path)
	if err != nil {
		panic(err)
	}
	if DB == nil {
		panic("DB is nil!")
	} else {
		fmt.Println("Vuejsto already standby!")
	}
	migrate(DB)
}

func migrate(db *sql.DB) {
	sql := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NOT NULL,
			done INTEGER NOT NULL
		);
	`

	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

type TaskCollection struct {
	Tasks []Task `json:"items"`
}

func GetTasks() TaskCollection {
	sql := "SELECT * FROM tasks;"
	rows, err := DB.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	result := TaskCollection{}

	for rows.Next() {
		task := Task{}
		err := rows.Scan(&task.ID, &task.Name, &task.Done)
		if err != nil {
			panic(err)
		}
		result.Tasks = append(result.Tasks, task)
	}

	return result
}

func PostTask(name string) (int64, error) {
	sql := "INSERT INTO tasks(name, done) VALUES(?, 0)"

	stmt, err := DB.Prepare(sql)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(name)
	if err != nil {
		panic(err)
	}

	return result.LastInsertId()
}

func PutTask(task Task) (int64, error) {
	var sql string

	sql = "UPDATE tasks SET name = ?, done = ? WHERE id = ?"

	stmt, err := DB.Prepare(sql)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	done := 0
	if task.Done {
		done = 1
	}

	result, err := stmt.Exec(task.Name, done, task.ID)

	if err != nil {
		panic(err)
	}

	return result.LastInsertId()
}

func DeleteTask(id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"

	stmt, err := DB.Prepare(sql)
	if err != nil {
		panic(err)
	}

	result, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}

	return result.RowsAffected()
}
