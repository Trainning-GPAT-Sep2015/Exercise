package loadConfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func FromFile(cfg interface{}, configPath string) error {
	absPath, err := filepath.Abs(configPath)
	if err != nil {
		return err
	}

	fmt.Println("Load config from file", absPath)
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, cfg)
	if err != nil {
		return err
	}

	return nil
}
