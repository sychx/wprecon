package printer

import (
	"bufio"
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

func ScanQ(text ...string) string {
	var prefix = color.Yellow("[?]").String()
	var textString = strings.Join(text, " ")

	_, err := io.WriteString(&stdout, prefix+" "+textString)

	if err != nil {
		Fatal(err)
	}

	scanner := bufio.NewReader(os.Stdin)

	response, err := scanner.ReadString('\n')

	if err != nil {
		Fatal(err)
	}

	response = strings.ToLower(response)

	if response == "\n" {
		return response
	}

	response = strings.ReplaceAll(response, "\n", "")

	return response
}

type l struct {
	text []string
}

// List ::
func List(text ...string) *l {
	return &l{text: text}
}

func (options *l) D() {
	var prefix = color.White("    —").String()
	var textString = strings.Join(options.text, " ")

	_, err := io.WriteString(&stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}

func (options *l) Done() {
	var prefix = color.Green("    —").String()
	var textString = strings.Join(options.text, " ")

	_, err := io.WriteString(&stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}

func (options *l) Danger() {
	var prefix = color.Red("    —").String()
	var textString = strings.Join(options.text, " ")

	_, err := io.WriteString(&stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}

func (options *l) Warning() {
	var prefix = color.Yellow("    —").String()
	var textString = strings.Join(options.text, " ")

	_, err := io.WriteString(&stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}

func (options *l) Info() {
	var prefix = color.Magenta("    —").String()
	var textString = strings.Join(options.text, " ")

	_, err := io.WriteString(&stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}

// TopLine :: "What is this struct for ?!" It will serve you some functions to write on the top line ... deleting the print of the NewTopLine function.
type TopLine struct {
	count  int
	Count2 int
}

// NewTopLine ::
func NewTopLine(text ...string) *TopLine {
	var topline TopLine

	var prefix = color.Yellow("[✲]").String()
	var textString = strings.Join(text, " ")

	_, err := io.WriteString(&stdout, prefix+" "+textString)

	if err != nil {
		panic(err)
	}

	return &topline
}

// DownLine :: An example of using this can be seen in the backup fuzzer file.
func (topline *TopLine) DownLine() {
	if topline.count <= 0 {
		_, err := io.WriteString(&stdout, "\n")

		if err != nil {
			panic(err)
		}
	}

	topline.count++
}

func (topline *TopLine) Done(text ...string) {
	var prefix = color.Green("[✔]").String()
	var textString = strings.Join(text, " ")

	fmt.Print("\033[G\033[K")
	_, err := io.WriteString(&stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}

func (topline *TopLine) Danger(text ...string) {
	var prefix = color.Red("[✗]").String()
	var textString = strings.Join(text, " ")

	fmt.Print("\033[G\033[K")
	_, err := io.WriteString(&stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}

func (topline *TopLine) Warning(text ...string) {
	var prefix = color.Yellow("[!]").String()
	var textString = strings.Join(text, " ")

	fmt.Print("\033[G\033[K")
	_, err := io.WriteString(&stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}

func (topline *TopLine) Progress(len int, text ...string) {

	var prefix = color.Yellow(fmt.Sprintf("[%d/%d]", topline.Count2, len)).String()
	var textString = strings.Join(text, " ")

	fmt.Print("\033[G\033[K")
	_, err := io.WriteString(&stdout, prefix+" "+textString)

	if err != nil {
		panic(err)
	}

	topline.Count2++
}
