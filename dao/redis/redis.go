package redis

import (
    "fmt"

    "github.com/go-redis/redis"
    "golang-template/config"
)

var (
    Nil = redis.Nil
)

type SliceCmd = redis.SliceCmd
type StringStringMapCmd = redis.StringStringMapCmd

type Client struct {
    client *redis.Client
}

// Init 初始化连接
func Init(globalConfig *config.GlobalConfig) (*Client, error) {
    cfg := globalConfig.RedisConfig
    client := redis.NewClient(&redis.Options{
        Addr:         fmt.Sprintf("%s:%d", cfg.Host, globalConfig.Port),
        Password:     cfg.Password, // no password set
        DB:           cfg.DB,       // use default DB
        PoolSize:     cfg.PoolSize,
        MinIdleConns: cfg.MinIdleConns,
    })

    _, err := client.Ping().Result()
    if err != nil {
        return nil, err
    }
    return &Client{
        client: client,
    }, nil
}

func (c *Client) Close() {
    _ = c.client.Close()
}
