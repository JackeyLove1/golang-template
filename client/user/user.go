package user

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "golang-template/config"
    "golang-template/dao/mysql"
    "golang-template/dao/redis"
    "golang-template/models"
    "golang-template/pkg/jwt"
    "golang-template/utils"
)

type Client struct {
    globalConfig *config.GlobalConfig
    db           *mysql.DB
    redis        *redis.Client
}

func SignUpHandler(c *gin.Context) {
    // 1.获取请求参数 2.校验数据有效性
    var fo models.RegisterForm
    if err := c.ShouldBindJSON(&fo); err != nil {
        utils.ResponseErrorWithMsg(c, utils.CodeInvalidParams, err.Error())
        return
    }
    // 3.注册用户
    /*
       err := mysql.Register(&models.User{
           UserName: fo.UserName,
           Password: fo.Password,
       })

       if errors.Is(err, mysql.ErrorUserExit) {
           utils.ResponseError(c, utils.CodeUserExist)
           return
       }
       if err != nil {
           zap.L().Error("mysql.Register() failed", zap.Error(err))
           utils.ResponseError(c, utils.CodeServerBusy)
           return
       }
    */
    utils.ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
    var u models.User
    if err := c.ShouldBindJSON(&u); err != nil {
        zap.L().Error("invalid params", zap.Error(err))
        utils.ResponseErrorWithMsg(c, utils.CodeInvalidParams, err.Error())
        return
    }
    /*
       if err := mysql.Login(&u); err != nil {
           zap.L().Error("mysql.Login(&u) failed", zap.Error(err))
           utils.ResponseError(c, utils.CodeInvalidPassword)
           return
       }
    */
    // 生成Token
    aToken, rToken, _ := jwt.GenToken(u.UserID)
    utils.ResponseSuccess(c, gin.H{
        "accessToken":  aToken,
        "refreshToken": rToken,
        "userID":       u.UserID,
        "username":     u.UserName,
    })
}

func RefreshTokenHandler(c *gin.Context) {
    rt := c.Query("refresh_token")
    // 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
    // 这里假设Token放在Header的Authorization中，并使用Bearer开头
    // 这里的具体实现方式要依据你的实际业务情况决定
    authHeader := c.Request.Header.Get("Authorization")
    if authHeader == "" {
        utils.ResponseErrorWithMsg(c, utils.CodeInvalidToken, "请求头缺少Auth Token")
        c.Abort()
        return
    }
    // 按空格分割
    parts := strings.SplitN(authHeader, " ", 2)
    if !(len(parts) == 2 && parts[0] == "Bearer") {
        utils.ResponseErrorWithMsg(c, utils.CodeInvalidToken, "Token格式不对")
        c.Abort()
        return
    }
    aToken, rToken, err := jwt.RefreshToken(parts[1], rt)
    fmt.Println(err)
    c.JSON(http.StatusOK, gin.H{
        "access_token":  aToken,
        "refresh_token": rToken,
    })
}
