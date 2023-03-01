package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Task struct {
	ID      uint64 `json:"id"`
	Todo    string `json:"todo"`
	Tanggal string `json:"date"`
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("pgx", "postgres://postgres:arvianta750@localhost:5432/ToDo_RPL")
	if err != nil {
		return nil, err
	}

	fmt.Println("Database connection established.")
	return db, nil
}

func InsertToDB(db *sql.DB, task Task) (*Task, error) {
	rows, err := db.Query("INSERT INTO todos (task, tanggal) VALUES ($1, $2) RETURNING id, task, tanggal", task.Todo, task.Tanggal)
	if err != nil {
		return nil, err
	}

	rows.Next()

	result := Task{}

	rows.Scan(&result.ID, &result.Todo, &result.Tanggal)
	return &result, nil
}

func GetAll(db *sql.DB) ([]Task, error) {
	var result []Task

	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task Task
		rows.Scan(&task.ID, &task.Todo, &task.Tanggal)

		result = append(result, task)
	}

	return result, nil
}

func DeleteFromDB(db *sql.DB, id int) error {
	fmt.Println(id)
	_, err := db.Query("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func deleteTODO(w http.ResponseWriter, r *http.Request) {
	db, err := Connect()
	if err != nil {
		fmt.Println(err)
	}

	id, err := strconv.Atoi(r.URL.Path[len("/Deletetask/"):])
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	err = DeleteFromDB(db, id)
	if err != nil {
		fmt.Println(err)
	}

	jsonMap := map[string]interface{}{
		"message": "Data berhasil dihapus",
	}

	b, err := json.Marshal(jsonMap)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(b))
}

func insertTODO(w http.ResponseWriter, r *http.Request) {
	db, err := Connect()
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	reqBody, _ := ioutil.ReadAll(r.Body)
	var post Task
	json.Unmarshal(reqBody, &post)

	json.NewEncoder(w).Encode(post)

	res, err := InsertToDB(db, post)
	if err != nil {
		fmt.Println(err)
	}

	jsonMap := map[string]interface{}{
		"message": "Data berhasil ditambah",
	}

	fmt.Println(res)

	b, err := json.Marshal(jsonMap)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(b))
}

func getTODO(w http.ResponseWriter, r *http.Request) {

	db, err := Connect()
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	res, err := GetAll(db)
	if err != nil {
		fmt.Println(err)
	}

	jsonMap := map[string]interface{}{
		"data": res,
	}

	b, err := json.Marshal(jsonMap)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(b))
}

func handleReqs() {

	http.HandleFunc("/Todolist", getTODO)
	http.HandleFunc("/Addtasks", insertTODO)
	http.HandleFunc("/Deletetask/", deleteTODO)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
func main() {
	handleReqs()
}
