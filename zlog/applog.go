package zlog

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var ctx context.Context

func Init(appCtx context.Context) {
	ctx = appCtx
}

func Debug(message string) {
	if ctx == nil {
		fmt.Println(message)
		return
	}
	runtime.LogDebug(ctx, message)
}

func Debugf(format string, args ...interface{}) {
	if ctx == nil {
		fmt.Println(fmt.Sprintf(format, args...))
		return
	}
	runtime.LogDebugf(ctx, format, args...)
}

func Info(message string) {
	if ctx == nil {
		fmt.Println(message)
		return
	}
	runtime.LogInfo(ctx, message)
}

func Infof(format string, args ...interface{}) {
	if ctx == nil {
		fmt.Println(fmt.Sprintf(format, args...))
		return
	}
	runtime.LogInfof(ctx, format, args...)
}

func Error(error error) {
	if ctx == nil {
		fmt.Println(error.Error())
		return
	}
	runtime.LogError(ctx, error.Error())
}

func ErrorStr(message string) {
	if ctx == nil {
		fmt.Println(message)
		return
	}
	runtime.LogError(ctx, message)
}

func ErrorStrf(format string, args ...interface{}) {
	if ctx == nil {
		fmt.Println(fmt.Sprintf(format, args...))
		return
	}
	runtime.LogErrorf(ctx, format, args...)
}
