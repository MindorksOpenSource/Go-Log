package golog

import (
	"fmt"
	"log"
	"runtime"
	"sync"
)

var mutex = &sync.Mutex{}
var wg sync.WaitGroup

// D : Default Logger
func D(message string) {
	generateMessage(message)
}

// E : Error Logger
func E(message string) {
	generateMessage("[ERROR] " + message)
}

func generateMessage(message string) {
	var resultMessage string
	if IsCallingFunctionEnabled == true {
		resultMessage += "[" + getCallingFunctionName() + "]"
	}
	resultMessage += " " + message

	wg.Add(1)
	go func() {
		defer wg.Done()
		printMessage(resultMessage)
	} ()

	wg.Wait()
}

func printMessage(resultMessage string) {
	mutex.Lock()
	defer mutex.Unlock()
	if IsTimeEnabled {
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
