package printer

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync/atomic"
)

/*
	Why did I choose to use "io.WriteString" instead of the standard golang Println ?!
	io.WriteString is much simpler/faster than the standard Println.
*/

// The color code
const (
	Red       string = "\u001b[31;1m"
	Blue      string = "\u001b[34;1m"
	Green     string = "\u001b[32;1m"
	Black     string = "\u001b[30;1m"
	White     string = "\u001b[37;1m"
	Yellow    string = "\u001b[33;1m"
	Magenta   string = "\u001b[35;1m"
	Cyan      string = "\u001b[36;1m"
	Reset     string = "\u001b[0m"
	Bold      string = "\u001b[1m"
	Underline string = "\u001b[4m"
	Reversed  string = "\u001b[7m"
)

var (
	stdin    = *os.Stdin
	stdout   = *os.Stdout
	stderr   = *os.Stderr
	Required = Red + "(Required)" + Reset
)

var (
	prefixWarning = Yellow + "[!]" + Reset
	prefixDanger  = Red + "[✗]" + Reset
	prefixFatal   = Red + "[!]" + Reset
	prefixDone    = Green + "[✔]" + Reset
	prefixScan    = Yellow + "[?]" + Reset
	prefixInfo    = Magenta + "[i]" + Reset

	prefixListWarning = Yellow + "    —" + Reset
	prefixListDefault = White + "    —" + Reset
	prefixListDanger  = Red + "    —" + Reset
	prefixListDone    = Green + "    —" + Reset

	prefixTopLine = Yellow + "[✲]" + Reset
)

var SeekCurrent int64 = 0

func init() {
	switch runtime.GOOS {
	case "windows":
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

// Println :: Just a normal Println.
// To avoid having to import the fmt package to use println i decided to "create" my own.
func Println(t ...interface{}) {
	fmt.Fprintln(&stdout, t...)
}

// Printf :: Just a normal Printf.
// To avoid having to import the fmt package to use printf i decided to "create" my own.
func Printf(format string, t ...interface{}) {
	fmt.Fprintf(&stdout, format, t...)
}

type endl struct{}

// endline :: In order to avoid writing at all times "printer.Println" or "fmt.Println", I created this function that will be returned on all normal printer.
func (options *endl) Endl() *endl {
	Println()

	return options
}

func Done(t ...string) *endl {
	var raw = strings.Join(t, " ")

	io.WriteString(&stdout, prefixDone+" "+raw+"\n")

	return &endl{}
}

func Bars(t string) *endl {
	var list = strings.Split(t, "\n")

	for num, txt := range list {
		if num+1 != len(list) {
			io.WriteString(&stdout, " |   "+txt+"\n")
		}
	}

	return &endl{}
}

func Danger(t ...string) *endl {
	var raw = strings.Join(t, " ")

	io.WriteString(&stdout, prefixDanger+" "+raw+"\n")

	return &endl{}
}

func Warning(t ...string) *endl {
	var raw = strings.Join(t, " ")

	io.WriteString(&stdout, prefixWarning+" "+raw+"\n")

	return &endl{}
}

func Info(t ...string) *endl {
	var raw = strings.Join(t, " ")

	io.WriteString(&stdout, prefixInfo+" "+raw+"\n")

	return &endl{}
}

func Fatal(t interface{}) {
	fmt.Fprint(&stdout, prefixFatal, " ")

	switch t.(type) {
	case error:
		fmt.Fprintln(&stderr, t)
	default:
		fmt.Fprintln(&stdout, t)
	}

	os.Exit(0)
}

func ScanQ(t ...string) string {
	var raw = strings.Join(t, " ")

	io.WriteString(&stdout, prefixScan+" "+raw)

	var scanner = bufio.NewReader(&stdin)
	var response, err = scanner.ReadString('\n')

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

type topics struct {
	text   string
	prefix string
}

func NewTopics(t ...string) *topics {
	var raw = strings.Join(t, " ")

	return &topics{text: raw}
}

func (options *topics) Prefix(s ...string) *topics {
	options.prefix = strings.Join(s, " ")

	return options
}

func (options *topics) Default() {
	io.WriteString(&stdout, options.prefix+prefixListDefault+" "+options.text+"\n")
}

func (options *topics) Done() {
	io.WriteString(&stdout, options.prefix+prefixListDone+" "+options.text+"\n")
}

func (options *topics) Danger() {
	io.WriteString(&stdout, options.prefix+prefixListDanger+" "+options.text+"\n")
}

func (options *topics) Warning() {
	io.WriteString(&stdout, options.prefix+prefixListWarning+" "+options.text+"\n")
}

type topline struct {
	*endl
}

func NewTopLine(t ...string) *topline {
	var raw = strings.Join(t, " ")

	io.WriteString(&stdout, prefixTopLine+" "+raw)

	return &topline{}
}

func (options *topline) Done(t ...string) {
	var raw = strings.Join(t, " ")

	options.Clean()
	Done(raw)
}

func (options *topline) Danger(t ...string) {
	var raw = strings.Join(t, " ")

	options.Clean()
	Danger(raw)
}

func (options *topline) Warning(t ...string) {
	var raw = strings.Join(t, " ")

	options.Clean()
	Warning(raw)
}

func (options *topline) Info(t ...string) {
	var raw = strings.Join(t, " ")

	options.Clean()
	Info(raw)
}

func (options *topline) Clean() {
	fmt.Fprint(&stdout, "\033[G\033[K")
}

func ResetSeek() {
	SeekCurrent = 0
}

func (options *topline) Progress(seek int, t ...string) {
	var prefix = Yellow + fmt.Sprintf("[%d/%d]", SeekCurrent, seek) + Reset
	var raw = strings.Join(t, " ")

	atomic.AddInt64(&SeekCurrent, 1)

	options.Clean()

	if int(SeekCurrent) <= seek {
		io.WriteString(&stdout, prefix+" "+raw)
	} else {
		io.WriteString(&stdout, prefix+" "+raw+"\n")
	}
}
