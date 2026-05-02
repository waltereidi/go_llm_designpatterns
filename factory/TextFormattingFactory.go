package factory

import "go_llm_designpatterns/interfaces"

type TextFormattingFactory struct {
	CreateStrategy func() interfaces.IPromptStrategy
}

func (tf *TextFormattingFactory) CreateStrategy() interfaces.IPromptStrategy {
	return &TextFormattingStrategy{}
}
