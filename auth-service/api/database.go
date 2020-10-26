package api

import (
	"database/sql"
	"log"
	"time"

	//MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

//DB represents the connection to the MySQL database
var (
	DB *sql.DB
)

//InitDB creates the MySQL database connection
func InitDB() *sql.DB {

	log.Println("attempting connections")

	var err error

	// Open a SQL connection to the docker container hosting the database server
	// Assign the connection to the "DB" variable
	// Look at how it's done in the other microservices!
	dbType := "mysql"
	username := "root"
	password := "root"
	ipAddress := "tcp(172.28.1.2:3306)"
	dbName := "/auth"
	// "YOUR CODE HERE"
	DB, err = sql.Open(dbType, username+":"+password+"@"+ipAddress+dbName) // possibly add DB?parseTime=true but I dont think it's necessary

	if err != nil {
		log.Println("couldnt connect")
		panic(err.Error())
	}

	err = DB.Ping()
	if err != nil {
		log.Println("couldnt ping")
		panic(err.Error())
	}
	_, err = DB.Query("SELECT * FROM users")
	for err != nil {
		log.Println("couldnt connect, waiting 20 seconds before retrying")
		time.Sleep(20 * time.Second)
		// Connect again, use the same connection function as you did above ^
		// YOUR CODE HERE
		DB, err = sql.Open(dbType, username+":"+password+"@"+ipAddress+dbName)
	}

	return DB
}
