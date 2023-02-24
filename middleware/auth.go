package middleware

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Auth(c *gin.Context) {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal(err)
	}
	// Get the Basic Authentication credentials
	user, password, hasAuth := c.Request.BasicAuth()
	if hasAuth && user == os.Getenv("USER_ONE")  && password == os.Getenv("PWD_ONE") || hasAuth && user == os.Getenv("USER_TWO") && password == os.Getenv("PWD_TWO") {
		c.Next()
		return
	} else if hasAuth && user == "" && password == "" {
		c.Abort()
		c.Writer.Write([]byte("Username atau password tidak boleh kosong"))
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		return
	} else {
		c.Abort()
		c.Writer.Write([]byte("Username atau password salah"))
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		return
	}

}