package config

import (
	"encoding/json"
	"fmt"
	"os"
)
const configFileName string = ".blogconfig.json"

type Config struct {
	DbURL string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func GetHomeFilePath() (string, error) {
	var filePath string
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return filePath, fmt.Errorf("Could not read file")
	}

	filePath = fmt.Sprintf("%s/%s", homeDir, configFileName)
	return filePath, err
}

func (c *Config) SetUser(name string) error{
	filePath, err := GetHomeFilePath()

	if err != nil {
		return err
	}

	c.CurrentUserName = name

	jsonBytes, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("Could now convert current config into bytes")
	}

	if err := os.WriteFile(filePath, jsonBytes, 0664); err != nil{
		return fmt.Errorf("Could not write into %s", filePath)
	}

	return nil
}

func Read() (*Config, error){

	var config Config

	filePath, err := GetHomeFilePath()
	if err != nil {
		return &config, err
	}
	
	fd, err := os.Open(filePath)
	if err != nil {
		return &config, fmt.Errorf("Could not open config file")
	}

	defer fd.Close()

	if err := json.NewDecoder(fd).Decode(&config); err != nil {
		return &config, fmt.Errorf("Could not decode json data")
		
	}

	return &config, nil
}


