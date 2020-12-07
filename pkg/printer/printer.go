package printer

import (
	"fmt"
	"os"

	color "github.com/logrusorgru/aurora" // This is color lib
)

// Println ::
func Println(text ...interface{}) {
	fmt.Fprintln(os.Stdout, text...)
}

// Danger ::
func Danger(text ...interface{}) {
	var prefix = color.Red("[!]").String()

	fmt.Fprint(os.Stdout, prefix, " ")
	fmt.Fprintln(os.Stdout, text...)
}

// Done ::
func Done(text ...interface{}) {
	var prefix = color.Green("[•]").String()

	fmt.Fprint(os.Stdout, prefix, " ")
	fmt.Fprintln(os.Stdout, text...)
}

// Warning ::
func Warning(text ...interface{}) {
	var prefix = color.Yellow("[•••]").String()

	fmt.Fprint(os.Stdout, prefix, " ")
	fmt.Fprintln(os.Stdout, text...)
}

// Loading ::
func Loading(text ...interface{}) {
	var prefix = color.Yellow("[*]").String()

	fmt.Fprint(os.Stdout, prefix, " ")
	fmt.Fprint(os.Stdout, text...)
}

// Loading ::
func LoadingDone(text ...interface{}) {
	var prefix = color.Green("[+]").String()

	fmt.Print("\033[G\033[K")
	fmt.Fprint(os.Stdout, prefix, " ")
	fmt.Fprintln(os.Stdout, text...)
}

// Loading ::
func LoadingDanger(text ...interface{}) {
	var prefix = color.Red("[!]").String()

	fmt.Print("\033[G\033[K")
	fmt.Fprint(os.Stdout, prefix, " ")
	fmt.Fprintln(os.Stdout, text...)
}

// Wait ::
func Wait(text ...interface{}) {
	var prefix = color.Green("[—]").String()

	fmt.Fprint(os.Stdout, prefix, " ")
	fmt.Fprintln(os.Stdout, text...)
}

func Fatal(text ...interface{}) {
	var prefix = color.Red("[!]").String()

	fmt.Fprint(os.Stdout, prefix, " ")
	fmt.Fprintln(os.Stdout, text...)

	os.Exit(0)
}

func Required(text ...interface{}) string {
	var sufix = color.Red("(Required)").Bold().String()

	return sufix
}
