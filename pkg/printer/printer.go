package printer

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	color "github.com/logrusorgru/aurora" // This is color lib
	"golang.org/x/net/html/charset"
)

// Required :: Constant with the word "required" in red.
var Required string = color.Red("(Required)").Bold().String()

// Println ::
func Println(text ...interface{}) {
	fmt.Fprintln(os.Stdout, text...)
}

// Decode :: This function will convert any string to UTF-8.
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

// Loading ::
func Loading(text ...string) {
	var prefix = color.Yellow("[✲]").String()
	var textString = strings.Join(text, " ")

	_, err := io.WriteString(os.Stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}

// LoadingDone ::
func LoadingDone(text ...string) {
	var prefix = color.Green("[✔]").String()
	var textString = strings.Join(text, " ")

	fmt.Print("\033[G\033[K")
	_, err := io.WriteString(os.Stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}

// LoadingDanger ::
func LoadingDanger(text ...string) {
	var prefix = color.Red("[✗]").String()
	var textString = strings.Join(text, " ")

	fmt.Print("\033[G\033[K")
	_, err := io.WriteString(os.Stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}

// LoadingWarning ::
func LoadingWarning(text ...string) {
	var prefix = color.Yellow("[!]").String()
	var textString = strings.Join(text, " ")

	fmt.Print("\033[G\033[K")
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
