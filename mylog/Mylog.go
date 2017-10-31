package mylog

import (
	"fmt"
	"os"
	"time"
)

var logDivPath = "./Json/Log/"
var logFilePath = time.Now().Format("2006-01-02") + ".txt"

func getFileHandle() *os.File {
	if _, err := os.Open(logDivPath + logFilePath); err != nil {
		os.Create(logDivPath + logFilePath)
	}

	// 以追加模式打开文件,并向文件写入
	fi, _ := os.OpenFile(logDivPath+logFilePath, os.O_RDWR|os.O_APPEND, 0)
	return fi
}

// AddLog : 添加记录
func AddLog(user string, command string, oldStr string, newStr string) {
	file := getFileHandle()
	if user != "" {
		fmt.Fprintf(file, "User:%s  ", user)
	}
	if command != "" {
		fmt.Fprintf(file, "Command:%s\n", command)
	}
	if oldStr != "" {
		fmt.Fprintf(file, "From:%s\n", oldStr)
	}
	if newStr != "" {
		fmt.Fprintf(file, "To:%s\n", newStr)
	}
	fmt.Fprintf(file, "\n")
	file.Close()
}
