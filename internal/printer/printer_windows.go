package printer

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
)

const (
	PREFIX_DANGER  = "[✗]"
	PREFIX_FATAL   = "[!]"
	PREFIX_DONE    = "[✔]"
	PREFIX_WARNING = "[!]"
	PREFIX_SCAN    = "[?]"
	PREFIX_INFO    = "[i]"

	PREFIX_LIST_DONE    = "    —"
	PREFIX_LIST_DANGER  = "    —"
	PREFIX_LIST_DEFAULT = "    —"
	PREFIX_LIST_WARNING = "    —"

	REQUIRED = "(Required)"
	WARNING  = "(Warning)"

	endl = "\n"
)

var (
	stdin    = *os.Stdin
	stdout   = *os.Stdout
	stderr   = *os.Stderr
)

func doPrintbs(a ...interface{}) (str string){
	for arg_num, arg := range a {
		if arg_num > 0 { str += " " }
		str += arg.(string)
	}

	return
}

func Println(a ...interface{}) {
	stdout.WriteString(doPrintbs(a...))
}

func Printf(format string, a ...interface{}) {
	stdout.WriteString(fmt.Sprintf(format, a...))
}

func Done(t ...interface{})  {
	stdout.WriteString(PREFIX_DONE + " " + doPrintbs(t...) + endl)
}

func Bars(t string)  {
	var list = strings.Split(t, "\n")

	for num, txt := range list {
		if num+1 != len(list) {
			stdout.WriteString(" |   "+txt+"\n")
		}
	}
}

func Danger(t ...interface{})  {
	stdout.WriteString(PREFIX_DANGER + " " + doPrintbs(t...) + endl)
}

func Warning(t ...interface{})  {
	stdout.WriteString(PREFIX_WARNING + " " + doPrintbs(t...) + endl)
}

func Info(t ...interface{})  {
	stdout.WriteString(PREFIX_INFO + " " + doPrintbs(t...) + endl)
}

func Fatal(t ...interface{}) {
	stderr.WriteString(RED + "[ERROR]" + doPrintbs(t...) +endl)
	syscall.Exit(0)
}

func ScanQ(t ...interface{}) string {
	stdout.WriteString(PREFIX_SCAN +" "+ doPrintbs(t...))

	var response, err = bufio.NewReader(&stdin).ReadString('\n')

	if err != nil {
		Fatal(err)
	}

	response = strings.ToLower(response)

	if response == "\n" {
		return response
	}

	response = strings.Replace(response, "\n", "", -1)

	return response
}
