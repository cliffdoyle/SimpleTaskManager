package database

import(
	"database/sql"
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