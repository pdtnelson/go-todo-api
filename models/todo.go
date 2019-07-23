package models

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/jinzhu/gorm"
)

// ToDo represents a todo from the database
type ToDo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// Routes organizes the handlers for ToDo API routes
func Routes(db *gorm.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{ToDoID}", GetToDoByID(db))
	router.Delete("/{ToDoID}", DeleteToDoByID(db))
	router.Post("/", CreateToDo(db))
	router.Get("/", GetAllToDos(db))
	router.Patch("/{ToDoID}", UpdateToDoByID(db))
	return router
}

// GetToDoByID use an id to attempt to find a specific ToDo
func GetToDoByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ToDoID := chi.URLParam(r, "ToDoID")
		var todo ToDo
		db.Where("ID = ?", ToDoID).First(&todo)
		render.JSON(w, r, todo)
	}
}

// DeleteToDoByID use an id to attempt to delete a specific ToDo
func DeleteToDoByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ToDoID := chi.URLParam(r, "ToDoID")
		var todo ToDo
		db.Where("ID = ?", ToDoID).Delete(&todo)
		response := make(map[string]string)
		response["message"] = "Deleted ToDo successfully"
		render.JSON(w, r, response)
	}
}

// CreateToDo creates a new ToDo
func CreateToDo(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var newToDo ToDo
		err := decoder.Decode(&newToDo)
		if err != nil {
			panic(err)
		}
		db.NewRecord(newToDo)
		db.Create(&newToDo)
		response := make(map[string]string)
		response["message"] = "Created ToDo successfully"
		render.JSON(w, r, response)
	}
}

// GetAllToDos returns all ToDos that are currently stored
func GetAllToDos(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todos []ToDo
		db.Find(&todos)
		render.JSON(w, r, todos)
	}
}

// UpdateToDoByID updates a specific ToDo
func UpdateToDoByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ToDoID := chi.URLParam(r, "ToDoID")
		// var todo ToDo
		// db.Where("ID = ?", ToDoID).First(&todo)
		// todo.Completed = true
		// db.Save(&todo)
		// response := make(map[string]string)
		// response["message"] = "Deleted ToDo successfully"
		// render.JSON(w, r, response)

		decoder := json.NewDecoder(r.Body)
		var requestToDo ToDo
		err := decoder.Decode(&requestToDo)
		if err != nil {
			panic(err)
		}
		db.Model(&requestToDo).Where("ID = ?", requestToDo.ID).Update(ToDo{Completed: requestToDo.Completed, Title: requestToDo.Title, Description: requestToDo.Description})
		response := make(map[string]string)
		response["message"] = "Created ToDo successfully"
		render.JSON(w, r, response)
	}
}
