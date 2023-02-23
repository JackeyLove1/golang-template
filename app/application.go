package app

import (
    "golang-template/client/user"
    "golang-template/config"
    "golang-template/dao/mysql"
    "golang-template/dao/redis"
    "golang-template/logger"
)

type App struct {
    GlobalConfig *config.GlobalConfig
    DB           *mysql.DB
    Redis        *redis.Client
    UserHandler  *user.Client
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

    db, err := mysql.Init(appConfig)
    panicError(err)
    defer db.Close()

    redisClient, err := redis.Init(appConfig)
    panicError(err)
    defer redisClient.Close()
    return &App{
        GlobalConfig: appConfig,
        DB:           db,
        Redis:        redisClient,
    }
}

func (a *App) Register() {
    // TODO
}
