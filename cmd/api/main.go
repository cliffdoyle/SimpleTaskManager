package main

import (
	"log"
	"net/http"
	"os"
	"github.com/cliffdoyle/SimpleTaskManager.git/internal/config"
	"github.com/cliffdoyle/SimpleTaskManager.git/internal/database"
	"github.com/cliffdoyle/SimpleTaskManager.git/internal/router"

)


func main(){
	//Ensure we're in the project root for relative paths
	ensureProjectRoot()

	//Load configuration
	cfg:=config.LoadConfig()

	//connect to database
	db,err:=database.Connect(cfg.GetDSN())
	if err!=nil{
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()
	log.Println("Connected to database")

	//run migrations
	if err:=database.Migrate(db); err!=nil{
		log.Fatalf("Error running migrations: %v", err)
	}
	log.Println("Migrations completed")
	//setup router
	r:=router.SetupRouter(db)
	//setup static file server for frontend
	fs:=http.FileServer(http.Dir("web"))
	r.PathPrefix("/").Handler(fs)

	//start server
	log.Printf("Starting server on %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))

}


func ensureProjectRoot() {
    // Check if migrations directory exists
    if _, err := os.Stat("migrations"); os.IsNotExist(err) {
        // Try going up one directory (if running from /cmd/api)
        if err := os.Chdir("../.."); err != nil {
            log.Printf("Warning: Could not change to project root directory: %v", err)
        }
    }
    
    // Log current directory
    dir, err := os.Getwd()
    if err == nil {
        log.Printf("Working directory: %s", dir)
    }
    
    // Verify migrations directory exists
    if _, err := os.Stat("migrations"); os.IsNotExist(err) {
        log.Printf("Warning: migrations directory not found. Migrations might fail.")
    }
}