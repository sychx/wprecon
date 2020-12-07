package printer

import (
	"fmt"
	"os"

	color "github.com/logrusorgru/aurora" // This is color lib
)

func Println(text ...interface{}) {
	fmt.Fprintln(os.Stdout, text...)
}

// Danger ::
func Danger(text ...interface{}) {
	prefix := color.Red("[!]").String()

	fmt.Fprint(os.Stdout, prefix, " ")
	fmt.Fprintln(os.Stdout, text...)
}

// Done ::
func Done(text ...interface{}) {
	prefix := color.Green("[•]").String()

	fmt.Fprint(os.Stdout, prefix, " ")
	fmt.Fprintln(os.Stdout, text...)
}

// Warning ::
func Warning(text ...interface{}) {
	prefix := color.Yellow("[•••]").String()

	fmt.Fprint(os.Stdout, prefix, " ")
	fmt.Fprintln(os.Stdout, text...)
}

// Loading ::
func Loading(text ...interface{}) {
	prefix := color.Green("[*]").String()

	fmt.Fprint(os.Stdout, prefix, " ")
	fmt.Fprintln(os.Stdout, text...)
}

// Wait ::
func Wait(text ...interface{}) {
	prefix := color.Green("[—]").String()

	fmt.Fprint(os.Stdout, prefix, " ")
	fmt.Fprintln(os.Stdout, text...)
}
