package routers

import (
	"database/sql"
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/dedenurr/management-akademik/controllers"
	"github.com/dedenurr/management-akademik/middleware"
	"github.com/dedenurr/management-akademik/repository"
	"github.com/dedenurr/management-akademik/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter(DB *sql.DB) *gin.Engine {
	//router gin
	router := gin.Default()

	//=== jwt
	errInit := middleware.JWTMiddleware().MiddlewareInit()

	if errInit != nil {
	  log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
  
	router.POST("/login", middleware.JWTMiddleware().LoginHandler)
  
	router.NoRoute(middleware.JWTMiddleware().MiddlewareFunc(), func(c *gin.Context) {
	  claims := jwt.ExtractClaims(c)
	  log.Printf("NoRoute claims: %#v\n", claims)
	  c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
  
	auth := router.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", middleware.JWTMiddleware().RefreshHandler)
	auth.Use(middleware.JWTMiddleware().MiddlewareFunc())
	{
	  auth.GET("/hello",middleware.HelloHandler)
	
	// === jwt
		//router dosens
		dosenRepository := repository.NewDosenRepository(DB)
		dosenService := service.NewDosenService(dosenRepository)
		dosenController := controllers.NewDosenController(dosenService)
		router.GET("/dosens", dosenController.ReadDosen)
		auth.POST("/dosens",  dosenController.CreateDosen)
		auth.PUT("/dosens/:nip",  dosenController.UpdateDosen)
		auth.DELETE("/dosens/:nip",  dosenController.DeleteDosen)

		// router mahasiswa
		mahasiswaRepository := repository.NewMahasiswaRepository(DB)
		mahasiswaService := service.NewMahasiswaService(mahasiswaRepository)
		mahasiswaController := controllers.NewMahasiswaController(mahasiswaService)
		router.GET("/mahasiswas",  mahasiswaController.ReadMahasiswa)
		auth.POST("/mahasiswas",  mahasiswaController.CreateMahasiswa)
		auth.PUT("/mahasiswas/:nim",  mahasiswaController.UpdateMahasiswa)
		auth.DELETE("/mahasiswas/:nim", mahasiswaController.DeleteMahasiswa)

		// router matakuliah
		mataKuliahRepository := repository.NewMataKuliahRepository(DB)
		mataKuliahService := service.NewMataKuliahService(mataKuliahRepository)
		mataKuliahController := controllers.NewMataKuliahController(mataKuliahService)
		router.GET("/matakuliahs",  mataKuliahController.ReadMataKuliah)
		auth.POST("/matakuliahs",  mataKuliahController.CreateMataKuliah)
		auth.PUT("/matakuliahs/:id",  mataKuliahController.UpdateMataKuliah)
		auth.DELETE("/matakuliahs/:id",  mataKuliahController.DeleteMataKuliah)

		// router perkuliahan
		perkuliahanRepository := repository.NewPerkuliahanRepository(DB)
		perkuliahanService := service.NewPerkuliahanService(perkuliahanRepository)
		perkuliahanController := controllers.NewPerkuliahanController(perkuliahanService)
		router.GET("/perkuliahans",  perkuliahanController.ReadPerkuliahan)
		auth.POST("/perkuliahans",  perkuliahanController.CreatePerkuliahan)
		auth.PUT("/perkuliahans/:id",  perkuliahanController.UpdatePerkuliahan)
		auth.DELETE("/perkuliahans/:id",  perkuliahanController.DeletePerkuliahan)
	} 
		
	
	return router
}
