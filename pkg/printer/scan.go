package printer

import (
	"bufio"
	"io"
	"os"
	"strings"

	color "github.com/logrusorgru/aurora"
)

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
