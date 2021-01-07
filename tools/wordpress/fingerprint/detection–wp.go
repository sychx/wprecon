package fingerprint

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
)

// Wordpress ::
type Wordpress struct {
	HTTP     *gohttp.HTTPOptions
	Verbose  bool
	accuracy float32
}

var wg sync.WaitGroup

// Detection :: This function should be used to perform wordpress detection.
/*
"How does this detection work?", I decided to make a 'percentage system' where I will check if each item in a list exists... and if that item exists it will add +1 to accuracy.
With "25.0" hits he says that wordpress is already detected. But it opens up an opportunity for you to choose whether to continue or not, because you are not 100% sure.
*/
func (options *Wordpress) Detection() {
	var question string

	wg.Add(2)

	go func() {
		options.directory()
		wg.Done()
	}()

	go func() {
		options.htmlcode()
		wg.Done()
	}()

	wg.Wait()

	options.accuracy = options.accuracy / 8 * 100

	accuracyString := fmt.Sprintf("%.2f%%", options.accuracy)

	if options.accuracy >= 62.5 {
		printer.Done("Wordpress confirmed with", accuracyString, "accuracy!")
	} else if options.accuracy < 62.5 && options.accuracy > 25.0 {
		printer.Warning("I'm not absolutely sure that this target is using wordpress!", accuracyString, "chance. do you wish to continue ? [Y/n]:")
		fmt.Print("\r")
		if fmt.Scan(&question); strings.ToLower(question) != "y" {
			printer.Fatal("Exiting...")
		}
	} else {
		printer.Fatal("This target is not running wordpress!")
	}

	printer.Println("")
}

func (options *Wordpress) htmlcode() {
	var payloads = [...]string{
		`<meta name="generator content="WordPress`,
		`<a href="http://www.wordpress.com">Powered by WordPress</a>`,
		`<link rel='https://api.wordpress.org/'`}

	response, err := gohttp.HTTPRequest(options.HTTP)

	if err != nil {
		printer.Fatal(err)
	}

	content, err := ioutil.ReadAll(response.Raw)

	if err != nil {
		printer.Fatal(err)
	}

	for _, value := range payloads {
		if strings.Contains(value, string(content)) {
			options.accuracy++
		}

		time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
	}
}

func (options *Wordpress) directory() {
	var directories = [...]string{
		"wp-content/uploads/",
		"wp-content/plugins/",
		"wp-content/themes/",
		"wp-includes/",
		"wp-admin"}

	for _, directory := range directories {
		options.HTTP.URL.Directory = directory

		request, err := gohttp.HTTPRequest(options.HTTP)

		if err != nil {
			printer.Fatal(err)
		}

		body, err := ioutil.ReadAll(request.Raw)

		if err != nil {
			printer.Fatal(err)
		}

		if directory == "wp-admin" && request.StatusCode == 200 || request.StatusCode == 403 {
			printer.Warning("Status Code:", fmt.Sprint(request.StatusCode), "—", "URL:", request.URL.Full)
			options.accuracy++
		} else if request.StatusCode == 200 || request.StatusCode == 403 {
			if strings.Contains(string(body), "Index Of") {
				printer.Done("Listing enable:", request.URL.Full)
			}

			options.accuracy++
		}
	}

}
