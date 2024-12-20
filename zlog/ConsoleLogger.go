package zlog

import (
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"os"
	"time"
)

// ConsoleLogger is a utility to applog messages to a number of destinations
type ConsoleLogger struct {
}

// NewConsoleLogger creates a new Logger.
func NewConsoleLogger() logger.Logger {
	return &ConsoleLogger{}
}

// Print works like Sprintf.
func (l *ConsoleLogger) Print(message string) {
	// 控制台输出
	fmt.Print(message)
}

func (l *ConsoleLogger) Println(message string) {
	timeStr := time.Now().Format("2006-01-02 15:04:05 ")
	l.Print(timeStr + message + "\n")
}

// Trace level logging. Works like Sprintf.
func (l *ConsoleLogger) Trace(message string) {
	l.Println("[TRACE] " + message)
}

// Debug level logging. Works like Sprintf.
func (l *ConsoleLogger) Debug(message string) {
	l.Println("[DEBUG] " + message)
}

// Info level logging. Works like Sprintf.
func (l *ConsoleLogger) Info(message string) {
	l.Println("[INFO]  " + message)
}

// Warning level logging. Works like Sprintf.
func (l *ConsoleLogger) Warning(message string) {
	l.Println("[WARN]  " + message)
}

// Error level logging. Works like Sprintf.
func (l *ConsoleLogger) Error(message string) {
	l.Println("[ERROR] " + message)
}

// Fatal level logging. Works like Sprintf.
func (l *ConsoleLogger) Fatal(message string) {
	l.Println("[FATAL] " + message)
	os.Exit(1)
}
