package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

//Establishes a connection to the PostgreSQL database
func Connect(dsn string)(*sql.DB, error){
	db,err:=sql.Open("postgres", dsn)
	if err !=nil{
		return nil,err
	}

	//Test the connection pool
	if err:=db.Ping(); err!=nil{
		return nil,err
	}
	return db,nil
}

//runs SQL migrations files in the migrations directory
func Migrate(db *sql.DB)error{
	//Get migration files
	files,err:=os.ReadDir("migrations")
	if err!=nil{
		return fmt.Errorf("failed to read migrations directory: %v", err)
	}

	for _,file:= range files{
		if filepath.Ext(file.Name())==".sql"{
			migrationPath:=filepath.Join("migrations", file.Name())
			log.Printf("Running migration: %s", migrationPath)
			//Read the migration file
			content,err:=os.ReadFile(migrationPath)
			if err!=nil{	
				return fmt.Errorf("failed to read migration file %s: %v", migrationPath, err)
			}
			//Execute the migration
			_,err=db.Exec(string(content))
			if err!=nil{
				return fmt.Errorf("failed to execute migration %s: %v", migrationPath, err)
			}
	}

}
	log.Println("Migrations completed successfully")
	return nil
}