package main

import ("fmt" 
"go_llm_designpatterns/interfaces"
)

type CommandFactory struct {
	commands map[string]interfaces.CommandPattern
}

func NewCommandFactory() *CommandFactory {
	return &CommandFactory{
		commands: map[string]CommandPattern{
			"create_user": &CreateUserCommand{},
			"delete_user": &DeleteUserCommand{},
		},
	}
}

func (f *CommandFactory) GetCommand(msgType string) (CommandPattern, error) {
	cmd, exists := f.commands[msgType]
	if !exists {
		return nil, fmt.Errorf("comando não encontrado: %s", msgType)
	}
	return cmd, nil
}
