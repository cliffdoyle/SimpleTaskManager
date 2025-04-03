// internal/router/router.go
package router

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/cliffdoyle/SimpleTaskManager.git/internal/handlers"
	"github.com/gorilla/mux"
)

// SetupRouter sets up the HTTP router
func SetupRouter(db *sql.DB) *mux.Router {
    r := mux.NewRouter()
    
    // Create handlers
    taskHandler := handlers.NewTaskHandler(db)
    
    // API endpoints
    r.HandleFunc("/api/tasks", taskHandler.GetTasks).Methods("GET")
    r.HandleFunc("/api/tasks/{id}", taskHandler.GetTask).Methods("GET")
    r.HandleFunc("/api/tasks", taskHandler.CreateTask).Methods("POST")
    r.HandleFunc("/api/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
    r.HandleFunc("/api/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")
    
    // Health check
    r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        // Check database connection
        err := db.Ping()
        if err != nil {
            http.Error(w, "Database connection failed", http.StatusServiceUnavailable)
            return
        }
        
        // Return success response
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
    }).Methods("GET")
    
    // Add middleware for CORS
    r.Use(corsMiddleware)
    
    return r
}

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}