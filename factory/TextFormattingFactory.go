package factory

import "go_llm_designpatterns/interfaces"

type TextFormattingFactory(i string ) struct {
}

func (tf *TextFormattingFactory) CreateStrategy() interfaces.IPromptStrategy {
	return &TextFormattingStrategy{}
}
