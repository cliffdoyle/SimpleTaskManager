package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/cliffdoyle/SimpleTaskManager.git/internal/models"
	"github.com/gorilla/mux"
)


type TaskHandler struct{
	DB *sql.DB
}

//NewTaskHandler creates a new TaskHandler
func NewTaskHandler(db *sql.DB)*TaskHandler{
	return &TaskHandler{DB: db}
}

//GetTasks returns all tasks
func (t *TaskHandler)GetTasks(w http.ResponseWriter, r *http.Request){
	//TODO implement me
	//querry tasks
	rows,err:=t.DB.Query(`SELECT id, title, description,status,created_at,updated_at
	FROM tasks
	ORDER BY created_at DESC
	`)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	//create a slice of tasks
	var tasks []models.Task
	//loop through the rows
	for rows.Next(){
		var task models.Task
		//scan the row into the task
		err:=rows.Scan(&task.ID,&task.Title,&task.Description,&task.Status,&task.CreatedAt,&task.UpdatedAt)
		if err!=nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//append the task to the slice
		tasks=append(tasks, task)
	}

	//Return Json response
	w.Header().Set("content-type","application/jsom")
	json.NewEncoder(w).Encode(tasks)

	
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

//GetTask returns a single task
func (t *TaskHandler)GetTask(w http.ResponseWriter, r *http.Request){
	//TODO implement me
	params:=mux.Vars(r)
	id:=params["id"]

	//querry task
	var task models.Task

	err:=t.DB.QueryRow(`SELECT id, title, description,status,created_at,updated_at
	FROM tasks
	WHERE id=$1`,id).Scan(&task.ID,&task.Title,&task.Description,&task.Status,&task.CreatedAt,&task.UpdatedAt)
	if err!=nil{
		if err==sql.ErrNoRows{
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Return Json response
	w.Header().Set("content-type","application/jsom")
	json.NewEncoder(w).Encode(task)

	http.Error(w, "Not implemented", http.StatusNotImplemented)
}