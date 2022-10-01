package error_handling

import (
	"github.com/AngraTeam/wprecon/internal/printer"
)

// Panic is a function that receives the error and handles it according to the condition by printing a log equivalent
// to fmt.Println() followed by a call to panic().
func Panic(err error, message ...string) {
	if err != nil {
		if message == nil {
			printer.Println(err)
		} else {
			printer.Println(message)
		}
		panic(err)
	}
}

// Fatal is a function that receives the error and handles it according to the condition by printing a log equivalent
// to Println() followed by a call to os.Exit(1).
func Fatal(err error, message ...string) {
	if err != nil {
		if message == nil {
			printer.Fatal(err)
		} else {
			printer.Fatal(message)
		}
	}
}
