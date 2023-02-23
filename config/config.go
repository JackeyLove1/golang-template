package config

import "fmt"

type GlobalConfig struct {
    Mode         string `yaml:"mode"`
    Port         int    `yaml:"port"`
    *LogConfig   `yaml:"log"`
    *MySQLConfig `yaml:"mysql"`
    *RedisConfig `yaml:"redis"`
}

func (globalConfig *GlobalConfig) String() string {
    return fmt.Sprintf("App Config Mode:%q, Port:%d", globalConfig.Mode, globalConfig.Port)
}

type MySQLConfig struct {
    Host         string `yaml:"host"`
    User         string `yaml:"user"`
    Password     string `yaml:"password"`
    DB           string `yaml:"db"`
    Port         int    `yaml:"port"`
    MaxOpenConns int    `yaml:"max_open_conns"`
    MaxIdleConns int    `yaml:"max_idle_conns"`
}

type RedisConfig struct {
    Host         string `yaml:"host"`
    Password     string `yaml:"password"`
    Port         int    `yaml:"port"`
    DB           int    `yaml:"db"`
    PoolSize     int    `yaml:"pool_size"`
    MinIdleConns int    `yaml:"min_idle_conns"`
}

type LogConfig struct {
    Level      string `yaml:"level"`
    Filename   string `yaml:"filename"`
    MaxSize    int    `yaml:"max_size"`
    MaxAge     int    `yaml:"max_age"`
    MaxBackups int    `yaml:"max_backups"`
}
