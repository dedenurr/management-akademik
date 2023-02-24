package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/dedenurr/management-akademik/database"
	"github.com/dedenurr/management-akademik/routers"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
	err error
)



func main()  {
	// load env
	godotenv.Load("config/.env")
	// validation load env
	if err != nil {
		fmt.Println("Failed Load File Environtment")
	}else{
		fmt.Println("Success Read File Environtment")
	}

	// insert variabel env to psqlInfo
	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"))

	// connect to postgress
	DB, err = sql.Open("postgres",psqlInfo)
	// test connection
	err = DB.Ping()
	
	if err != nil {
		fmt.Println("Connection to database is failed")
		panic(err)
	}else{
		fmt.Println("Successfully make connection to database")
	}

	// migrasi table database
	database.DbMigrate(DB)

	//close db
	defer DB.Close()

	// start server 
	PORT := ":" + os.Getenv("PORT")
	routers.SetupRouter(DB).Run(PORT) 
} 
