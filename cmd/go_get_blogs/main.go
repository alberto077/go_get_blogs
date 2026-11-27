package main

import (
	"fmt"
	"os"

	"github.com/alberto077/go_get_blogs/internal/commands"
	"github.com/alberto077/go_get_blogs/internal/config"
)


func main() {
/*
	loginName := flag.String("login", "", "User that wants to login")

	flag.Parse()

	if *loginName == "" {
		fmt.Println("Login name is required")
		os.Exit(1) 
	}

*/
	// Would probably have a while loop here to keep reading commands
	appState := config.NewState()

	cmdRegistry := commands.CommandRegistry{
		Commands: make(map[string]func(s *config.State, cmd commands.Command) error),
	}

	cmdRegistry.Register("login", commands.HandleLogin)

	if len(os.Args) < 2 {
		return 
	}

	cmd := commands.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := cmdRegistry.Run(appState, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	

	fmt.Printf("Config db url: %s, CurrentUser %s\n", appState.CurrConfig.DbURL, appState.CurrConfig.CurrentUserName)




}
