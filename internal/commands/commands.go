package commands

import (
	"github.com/alberto077/go_get_blogs/internal/config"
)

type CommandRegistry struct {
	Commands map[string]func(s *config.State, cmd Command) error
}

type Command struct {
	Name string
	Args []string

}

func (c *CommandRegistry) Run (s *config.State, cmd Command) error{
	if err := c.Commands[cmd.Name](s, cmd); err != nil {
		return err
	}
	return nil
}

func (c *CommandRegistry) Register (name string, f func(s *config.State, cmd Command) error){
	c.Commands[name] = f
}

