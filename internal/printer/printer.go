package printer

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
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
	Warningx  = Yellow + "(Warning)" + Reset
)

var (
	prefix_warning = Yellow + "[!]" + Reset
	prefix_danger  = Red + "[✗]" + Reset
	prefix_fatal   = Red + "[!]" + Reset
	prefix_done    = Green + "[✔]" + Reset
	prefix_scan    = Yellow + "[?]" + Reset
	prefix_info    = Magenta + "[i]" + Reset

	prefix_list_warning = Yellow + "    —" + Reset
	prefix_list_default = White + "    —" + Reset
	prefix_list_danger  = Red + "    —" + Reset
	prefix_list_done    = Green + "    —" + Reset

	prefix_top_line = Yellow + "[✲]" + Reset
)

var SeekCurrent int64 = 0

func init() {
	switch runtime.GOOS {
	case "windows":
		prefix_danger = "[✗]"
		prefix_fatal = "[!]"
		prefix_done = "[✔]"
		prefix_warning = "[!]"
		prefix_scan = "[?]"
		prefix_info = "[i]"

		prefix_list_danger = "    —"
		prefix_list_done = "    —"
		prefix_list_default = "    —"
		prefix_list_warning = "    —"

		prefix_top_line = "[✲]"
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

func Done(t ...string)  {
	var raw = strings.Join(t, " ")

	io.WriteString(&stdout, prefix_done+" "+raw+"\n")
}

func Bars(t string)  {
	var list = strings.Split(t, "\n")

	for num, txt := range list {
		if num+1 != len(list) {
			io.WriteString(&stdout, " |   "+txt+"\n")
		}
	}
}

func Danger(t ...string)  {
	var raw = strings.Join(t, " ")

	io.WriteString(&stdout, prefix_danger+" "+raw+"\n")
}

func Warning(t ...string)  {
	var raw = strings.Join(t, " ")

	io.WriteString(&stdout, prefix_warning+" "+raw+"\n")
}

func Info(t ...string)  {
	var raw = strings.Join(t, " ")

	io.WriteString(&stdout, prefix_info+" "+raw+"\n")
}

func Fatal(t interface{}) {
	fmt.Fprint(&stdout, prefix_fatal, " ")

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

	io.WriteString(&stdout, prefix_scan+" "+raw)

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

func (this *topics) Prefix(s ...string) *topics {
	this.prefix = strings.Join(s, " ")

	return this
}

func (this *topics) Default() {
	io.WriteString(&stdout, this.prefix+prefix_list_default+" "+this.text+"\n")
}

func (this *topics) Done() {
	io.WriteString(&stdout, this.prefix+prefix_list_done+" "+this.text+"\n")
}

func (this *topics) Danger() {
	io.WriteString(&stdout, this.prefix+prefix_list_danger+" "+this.text+"\n")
}

func (this *topics) Warning() {
	io.WriteString(&stdout, this.prefix+prefix_list_warning+" "+this.text+"\n")
}

func ResetSeek() {
	SeekCurrent = 0
}