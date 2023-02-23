package routers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "golang-template/client/user"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    v1 := r.Group("/api/v1")
    v1.POST("/login", user.LoginHandler)
    v1.POST("/signup", user.SignUpHandler)
    v1.GET("/refresh_token", user.RefreshTokenHandler)

    v1.Use(http.JWTAuthMiddleware())
    {

        v1.GET("/ping", func(c *gin.Context) {
            c.String(http.StatusOK, "pong")
        })

    }

    r.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "msg": "404",
        })
    })
    return r
}
