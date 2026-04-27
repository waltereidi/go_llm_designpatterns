package interfaces

type IPromptFactory interface {
	createStrategy() IPromptStrategy
}
