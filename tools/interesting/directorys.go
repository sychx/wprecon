package interesting

import (
	"strings"

	"github.com/blackcrw/wprecon/internal/database"
	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/net"
)

func DirectoryPlugins() (models.InterestingModel, error) {
	var http = net.NewNETClient()
	http.SetURL(database.Memory.GetString("Options URL")).SetURLDirectory(database.Memory.GetString("HTTP wp-content") + "/plugins/")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))

	var response, err = http.Runner()

	if err != nil { return models.InterestingModel{}, err }

	if strings.Contains(response.Raw, "Index of") {
		database.Memory.AddInSlice("HTTP Index Of's", response.URL.Full)
		database.Memory.SetString("HTTP wp-content/plugins Index Of Raw", response.Raw)
	}

	var models_interesting = models.InterestingModel{Url: response.URL.Full, Raw: response.Raw, Confidence: -1, FoundBy: "Direct Access"}

	if response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		models_interesting.Confidence = 100
	}

	return models_interesting, nil
}

func DirectoryUploads() (models.InterestingModel, error) {
	var http = net.NewNETClient().SetURL(database.Memory.GetString("Options URL"))
	http.SetURLDirectory(database.Memory.GetString("HTTP wp-content") + "/uploads/")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))

	var response, err = http.Runner()

	if err != nil { return models.InterestingModel{}, err }

	if strings.Contains(response.Raw, "Index of") {
		database.Memory.AddInSlice("HTTP Index Of's", response.URL.Full)
		database.Memory.SetString("HTTP wp-content/uploads Index Of Raw", response.Raw)
	}

	var models_interesting = models.InterestingModel{Url: response.URL.Full, Raw: response.Raw, Confidence: -1, FoundBy: "Direct Access"}

	if response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		models_interesting.Confidence = 100
	}

	return models_interesting, nil
}

func DirectoryThemes() (models.InterestingModel, error) {
	var http = net.NewNETClient()
	http.SetURL(database.Memory.GetString("Options URL")).SetURLDirectory(database.Memory.GetString("HTTP wp-content") + "/themes/")
	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))

	var response, err = http.Runner()

	if err != nil { return models.InterestingModel{}, err }

	if strings.Contains(response.Raw, "Index of") {
		database.Memory.AddInSlice("HTTP Index Of's", response.URL.Full)
		database.Memory.SetString("HTTP wp-content/themes Index Of Raw", response.Raw)
	}

	var models_interesting = models.InterestingModel{Url: response.URL.Full, Raw: response.Raw, Confidence: -1, FoundBy: "Direct Access"}

	if response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		models_interesting.Confidence = 100
	}

	return models_interesting, nil
}
