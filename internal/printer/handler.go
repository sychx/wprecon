package printer

// Panic is a function that receives the error and handles it according to the condition by printing a log equivalent
// to fmt.Println() followed by a call to panic().
func HandlingPanic(err error, message ...string) {
	if err != nil && message == nil {
		Println(err)
	} else if err != nil && message != nil {
		Println(message)
	}

	panic(err)	
}

// Fatal is a function that receives the error and handles it according to the condition by printing a log equivalent
func HandlingFatal(err error, message ...string) {
	if err != nil && message == nil {
		Fatal(err)
	} else if err != nil && message != nil {
		Fatal(message)
	}
}
