package utils

import (
	"fmt"
)

func printGoRoutineError(routineName string, r interface{}) {
	fmt.Println(fmt.Sprintf("Go routine '%s' error: %s", routineName, r))
}

func RecoverGoRoutineWithHandler(routineName string, function func()) {
	if r := recover(); r != nil {
		printGoRoutineError(routineName, r)
		function()
	}
}

func RecoverGoRoutine(routineName string) {
	if r := recover(); r != nil {
		printGoRoutineError(routineName, r)
	}
}
