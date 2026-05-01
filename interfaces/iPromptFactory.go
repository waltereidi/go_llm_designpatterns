package interfaces

type IPromptFactory interface {
	CreateStrategy() IPromptStrategy
}
