package config

import (
    "fmt"
    "os"
    "path/filepath"

    "gopkg.in/yaml.v2"
)

const (
    DefaultConfigFilePath = "config.yaml"
)

func LoadAppConfig() (*GlobalConfig, error) {
    appConfigPath, err := os.Getwd()
    if err != nil {
        return nil, err
    }
    appConfigPath = filepath.Join(appConfigPath, DefaultConfigFilePath)
    config := &GlobalConfig{}
    file, err := os.Open(appConfigPath)
    if err != nil {
        return nil, fmt.Errorf("open app config file failed, err: %v", err)
    }
    if err := yaml.NewDecoder(file).Decode(&config); err != nil {
        return nil, err
    }
    return config, nil
}
