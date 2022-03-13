package printer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
)

const (
	RED       = "\u001b[31;1m"
	BLUE      = "\u001b[34;1m"
	GREEN     = "\u001b[32;1m"
	BLACK     = "\u001b[30;1m"
	WHITE     = "\u001b[37;1m"
	YELLOW    = "\u001b[33;1m"
	MAGENTA   = "\u001b[35;1m"
	CYAN      = "\u001b[36;1m"
	RESET     = "\u001b[0m"
	BOLD      = "\u001b[1m"
	UNDERLINE = "\u001b[4m"
	REVERSED  = "\u001b[7m"

	PREFIX_DONE    = GREEN   + "[✔]" +RESET
	PREFIX_DANGER  = RED     + "[✗]" +RESET
	PREFIX_FATAL   = RED     + "[!]" +RESET
	PREFIX_INFO    = MAGENTA + "[i]" +RESET
	PREFIX_WARNING = YELLOW  + "[!]" +RESET
	PREFIX_SCAN    = YELLOW  + "[?]" +RESET

	PREFIX_LIST_DONE    = GREEN   + "    —" +RESET
	PREFIX_LIST_DANGER  = RED     + "    —" +RESET
	PREFIX_LIST_DEFAULT = WHITE   + "    —" +RESET
	PREFIX_LIST_WARNING = YELLOW  + "    —" +RESET

	REQUIRED = RED    + "(Required)" +RESET
	WARNING  = YELLOW + "(Warning)"  +RESET

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
	stdout.WriteString(doPrintbs(a...) +endl)
}

func Printf(format string, a ...interface{}) {
	stdout.WriteString(fmt.Sprintf(format, a...))
}

func Done(t ...interface{})  {
	stdout.WriteString(PREFIX_DONE + " " + doPrintbs(t...) +endl)
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
	stdout.WriteString(PREFIX_DANGER + " " + doPrintbs(t...) +endl)
}

func Warning(t ...interface{})  {
	stdout.WriteString(PREFIX_WARNING + " " + doPrintbs(t...) +endl)
}

func Info(t ...interface{})  {
	stdout.WriteString(PREFIX_INFO + " " + doPrintbs(t...) +endl)
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