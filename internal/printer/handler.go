package printer

// Fatal is a function that receives the error and handles it according to the condition by printing a log equivalent
func HandlingFatal(err error) {
    if err != nil { Fatal(err) }
}