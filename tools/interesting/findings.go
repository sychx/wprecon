package interesting

import (
	"fmt"

	"github.com/blackcrw/wprecon/internal/database"
	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/net"
	"github.com/blackcrw/wprecon/internal/printer"
)

func WPCron() models.InformationsModel {
	var http = net.NewNETClient()
	http.SetURL(database.Memory.GetString("Target")).SetURLDirectory(database.Memory.GetString("HTTP wp-content") + "/wp-cron.php")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))
	http.OnFirewallDetection(true)

	var response, err = http.Runner()

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	return models.InformationsModel{Url: response.URL.Full, Status: response.Response.StatusCode, Raw: response.Raw, Confidence: 100, FoundBy: "Direct Access"}
}

func PHPDisabled() models.InformationsModel {
	var http = net.NewNETClient()
	http.SetURL(database.Memory.GetString("Target")).SetURLDirectory(database.Memory.GetString("HTTP wp-content") + "/wp-includes/version.php")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))
	http.OnFirewallDetection(true)

	var response, err = http.Runner()
	
	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}
	
	return models.InformationsModel{Url: response.URL.Full, Status: response.Response.StatusCode, Raw: response.Raw, Confidence: 100, FoundBy: "Direct Access"}
}

