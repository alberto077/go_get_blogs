package main

import (
	"fmt"
	"github.com/alberto077/go_get_blogs/internal/config"
)


func main() {
	BaseConfig, err := config.Read()
	if err != nil {
		fmt.Println("Could not grab a new config")
	}

	BaseConfig.SetUser("Alberto")

	newConfig, err := config.Read()
	if err != nil {
		fmt.Println("Could read new config")
	}

	fmt.Printf("Config db url: %s, CurrentUser %s\n", newConfig.DbURL, newConfig.CurrentUserName)


}
