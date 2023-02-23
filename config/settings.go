package settings

import (
    "fmt"
    "os"

    "gopkg.in/yaml.v2"
)

const (
    DefaultConfigFilePath = "./conf/config.yaml"
)

func LoadAppConfig(appConfigPath string) (*AppConfig, error) {
    config := &AppConfig{}
    file, err := os.Open(appConfigPath)
    if err != nil {
        return nil, fmt.Errorf("open app config file failed, err: %v", err)
    }
    if err := yaml.NewDecoder(file).Decode(&config); err != nil {
        return nil, err
    }
    return config, nil
}
