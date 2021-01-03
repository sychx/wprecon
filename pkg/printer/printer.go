package printer

import (
	"fmt"
	"io"
	"os"
	"strings"

	color "github.com/logrusorgru/aurora" // This is color lib
)

// Required :: Constant with the word "required" in red.
var Required string = color.Red("(Required)").Bold().String()

// Println ::
func Println(text ...interface{}) {
	fmt.Fprintln(os.Stdout, text...)
}

// Danger ::
func Danger(text ...string) {
	var prefix = color.Red("[✗]").String()
	var textString = strings.Join(text, " ")

	_, err := io.WriteString(os.Stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}

// Done ::
func Done(text ...string) {
	var prefix = color.Green("[✔]").String()
	var textString = strings.Join(text, " ")

	_, err := io.WriteString(os.Stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}

// Warning ::
func Warning(text ...string) {
	var prefix = color.Yellow("[!]").String()
	var textString = strings.Join(text, " ")

	_, err := io.WriteString(os.Stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}

// Wait ::
func Wait(text ...string) {
	var prefix = color.Green("[—]").String()
	var textString = strings.Join(text, " ")

	_, err := io.WriteString(os.Stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}

// Fatal ::
func Fatal(text ...interface{}) {
	var prefix = color.Red("[!]").String()

	fmt.Fprint(os.Stdout, prefix, " ")
	fmt.Fprintln(os.Stdout, text...)

	os.Exit(0)
}
