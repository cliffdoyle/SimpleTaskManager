package models


type Task struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status string `json:"completed"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`

}