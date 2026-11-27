package commands

import (
	"fmt"
	"github.com/alberto077/go_get_blogs/internal/config"
)


func HandleLogin(s *config.State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("login handler expects a username")
	}

	if err := s.CurrConfig.SetUser(cmd.Args[0]); err != nil {
		return err
	}

	fmt.Printf("User %s has been set\n", cmd.Args[0])

	return nil
}
