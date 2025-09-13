package model

type Notifier interface {
	PostMessage(message string) error
}
