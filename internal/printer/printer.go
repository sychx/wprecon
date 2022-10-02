package printer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
)

var (
	stdin    = *os.Stdin
	stdout   = *os.Stdout
	stderr   = *os.Stderr
)

func doPrintbs(a ...interface{}) (str string){
	for count, arg := range a {
		if count > 0 { str += " " }

		switch arg.(type) {
		case string: str += arg.(string)
		case error: str += arg.(error).Error()
		case int: str += fmt.Sprint(arg.(int))
		}
	}

	return
}

func Println(a ...interface{}) {
	stdout.WriteString(doPrintbs(a...) +"\n")
}

func Printf(format string, a ...interface{}) {
	stdout.WriteString(fmt.Sprintf(format, a...))
}

func Print(a ...interface{}) {
	stdout.WriteString(doPrintbs(a...))
}

func Done(t ...interface{})  {
	stdout.WriteString(PREFIX_DONE + " " + doPrintbs(t...) +"\n")
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
	stdout.WriteString(PREFIX_DANGER + " " + doPrintbs(t...) +"\n")
}

func Warning(t ...interface{})  {
	stdout.WriteString(PREFIX_WARNING + " " + doPrintbs(t...) +"\n")
}

func Info(t ...interface{})  {
	stdout.WriteString(PREFIX_INFO + " " + doPrintbs(t...) +"\n")
}

func Fatal(t ...interface{}) {
	stderr.WriteString(RED + "[ERROR]" + doPrintbs(t...) +"\n")
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