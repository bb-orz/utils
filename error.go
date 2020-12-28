package utils

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

/*
通用错误处理函数,用于需要记录位置的错误处理
Usage:
if !EH(err) {
	// return error or transfer error
}
*/
func EH(err error) bool {
	if err != nil {
		if _, file, line, ok := runtime.Caller(1); ok {
			path := file + "" + strconv.Itoa(line)
			fmt.Printf("[Error]:%s --- [Location]:%s", err.Error(), path)
		}
		return false
	}
	return true
}

/*
注明调用深度的错误处理函数
Usage:
if !EHCaller(err,2) {
	// return error or transfer error
}
*/
func EHCaller(err error, callerDepth int) bool {
	if err != nil {
		if _, file, line, ok := runtime.Caller(callerDepth); ok {
			path := file + "" + strconv.Itoa(line)
			fmt.Printf("[Error]:%s --- [Location]:%s --- [Caller Depth]:%d", err.Error(), path, callerDepth)
		}
		return false
	}
	return true
}

func Emsg(err error, msg string) bool {
	if err != nil {
		if _, file, line, ok := runtime.Caller(1); ok {
			path := file + "" + strconv.Itoa(line)
			fmt.Printf("[Error]:%s --- [Location]:%s --- [Message]:%s", err.Error(), path, msg)
		}
		return false
	}
	return true
}

// 通用错误退出函数
func EF(err error) {
	if err != nil {
		if _, file, line, ok := runtime.Caller(1); ok {
			path := file + "" + strconv.Itoa(line)
			fmt.Printf("[Fatal]:%s --- [Location]:%s", err.Error(), path)
		}
		os.Exit(-1)
	}
}
