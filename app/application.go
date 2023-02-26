package app

import (
    "github.com/sony/sonyflake"
    "golang-template/client/user"
    "golang-template/config"
    "golang-template/dao/mysql"
    "golang-template/dao/redis"
    "golang-template/dao/sqlite"
    "golang-template/logger"
)

type App struct {
    GlobalConfig *config.GlobalConfig
    MySQL        *mysql.DB
    Redis        *redis.Client
    UserHandler  *user.Client
    SQLite       *sqlite.DB
    IDGenerator  *sonyflake.Sonyflake
}

func panicError(err error) {
    if err != nil {
        panic(err)
    }
}

func MustNew() *App {
    appConfig, err := config.LoadAppConfig()
    panicError(err)

    _, err = logger.Init(appConfig, appConfig.Mode)
    panicError(err)

    mysqlClient, err := mysql.Init(appConfig)
    panicError(err)
    defer mysqlClient.Close()

    redisClient, err := redis.Init(appConfig)
    panicError(err)
    defer redisClient.Close()

    return &App{
        GlobalConfig: appConfig,
        MySQL:        mysqlClient,
        Redis:        redisClient,
    }
}

func (a *App) Register() {
    // TODO
}
