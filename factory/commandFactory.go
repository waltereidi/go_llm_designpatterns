package factory

import (
	"fmt"
	"go_llm_designpatterns/handlers"
	"go_llm_designpatterns/interfaces"
)

type CommandFactory struct {
	commands map[string]interfaces.CommandPattern
}

func NewCommandFactory() *CommandFactory {
	return &CommandFactory{
		commands: map[string]interfaces.CommandPattern{
			"syntax_validation": &handlers.SintaxValidationCommand{},
		},
	}
}

func (f *CommandFactory) GetCommand(msgType string) (interfaces.CommandPattern, error) {
	cmd, exists := f.commands[msgType]
	if !exists {
		return nil, fmt.Errorf("comando não encontrado: %s", msgType)
	}
	return cmd, nil
}
