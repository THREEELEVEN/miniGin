package day6

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

//	func trace(message string) string {
//		// linter 是一个无符号的整形，可以保存一个指针地址
//		var pcStack [32]uintptr
//		n := runtime.Callers(3, pcStack[:])
//
//		var str strings.Builder
//		str.WriteString(message + "\nTraceback: ")
//
//		frames := runtime.CallersFrames(pcStack[:n])
//		for {
//			frame, more := frames.Next()
//			str.WriteString(fmt.Sprintf("\n\t%s:%d", frame.File, frame.Line))
//			if !more {
//				break
//			}
//		}
//		return str.String()
//	}
func trace(message string) string {
	// 跳过自身调用栈的信息，只打印发生Panic的行信息
	skip := 4
	// Caller 获取调用函数信息
	// 返回发生Panic的文件及具体行数
	_, file, line, _ := runtime.Caller(skip)
	var str strings.Builder
	str.WriteString(message + "\nTraceback: ")
	str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	return str.String()
}

func recovery() HandlerFunc {
	return func(ctx *Context) {
		defer func() {
			err := recover()
			if err != nil {
				message := fmt.Sprintf("%s", err)
				//log.Printf("%s\n", trace(message))
				//log.Printf("\x1b[%dm%s\x1b[0m\n", 31, trace(message))
				log.Printf("\x1b[31m%s\x1b[0m\n", trace(message))
				ctx.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		ctx.Next()
	}
}
