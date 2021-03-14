package printer

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	color "github.com/logrusorgru/aurora"
)

var (
	stdin    = *os.Stdin
	stdout   = *os.Stdout
	stderr   = *os.Stderr
	line     = &ln{}
	zfill    = &z{}
	Required = color.Red("(Required)").Bold().String()
)

var (
	prefixDanger  = color.Red("[✗]").String()
	prefixFatal   = color.Red("[!]").String()
	prefixDone    = color.Green("[✔]").String()
	prefixWarning = color.Yellow("[!]").String()
	prefixScan    = color.Yellow("[?]").String()
	prefixInfo    = color.Magenta("[i]").String()

	prefixListDanger  = color.Red("    —").String()
	prefixListDone    = color.Green("    —").String()
	prefixListDefault = color.White("    —").String()
	prefixListWarning = color.Yellow("    —").String()

	prefixTopLine = color.Yellow("[✲]").String()
)

var seekCurrent = 1

func init() {
	if runtime.GOOS == "windows" {
		prefixDanger = "[✗]"
		prefixFatal = "[!]"
		prefixDone = "[✔]"
		prefixWarning = "[!]"
		prefixScan = "[?]"
		prefixInfo = "[i]"

		prefixListDanger = "    —"
		prefixListDone = "    —"
		prefixListDefault = "    —"
		prefixListWarning = "    —"

		prefixTopLine = "[✲]"
	}
}

type ln struct{}

func (l *ln) L() *ln {
	fmt.Fprintln(&stdout)

	return l
}

// Println ::
func Println(t ...interface{}) {
	fmt.Fprintln(&stdout, t...)
}

func Printf(format string, t ...interface{}) {
	fmt.Fprintf(&stdout, format, t...)
}

func Done(t ...string) *ln {
	var raw = strings.Join(t, " ")

	io.WriteString(&stdout, prefixDone+" "+raw+"\n")

	return line
}

func Bars(t string) *ln {
	var list = strings.Split(t, "\n")

	for num, txt := range list {
		if num+1 != len(list) {
			fmt.Fprintln(&stdout, " |  ", txt)
		}
	}

	return line
}

func Danger(t ...string) *ln {
	var raw = strings.Join(t, " ")

	io.WriteString(&stdout, prefixDanger+" "+raw+"\n")

	return line
}

func Warning(t ...string) *ln {
	var raw = strings.Join(t, " ")

	io.WriteString(&stdout, prefixWarning+" "+raw+"\n")

	return line
}

func Info(t ...string) *ln {
	var raw = strings.Join(t, " ")

	io.WriteString(&stdout, prefixInfo+" "+raw+"\n")

	return line
}

func Fatal(t ...interface{}) {
	fmt.Fprint(&stdout, prefixFatal, " ")
	fmt.Fprintln(&stdout, t...)

	os.Exit(0)
}

func ScanQ(t ...string) string {
	var raw = strings.Join(t, " ")

	io.WriteString(&stdout, prefixScan+" "+raw)

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
	text   []string
	prefix string
}

// List ::
func List(text ...string) *l {
	return &l{text: text}
}

func (options *l) Prefix(s ...string) *l {
	options.prefix = strings.Join(s, " ")

	return options
}

func (options *l) D() *ln {
	var raw = strings.Join(options.text, " ")

	io.WriteString(&stdout, options.prefix+prefixListDefault+" "+raw+"\n")

	return line
}

func (options *l) Done() *ln {
	var raw = strings.Join(options.text, " ")

	io.WriteString(&stdout, options.prefix+prefixListDone+" "+raw+"\n")

	return line
}

func (options *l) Danger() *ln {
	var raw = strings.Join(options.text, " ")

	io.WriteString(&stdout, options.prefix+prefixListDanger+" "+raw+"\n")

	return line
}

func (options *l) Warning() *ln {
	var raw = strings.Join(options.text, " ")

	io.WriteString(&stdout, options.prefix+prefixListWarning+" "+raw+"\n")

	return line
}

type topline struct{}

func NewTopLine(t ...string) *topline {
	var raw = strings.Join(t, " ")

	io.WriteString(&stdout, prefixTopLine+" "+raw)

	return &topline{}
}

func (options *topline) Done(t ...string) {
	var raw = strings.Join(t, " ")

	fmt.Fprint(&stdout, "\033[G\033[K")
	io.WriteString(&stdout, prefixDone+" "+raw+"\n")
}

func (options *topline) Danger(t ...string) {
	var raw = strings.Join(t, " ")

	fmt.Fprint(&stdout, "\033[G\033[K")
	io.WriteString(&stdout, prefixDanger+" "+raw+"\n")
}

func (options *topline) Warning(t ...string) {
	var raw = strings.Join(t, " ")

	fmt.Fprint(&stdout, "\033[G\033[K")
	io.WriteString(&stdout, prefixWarning+" "+raw+"\n")
}

func (options *topline) Info(t ...string) {
	var raw = strings.Join(t, " ")

	fmt.Fprint(&stdout, "\033[G\033[K")
	io.WriteString(&stdout, prefixInfo+" "+raw+"\n")
}

type z struct{}

func (options *z) Fill() {
	seekCurrent = 0
}

func (options *topline) Progress(seek int, t ...string) *z {
	var prefix = color.Yellow(fmt.Sprintf("[%d/%d]", seekCurrent, seek)).String()
	var raw = strings.Join(t, " ")

	seekCurrent++

	fmt.Fprint(&stdout, "\033[G\033[K")
	io.WriteString(&stdout, prefix+" "+raw)

	return zfill
}
