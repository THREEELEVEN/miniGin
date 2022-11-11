package day6

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"testing"
)

func TestCallers(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			message := fmt.Sprintf("%s", err)
			log.Printf("%s\n", Trace(message))
		}
	}()

	panic("XXX")
}

func Trace(message string) string {
	// 跳过自身调用栈的信息，只打印发生Panic的行信息
	skip := 3
	// 返回Panic的文件及具体行数
	_, file, line, _ := runtime.Caller(skip)
	var str strings.Builder
	str.WriteString(message + "\nTraceback: ")
	str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	return str.String()
}
