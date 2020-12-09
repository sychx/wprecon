package printer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	color "github.com/logrusorgru/aurora" // This is color lib
	"golang.org/x/net/html/charset"
)

// Println ::
func Println(text ...interface{}) {
	fmt.Fprintln(os.Stdout, text...)
}

// Decode :: Convert to UTF-8
func Decode(text string) string {
	r, err := charset.NewReader(strings.NewReader(text), "latin1")
	if err != nil {
		log.Fatal(err)
	}
	result, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	return string(result)
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

// LoadingDone ::
func LoadingDone(text ...interface{}) {
	var prefix = color.Green("[+]").String()

	fmt.Print("\033[G\033[K")
	fmt.Fprint(os.Stdout, prefix, " ")
	fmt.Fprintln(os.Stdout, text...)
}

// LoadingDanger ::
func LoadingDanger(text ...interface{}) {
	var prefix = color.Red("[!]").String()

	fmt.Print("\033[G\033[K")
	fmt.Fprint(os.Stdout, prefix, " ")
	fmt.Fprintln(os.Stdout, text...)
}

// LoadingWarning ::
func LoadingWarning(text ...interface{}) {
	var prefix = color.Yellow("[•••]").String()

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

// Fatal ::
func Fatal(text ...interface{}) {
	var prefix = color.Red("[!]").String()

	fmt.Fprint(os.Stdout, prefix, " ")
	fmt.Fprintln(os.Stdout, text...)

	os.Exit(0)
}

// Required ::
func Required(text ...interface{}) string {
	var sufix = color.Red("(Required)").Bold().String()

	return sufix
}
