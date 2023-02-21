package routers

/* import (
	"database/sql"

	"github.com/dedenurr/management-akademik/controllers"
	"github.com/dedenurr/management-akademik/repository"
	"github.com/dedenurr/management-akademik/service"
	"github.com/gin-gonic/gin"
)

var (
	db  *sql.DB
)

func StartServer() *gin.Engine {
	router := gin.Default()

	//router dosens
	dosenRepository := repository.NewDosenRepository(db)
	dosenService := service.NewDosenService(dosenRepository)
	dosenHandler := controllers.InitDosenHandler(dosenService)
	dosens := router.Group("/dosens")
	dosens.GET("", dosenHandler.GetDosen)
	dosens.POST("", dosenHandler.InsertDosen)
	return router
}  */