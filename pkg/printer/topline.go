package printer

import (
	"fmt"
	"io"
	"os"
	"strings"

	color "github.com/logrusorgru/aurora"
)

// TopLine :: "What is this struct for ?!" It will serve you some functions to write on the top line ... deleting the print of the NewTopLine function.
type TopLine struct {
	Count int
}

// NewTopLine ::
func NewTopLine(text ...string) *TopLine {
	var topline TopLine

	var prefix = color.Yellow("[✲]").String()
	var textString = strings.Join(text, " ")

	_, err := io.WriteString(os.Stdout, prefix+" "+textString)

	if err != nil {
		panic(err)
	}

	return &topline
}

// DownLine :: An example of using this can be seen in the backup fuzzer file.
func (topline *TopLine) DownLine() {

	if topline.Count <= 0 {
		_, err := io.WriteString(os.Stdout, "\n")

		if err != nil {
			panic(err)
		}
	}

	topline.Count++
}

// Done ::
func (topline *TopLine) Done(text ...string) {
	var prefix = color.Green("[✔]").String()
	var textString = strings.Join(text, " ")

	fmt.Print("\033[G\033[K")
	_, err := io.WriteString(os.Stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}

// Danger ::
func (topline *TopLine) Danger(text ...string) {
	var prefix = color.Red("[✗]").String()
	var textString = strings.Join(text, " ")

	fmt.Print("\033[G\033[K")
	_, err := io.WriteString(os.Stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}

// Warning ::
func (topline *TopLine) Warning(text ...string) {
	var prefix = color.Yellow("[!]").String()
	var textString = strings.Join(text, " ")

	fmt.Print("\033[G\033[K")
	_, err := io.WriteString(os.Stdout, prefix+" "+textString+"\n")

	if err != nil {
		panic(err)
	}
}
