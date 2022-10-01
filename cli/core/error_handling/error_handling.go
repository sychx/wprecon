package error_handling

import (
	"fmt"
	"os"
)

// Panic is a function that receives the error and handles it according to the condition by printing a log equivalent
// to fmt.Println() followed by a call to panic().
func Panic(err error, message ...string) {
	if err != nil {
		if message == nil {
			fmt.Println(err)
		} else {
			fmt.Println(message)
		}
		panic(err)
	}
}

// Fatal is a function that receives the error and handles it according to the condition by printing a log equivalent
// to Println() followed by a call to os.Exit(1).
func Fatal(err error, message ...string) {
	if err != nil {
		if message == nil {
			fmt.Println(err)
		} else {
			fmt.Println(message)
		}
		os.Exit(1)
	}
}
