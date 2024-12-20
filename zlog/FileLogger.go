package zlog

import (
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"os"
	"time"
)

// FileLogger is a utility to applog messages to a number of destinations
type FileLogger struct {
	filename string
}

// NewFileLogger creates a new Logger.
func NewFileLogger(filename string) logger.Logger {
	return &FileLogger{
		filename: filename,
	}
}

// Print works like Sprintf.
func (l *FileLogger) Print(message string) {
	// 控制台输出
	fmt.Print(message)

	// 文件输出
	f, err := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(f)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	if _, err := f.WriteString(message); err != nil {
		fmt.Printf("Error writing file: %v\n", err)
	}
}

func (l *FileLogger) Println(message string) {
	timeStr := time.Now().Format("2006-01-02 15:04:05 ")
	l.Print(timeStr + message + "\n")
}

// Trace level logging. Works like Sprintf.
func (l *FileLogger) Trace(message string) {
	l.Println("[TRACE] " + message)
}

// Debug level logging. Works like Sprintf.
func (l *FileLogger) Debug(message string) {
	l.Println("[DEBUG] " + message)
}

// Info level logging. Works like Sprintf.
func (l *FileLogger) Info(message string) {
	l.Println("[INFO]  " + message)
}

// Warning level logging. Works like Sprintf.
func (l *FileLogger) Warning(message string) {
	l.Println("[WARN]  " + message)
}

// Error level logging. Works like Sprintf.
func (l *FileLogger) Error(message string) {
	l.Println("[ERROR] " + message)
}

// Fatal level logging. Works like Sprintf.
func (l *FileLogger) Fatal(message string) {
	l.Println("[FATAL] " + message)
	os.Exit(1)
}
