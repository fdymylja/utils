package utils

import (
	"path/filepath"
	"runtime"
	"strings"
)

// CallerFunctionName returns the caller of the function calling CallerFunctionName (inception!)
func CallerFunctionName() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "UnknownFunction"
	}
	info := runtime.FuncForPC(pc)
	if info == nil {
		return "UnknownFunction"
	}
	fPath := info.Name()
	fileFunc, err := filepath.Abs(fPath)
	if err != nil {
		return "UnknownFunction"
	}
	fileFuncS := strings.Split(fileFunc, ".")
	funcName := fileFuncS[len(fileFuncS)-1]
	return funcName
}
