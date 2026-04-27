package factory

import "interfaces"

type TextFormattingFactory struct {
	createStrategy interfaces.IPromptStrategy
}

func createStrategy() *interfaces.IPromptStrategy {
	return &TextFormattingStrategy{


	}
}
