package settings

import "testing"

func TestLoadAppConfig(t *testing.T) {
    app, err := LoadAppConfig(DefaultConfigFilePath)
    if err != nil {
        t.Errorf("LoadAppConfig Error:%v", err)
    }
    t.Log(app.String())
}
