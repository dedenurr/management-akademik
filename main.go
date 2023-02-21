package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/dedenurr/management-akademik/controllers"
	"github.com/dedenurr/management-akademik/database"
	"github.com/dedenurr/management-akademik/repository"
	"github.com/dedenurr/management-akademik/service"
	"github.com/gin-gonic/gin"

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

	//router gin
	router := gin.Default()

	//router dosens
	dosenRepository := repository.NewDosenRepository(DB)
	dosenService := service.NewDosenService(dosenRepository)
	dosenController := controllers.NewDosenController(dosenService)
	dosens := router.Group("/dosens")
	dosens.GET("", dosenController.ReadDosen)
	dosens.POST("", dosenController.CreateDosen)
	dosens.PUT("/:nip",dosenController.UpdateDosen)
	dosens.DELETE("/:nip",dosenController.DeleteDosen)

	mahasiswaRepository := repository.NewMahasiswaRepository(DB)
	mahasiswaService := service.NewMahasiswaService(mahasiswaRepository)
	mahasiswaController := controllers.NewMahasiswaController(mahasiswaService)
	mahasiswas := router.Group("/mahasiswas")
	mahasiswas.GET("", mahasiswaController.ReadMahasiswa)
	mahasiswas.POST("", mahasiswaController.CreateMahasiswa)
	mahasiswas.PUT("/:nip",mahasiswaController.UpdateMahasiswa)
	mahasiswas.DELETE("/:nip",mahasiswaController.DeleteMahasiswa)

	PORT := ":8080"
	router.Run(PORT)

	/* PORT := ":8080"
	routers.StartServer().Run(PORT) */

}