package cmd

import (
	"fmt"
	"strings"

	. "github.com/blackcrw/wprecon/cli/config"
	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
	"github.com/blackcrw/wprecon/pkg/scripts"
	"github.com/blackcrw/wprecon/tools/wordpress/commons"
	"github.com/blackcrw/wprecon/tools/wordpress/enumerate"
	"github.com/blackcrw/wprecon/tools/wordpress/extensions"
	"github.com/blackcrw/wprecon/tools/wordpress/security"
	"github.com/spf13/cobra"
)

func RootOptionsPreRun(cmd *cobra.Command, args []string) {
	target, _ := cmd.Flags().GetString("url")
	tor, _ := cmd.Flags().GetBool("tor")
	force, _ := cmd.Flags().GetBool("force")
	verbose, _ := cmd.Flags().GetBool("verbose")
	randomuseragent, _ := cmd.Flags().GetBool("random-agent")
	tlscertificateverify, _ := cmd.Flags().GetBool("tlscertificateverify")
	scriptsS, _ := cmd.Flags().GetString("scripts")

	InfosWprecon.Force = force
	InfosWprecon.Target = target
	InfosWprecon.Verbose = verbose
	InfosWprecon.OtherInformationsBool["http.options.tor"] = tor
	InfosWprecon.OtherInformationsBool["http.options.randomuseragent"] = randomuseragent
	InfosWprecon.OtherInformationsBool["http.options.tlscertificateverify"] = tlscertificateverify

	InfosWprecon.OtherInformationsString["scripts.name"] = scriptsS

	response := gohttp.SimpleRequest(target)

	InfosWprecon.OtherInformationsString["target.http.index.raw"] = response.Raw
}

func RootOptionsRun(cmd *cobra.Command, args []string) {
	aggressivemode, _ := cmd.Flags().GetBool("aggressive-mode")
	detectionwaf, _ := cmd.Flags().GetBool("detection-waf")

	if confidence := wordpresscheck(); confidence >= 49.9 && !InfosWprecon.Force {
		confidenceString := fmt.Sprintf("%.2f%%", confidence)
		printer.Done("Wordpress confirmed with", confidenceString, "confidence!")
		printer.Println()
	} else if confidence < 49.9 && confidence > 33.3 && !InfosWprecon.Force {
		confidenceString := fmt.Sprintf("%.2f%%", confidence)

		if q := printer.ScanQ("I'm not absolutely sure that this target is using wordpress!", confidenceString, "chance. do you wish to continue ? [Y]es | [n]o : "); q != "y" {
			printer.Fatal("Exiting...")
		}
		printer.Println()
	} else {
		printer.Fatal("This target is not running wordpress!")
	}

	if detectionwaf || aggressivemode {
		security.WAFAgressiveDetection()
		printer.Println()
	}

	if scriptsS := InfosWprecon.OtherInformationsString["scripts.name"]; scriptsS != "" {
		L, _ := scripts.Initialize(scriptsS)

		scripts.Run(L)
	}

	if InfosWprecon.Verbose {
		if sitemapResponse := commons.Sitemap(); InfosWprecon.Verbose {
			printer.Warning("Sitemap.xml found:", sitemapResponse.URL.Full)
		}

		if robotsResponse := commons.Robots(); robotsResponse.Raw != "" {
			printer.Warning("Robots.txt file text:")
			printer.Println(robotsResponse.Raw)
		}

		printer.Println()
	}

	if enumP := enumerate.UsersEnumeratePassive(); len(enumP) > 0 && !aggressivemode {
		printer.Done(":: Username(s) Enumerate Passive Mode ::")
		for _, username := range enumP {
			printer.Done(username)
		}
		printer.Println()
	} else if enumA := enumerate.UsersEnumerateAgressive(); len(enumA) > 0 && aggressivemode {
		printer.Done(":: Username(s) Enumerate Agressive Mode ::")
		for _, username := range enumA {
			printer.Done(username)
		}
		printer.Println()
	} else if len(enumA) <= 0 && aggressivemode {
		printer.Danger("Unfortunately no user was found.")
		printer.Println()
	} else {
		printer.Danger("Unfortunately no user was found. Try to use agressive mode: --agressive-mode")
		printer.Println()
	}

	if enumP := enumerate.PluginsEnumeratePassive(); len(enumP) > 0 && !aggressivemode {
		printer.Done(":: Plugin(s) Enumerate Passive Mode ::")
		for name, version := range enumP {
			pluginenum(name, version)
		}
		printer.Println()
	} else if enumA := enumerate.PluginsEnumerateAgressive(); len(enumA) > 0 && aggressivemode {
		printer.Done(":: Plugin(s) Enumerate Agressive Mode ::")
		for name, version := range enumA {
			pluginenum(name, version)
		}
		printer.Println()
	} else if len(enumA) <= 0 && aggressivemode {
		printer.Danger("Unfortunately I was unable to passively list any plugin.")
		printer.Println()
	} else {
		printer.Danger("Unfortunately I was unable to passively list any plugin. Try to use aggressive mode: --aggressive-mode")
		printer.Println()
	}

	if enumP := enumerate.ThemesEnumeratePassive(); len(enumP) > 0 && !aggressivemode {
		printer.Done(":: Theme(s) Enumerate Passive Mode ::")
		for name, version := range enumP {
			printer.Done("Version:", version+"\t", "Plugins:", name)
		}
		if InfosWprecon.Verbose {
			printer.Warning("Unfortunately wprecon doesn't have vulns for themas *yet*.")
		}
		printer.Println()
	} else if enumA := enumerate.ThemesEnumerateAgressive(); len(enumA) > 0 && aggressivemode {
		printer.Done(":: Theme(s) Enumerate Agressive Mode ::")
		for name, version := range enumA {
			printer.Done("Version:", version+"\t", "Plugin:", name)
		}
		if InfosWprecon.Verbose {
			printer.Warning("Unfortunately wprecon doesn't have vulns for themas *yet*.")
		}
		printer.Println()
	} else if len(enumA) <= 0 && aggressivemode {
		printer.Danger("Unfortunately I was unable to passively list any theme.")
		printer.Println()
	} else {
		printer.Danger("Unfortunately I was unable to passively list any theme. Try to use aggressive mode: --aggressive-mode")
		printer.Println()
	}
}

// Detection :: This function should be used to perform wordpress detection.
// "How does this detection work?", I decided to make a 'percentage system' where I will check if each item in a list exists... and if that item exists it will add +1 to accuracy.
// With "16.6" hits he says that wordpress is already detected. But it opens up an opportunity for you to choose whether to continue or not, because you are not 100% sure.
func wordpresscheck() float32 {
	var confidence float32
	var payloads = [...]string{
		`<meta name="generator content="WordPress`,
		`<a href="http://www.wordpress.com">Powered by WordPress</a>`,
		`<link rel='https://api.wordpress.org/'`}

	if has, response := commons.AdminPage(); has == "true" {
		printer.Done("The admin page found:", response.URL.Full)
		confidence++
	} else if has == "redirect" {
		printer.Warning("The admin page is being redirected to:", response.Response.Header.Get("Location"))
		confidence++
	}

	if response := commons.DirectoryPlugins(); response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		confidence++
	}
	if response := commons.DirectoryThemes(); response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		confidence++
	}

	for _, payload := range payloads {
		if strings.Contains(InfosWprecon.OtherInformationsString["target.http.index.raw"], payload) {
			confidence++
		}
	}

	return confidence / 6 * 100
}

func pluginenum(name string, version string) {
	printer.Done("Version:", version+"\t", "Plugin:", name)

	pntl := printer.NewTopLine("Find Vuln...")
	if x := extensions.GetVuln(name, version); len(x.Vulnerabilities) > 0 {
		pntl.Done("Vuln Title:", x.Vulnerabilities[0].Title)
		printer.Done("Vuln Title:", x.Vulnerabilities[0].Version)
		printer.Done("Vuln Plublish:", x.Vulnerabilities[0].Published)

		for _, value := range x.Vulnerabilities[0].References {
			printer.Done("Reference(s):", value)
		}
	} else {
		pntl.Danger("Vuln not found...")
	}
}
