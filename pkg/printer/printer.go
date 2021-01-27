package printer

import (
	"fmt"
	"io"
	"os"
	"strings"
	"syscall"

	color "github.com/logrusorgru/aurora"
)

// Required :: Constant with the word "required" in red.
var Required = color.Red("(Required)").Bold().String()
var stdout = *os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")
var stderr = *os.NewFile(uintptr(syscall.Stdout), "/dev/stderr")

// Println ::
func Println(text ...interface{}) {
	fmt.Fprintln(&stdout, text...)
}

// Done ::
func Done(text ...string) {
	var prefix = color.Green("[✔]").String()
	var textString = strings.Join(text, " ")

	_, err := io.WriteString(&stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}

// Danger ::
func Danger(text ...string) {
	var prefix = color.Red("[✗]").String()
	var textString = strings.Join(text, " ")

	_, err := io.WriteString(&stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}

// Warning ::
func Warning(text ...string) {
	var prefix = color.Yellow("[!]").String()
	var textString = strings.Join(text, " ")

	_, err := io.WriteString(&stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}

// List ::
func List(text ...string) {
	var prefix = color.White("    —").String()
	var textString = strings.Join(text, " ")

	_, err := io.WriteString(&stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}

// Info ::
func Info(text ...string) {
	var prefix = color.Magenta("[i]").String()
	var textString = strings.Join(text, " ")

	_, err := io.WriteString(&stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}

// Fatal ::
func Fatal(text ...interface{}) {
	var prefix = color.Red("[!]").String()

	fmt.Fprint(&stderr, prefix, " ")
	fmt.Fprintln(&stderr, text...)

	os.Exit(0)
}
