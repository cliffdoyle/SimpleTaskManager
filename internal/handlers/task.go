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

//CreateTask creates a new task
func (t *TaskHandler)CreateTask(w http.ResponseWriter, r *http.Request){
	//parse the request body
	var task models.Task
	err:=json.NewDecoder(r.Body).Decode(&task)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//validate the task
	if task.Title==""{
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}	
	if task.Status==""{
		task.Status="pending"
	}

	//insert the task into the database
	err=t.DB.QueryRow(`INSERT INTO tasks(title,description,status,created_at,updated_at) VALUES ($1,$2,$3,NOW(),NOW()) RETURNING id,title,description,status,created_at,updated_at`,task.Title,task.Description,task.Status,task.CreatedAt,task.UpdatedAt).Scan(&task.ID,&task.Title,&task.Description,&task.Status,&task.CreatedAt,&task.UpdatedAt)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Return Json response
	w.Header().Set("content-type","application/jsom")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)

	// http.Error(w, "Not implemented", http.StatusNotImplemented)
}

//UpdateTask updates a task
func (t *TaskHandler)UpdateTask(w http.ResponseWriter, r *http.Request){
	//Get Id from the url
	params:=mux.Vars(r)
	id:=params["id"]
	//parse the request body			
	var task models.Task
	err:=json.NewDecoder(r.Body).Decode(&task)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//validate the task
	if task.Title==""{		
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	//update the task in the database
	err=t.DB.QueryRow(`UPDATE tasks SET title=$1,description=$2,status=$3,updated_at=NOW() WHERE id=$4 RETURNING id,title,description,status,created_at,updated_at`,task.Title,task.Description,task.Status,id).Scan(&task.ID,&task.Title,&task.Description,&task.Status,&task.CreatedAt,&task.UpdatedAt)
	if err!=nil{
		if err==sql.ErrNoRows{
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err!=nil{
		if err==sql.ErrNoRows{
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}else{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	//Return Json response
	w.Header().Set("content-type","application/jsom")
	json.NewEncoder(w).Encode(task)
}

//deleteTask deletes a task
func (t *TaskHandler)DeleteTask(w http.ResponseWriter, r *http.Request){
	//Get Id from the url
	params:=mux.Vars(r)
	id:=params["id"]

	//Delete the task from the database
	result,err:=t.DB.Exec("DELETE FROM tasks WHERE id=$1",id)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//check if the task was deleted
	rowsAfected,err:=result.RowsAffected()
	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAfected==0{
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	//Return Json response
	w.Header().Set("content-type","application/jsom")
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(map[string]string{"message":"Task deleted successfully"})
}