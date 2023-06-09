package utils

import (
	"github.com/pkg/errors"
	"runtime"
	"strings"
)

func Wrap(err error, message string) error {
	return errors.Wrap(err, "==> "+message)
}

func GetSelfFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return cleanUpFuncName(runtime.FuncForPC(pc).Name())
}
func cleanUpFuncName(funcName string) string {
	end := strings.LastIndex(funcName, ".")
	if end == -1 {
		return ""
	}
	return funcName[end+1:]
}
