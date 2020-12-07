package printer

import (
	"fmt"
	"os"

	color "github.com/logrusorgru/aurora" // This is color lib
)

/*
Println :: As I was wanting to make the println more cute with error indicators and etc ...
I preferred to do this function/this package.

But you must be in doubt as to whether there is a 'Verbose' parameter, the question is simple: As I wanted to do something as simple and with fewer functions as possible, I preferred to pass this 'Verbose' here, because it would be very simple there ... when making the verbose mode, and it would be more organized ... the data and etc ...
If the print you are going to make will not use verbose mode, you can enter the value 'false' which will have no problem.
*/

func Println(text ...interface{}) {
	fmt.Fprintln(os.Stdout, text...)
}

// Danger ::
func Danger(text ...interface{}) {
	prefix := color.Red("[!]").String()
	fmt.Fprintln(os.Stdout, prefix, text)

}

// Done ::
func Done(text ...interface{}) {
	prefix := color.Green("[•]").String()
	fmt.Fprintln(os.Stdout, prefix, text)

}

// Warning ::
func Warning(text ...interface{}) {
	prefix := color.Yellow("[•••]").String()
	fmt.Fprintln(os.Stdout, prefix, text)

}

// Loading ::
func Loading(text ...interface{}) {
	prefix := color.Green("[*]").String()
	fmt.Fprintln(os.Stdout, prefix, text)

}

// Wait ::
func Wait(text ...interface{}) {
	prefix := color.Green("[—]").String()
	fmt.Fprintln(os.Stdout, prefix, text)
}
