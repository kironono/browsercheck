package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kironono/browsercheck/internal/model"
	"gopkg.in/yaml.v2"
)

type FileConfig struct {
	path string
}

func NewFileConfig(path string) *FileConfig {
	return &FileConfig{
		path: path,
	}
}

func (fc *FileConfig) Load() (*model.Config, error) {
	f, err := os.Open(fc.path)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s for reading", fc.path)
	}
	defer f.Close()

	body, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	loadedConfig := &model.Config{}
	if err = yaml.Unmarshal(body, &loadedConfig); err != nil {
		return nil, err
	}

	loadedConfig.SetDefaults()

	return loadedConfig, nil
}
