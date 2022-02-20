package interesting

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/blackcrw/wprecon/internal/database"
	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/net"
	"github.com/blackcrw/wprecon/internal/printer"
	"github.com/blackcrw/wprecon/internal/text"
	"github.com/blackcrw/wprecon/internal/wordlist"
)

func WordpressFirewall() (*models.InterestingModel, error) {
	for _, firewall_name := range wordlist.WPFirewall {
		var net_client = net.NewNETClient()
		net_client.SetURL(database.Memory.GetString("Options URL"))
		net_client.SetURLDirectory(database.Memory.GetString("HTTP wp-content") + "/plugins/" + firewall_name)
		net_client.OnRandomUserAgent(database.Memory.GetBool("Options Random Agent"))

		var net_response, net_err = net_client.Runner()

		if net_err != nil {
			return &models.InterestingModel{}, net_err
		}

		if net_response.Response.StatusCode == 200 || net_response.Response.StatusCode == 403 {
			for _, important_file := range text.FindImportantFiles(net_response.Raw) {
				var net_response_important = net.SimpleRequest(net_response.URL.Full+important_file[1])

				if version_stable_tag := text.GetVersionByStableTag(net_response_important.Raw); len(version_stable_tag) != 0 {
					return &models.InterestingModel{Name: firewall_name, Url: net_response.URL.Full, FoundBy: "Important Files - Stable Tag", Status: net_response.Response.StatusCode }, nil
				} else if version_changelog := text.GetVersionByChangelog(net_response_important.Raw); len(version_changelog) != 0 {
					return &models.InterestingModel{Name: firewall_name, Url: net_response.URL.Full, FoundBy: "Important Files - Changelog", Status: net_response.Response.StatusCode }, nil
				}
			}
		}
	}

	return &models.InterestingModel{}, nil
}

func WordPressVersion() string {
	var raw = database.Memory.GetString("HTTP Index Raw")

	var regxp = regexp.MustCompile("<meta name=\"generator\" content=\"WordPress ([0-9.-]*).*?")

	for _, slice_bytes := range regxp.FindAllSubmatch([]byte(raw), -1) {
		var version = string(slice_bytes[1][:])

		database.Memory.SetString("HTTP WordPress Version", version)
	}

	return database.Memory.GetString("HTTP WordPress Version")
}

func WordpressCheck() float32 {
	var wait_group sync.WaitGroup
	var confidence float32

	var payloads = [4]string{
		"<meta name=\"generator content=\"WordPress",
		"<a href=\"http://www.wordpress.com\">Powered by WordPress</a>",
		"<link rel=\"https://api.wordpress.org/",
		"<link rel=\"https://api.w.org/\"",
	}

	wait_group.Add(4)

	go func(){ if check, err := AdminPage();        check.Confidence == 100 { confidence++ } else if err != nil { printer.Danger(fmt.Sprintf("%s", err)) } ; wait_group.Done() }()
	go func(){ if check, err := DirectoryPlugins(); check.Confidence == 100 { confidence++ } else if err != nil { printer.Danger(fmt.Sprintf("%s", err)) } ; wait_group.Done() }()
	go func(){ if check, err := DirectoryThemes();  check.Confidence == 100 { confidence++ } else if err != nil { printer.Danger(fmt.Sprintf("%s", err)) } ; wait_group.Done() }()
	go func(){ if check, err := DirectoryUploads(); check.Confidence == 100 { confidence++ } else if err != nil { printer.Danger(fmt.Sprintf("%s", err)) } ; wait_group.Done() }()
	
	for _, payload := range payloads {
		if strings.Contains(database.Memory.GetString("HTTP Index Raw"), payload) {
			confidence++
		}
	}
	
	wait_group.Wait()

	return confidence / 8 * 100
}

func WPCron() (*models.InterestingModel, error) {
	var http = net.NewNETClient()
	http.SetURL(database.Memory.GetString("Options URL")).SetURLDirectory(database.Memory.GetString("HTTP wp-content") + "/wp-cron.php")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))

	var response, err = http.Runner()

	if err != nil { return &models.InterestingModel{}, err }

	return &models.InterestingModel{Url: response.URL.Full, Status: response.Response.StatusCode, Raw: response.Raw, Confidence: 100, FoundBy: "Direct Access"}, nil
}

func PHPDisabled() (*models.InterestingModel, error) {
	var http = net.NewNETClient()
	http.SetURL(database.Memory.GetString("Options URL")).SetURLDirectory(database.Memory.GetString("HTTP wp-content") + "/wp-includes/version.php")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))

	var response, err = http.Runner()
	
	if err != nil { return &models.InterestingModel{}, err }
	
	return &models.InterestingModel{Url: response.URL.Full, Status: response.Response.StatusCode, Raw: response.Raw, Confidence: 100, FoundBy: "Direct Access"}, nil
}

func XMLRPC() (*models.InterestingModel, error) {
	var http = net.NewNETClient()
	http.SetURL(database.Memory.GetString("Options URL")).SetURLDirectory("xmlrpc.php")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))

	var response, err = http.Runner()

	if err != nil { return &models.InterestingModel{}, err }

	var model = models.InterestingModel{Url: response.URL.Full, Status: response.Response.StatusCode, Raw: response.Raw, Confidence: -1, FoundBy: "Direct Access"}

	if strings.Contains(response.Raw, "XML-RPC server accepts POST requests only.") {
		model.Confidence += 60
	}

	if response.Response.StatusCode == 405 {
		model.Confidence += 20
	}
	
	if strings.Contains(response.Raw, "This error was generated by Mod_Security.") {
		model.Confidence = 80
		model.FoundBy = "Direct Access (Mod_Security)"
	}
	
	if strings.Contains(database.Memory.GetString("HTTP Index Raw"), "xmlrpc.php") {
		model.Confidence += 10
	}

	return &model, nil
}