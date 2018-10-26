package gologger

import (
	"fmt"
	"golog/configure"
	"log"
	"runtime"
)

// Default Logger
func D(message string) {
	printMessage(message)
}

// Error Logger
func E(message string) {
	printMessage("[ERROR] " + message)
}

func printMessage(message string) {
	var resultMessage string
	if configure.IsCallingFunctionEnabled == true {
		resultMessage += "[" + getCallingFunctionName() + "]"
	}
	resultMessage += " " + message
	if configure.IsTimeEnabled {
		log.Println(resultMessage)
		return
	}
	fmt.Println(resultMessage)
}


func getCallingFunctionName() string {
	pc, _, _, ok := runtime.Caller(3)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		return details.Name()
	}
	return ""
}