package middleware

import (
	"log"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type login struct {
  Username string `form:"username" json:"username" binding:"required"`
  Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

func HelloHandler(c *gin.Context) {
  claims := jwt.ExtractClaims(c)
  user, _ := c.Get(identityKey)
  c.JSON(200, gin.H{
    "userID":   claims[identityKey],
    "userName": user.(*User).UserName,
    "text":     "Hello World.",
  })
}

// User demo
type User struct {
  UserName  string
  FirstName string
  LastName  string
}

func JWTMiddleware() *jwt.GinJWTMiddleware {

  // the jwt middleware
  authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
    Realm:       "test zone",
    Key:         []byte("sdf12341534fdsf"),
    Timeout:     time.Hour,
    MaxRefresh:  time.Hour,
    IdentityKey: identityKey,
    PayloadFunc: func(data interface{}) jwt.MapClaims {
      if v, ok := data.(*User); ok {
        return jwt.MapClaims{
          identityKey: v.UserName,
        }
      }
      return jwt.MapClaims{}
    },
    IdentityHandler: func(c *gin.Context) interface{} {
      claims := jwt.ExtractClaims(c)
      return &User{
        UserName: claims[identityKey].(string),
      }
    },
    Authenticator: func(c *gin.Context) (interface{}, error) {
      var loginVals login
      if err := c.ShouldBind(&loginVals); err != nil {
        return "", jwt.ErrMissingLoginValues
      }
      userID := loginVals.Username
      password := loginVals.Password

      if (userID == os.Getenv("USER_ONE") && password == os.Getenv("PWD_ONE")) || (userID ==os.Getenv("USER_TWO") && password == os.Getenv("PWD_TWO")) {
        return &User{
          UserName:  userID,
          LastName:  "Bo-Yi",
          FirstName: "Wu",
        }, nil
      }

      return nil, jwt.ErrFailedAuthentication
    },
    Authorizator: func(data interface{}, c *gin.Context) bool {
      if v, ok := data.(*User); ok && v.UserName == os.Getenv("USER_ONE") {
        return true
      }

      return false
    },
    Unauthorized: func(c *gin.Context, code int, message string) {
      c.JSON(code, gin.H{
        "code":    code,
        "message": message,
      })
    },
 
    TokenLookup: "header: Authorization, query: token, cookie: jwt",


    TokenHeadName: "Bearer",

    TimeFunc: time.Now,
  })
  
  if err != nil {
    log.Fatal("JWT Error:" + err.Error())
  }

  return authMiddleware
}