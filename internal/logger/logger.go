package logger

import (
	"fmt"
	"time"
)

// Logger интерфейс определяет методы логгера
type Logger interface {
	Debug(message string)
	Error(err error)
}

// simpleLogger реализует интерфейс Logger
type simpleLogger struct{}

// NewLogger создает новый экземпляр simpleLogger
func NewLogger() Logger {
	return &simpleLogger{}
}

// Debug логирует отладочные сообщения
func (l *simpleLogger) Debug(message string) {
	// if message != "" {
	fmt.Printf("[%s] DEBUG: %s\n", time.Now().Format(time.RFC3339), message)
	// }
}

// Error логирует ошибки
func (l *simpleLogger) Error(err error) {
	fmt.Printf("[%s] ERROR: %v\n", time.Now().Format(time.RFC3339), err)
}
