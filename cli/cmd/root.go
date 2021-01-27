package cmd

import (
	"fmt"
	"strings"

	. "github.com/blackbinn/wprecon/cli/config"
	"github.com/blackbinn/wprecon/pkg/gohttp"
	"github.com/blackbinn/wprecon/pkg/printer"
	"github.com/blackbinn/wprecon/pkg/scripts"
	"github.com/blackbinn/wprecon/pkg/text"
	"github.com/blackbinn/wprecon/tools/wordpress/commons"
	"github.com/blackbinn/wprecon/tools/wordpress/enumerate"
	"github.com/blackbinn/wprecon/tools/wordpress/extensions"
	"github.com/blackbinn/wprecon/tools/wordpress/security"
	"github.com/spf13/cobra"
)

func RootOptionsRun(cmd *cobra.Command, args []string) {
	aggressivemode, _ := cmd.Flags().GetBool("aggressive-mode")
	detectionwaf, _ := cmd.Flags().GetBool("detection-waf")

	if confidence := wordpresscheck(); confidence >= 40.0 {
		confidenceString := fmt.Sprintf("%.2f%%", confidence)
		printer.Done("WordPress confirmed with", confidenceString, "confidence!")
		printer.Println()
	} else if confidence < 40.0 && confidence > 15.0 && !InfosWprecon.Force {
		confidenceString := fmt.Sprintf("%.2f%%", confidence)

		if q := printer.ScanQ("I'm not absolutely sure that this target is using wordpress!", confidenceString, "chance. do you wish to continue ? [Y]es | [n]o : "); q != "y" && q != "\n" {
			printer.Fatal("Exiting...")
		}
		printer.Println()
	} else if confidence < 15.0 && !InfosWprecon.Force {
		printer.Fatal("This target is not running wordpress!")
	}

	if detectionwaf || aggressivemode {
		if waf := security.WAFAgressiveDetection(); waf != nil {
			name := strings.ReplaceAll(waf.URL.Directory, "wp-content/plugins/", "")
			name = strings.ReplaceAll(name, "/", "")
			name = strings.ReplaceAll(name, "-", " ")
			name = strings.Title(name)

			printer.Done("Web Application Firewall (WAF):", name, "(Aggressive Detection)")
			printer.List("Location:", waf.URL.Full)
			printer.List("Status Code:", fmt.Sprint(waf.Response.Status))

			if importantfile := text.GetOneImportantFile(waf.Raw); importantfile != "" {
				response := gohttp.SimpleRequest(InfosWprecon.Target, waf.URL.Directory+importantfile)

				if readme := text.GetVersionStableTag(response.Raw); readme != "" {
					printer.Warning("Version \t:", readme)
				} else if changelog := text.GetVersionChangelog(response.Raw); changelog != "" {
					printer.Warning("Version \t:", changelog)
				}
			}

			if scan := printer.ScanQ("Do you wish to continue ?! [Y]es | [n]o : "); scan != "y" && scan != "\n" {
				printer.Fatal("Exiting...")
			}
		} else {
			printer.Warning(":: No WAF was detected! But that doesn't mean it doesn't. ::")
		}

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

	if usersEnumerateP, response := enumerate.UsersEnumeratePassive(); len(usersEnumerateP) > 0 && !aggressivemode {
		printer.Done("WordPress Users:")
		for _, username := range usersEnumerateP {
			printer.List(username, "("+InfosWprecon.OtherInformationsString["target.http.users.method"]+")")
		}
		printer.List("All users were found at:", response.URL.Full)
	} else if usersEnumerateA, response := enumerate.UsersEnumerateAgressive(); len(usersEnumerateA) > 0 && aggressivemode {
		printer.Done("WordPress Users:")
		for _, username := range usersEnumerateA {
			printer.List(username, "("+InfosWprecon.OtherInformationsString["target.http.users.method"]+")")
		}
		printer.List("All users were found at:", response.URL.Full)
	} else if len(usersEnumerateA) <= 0 && aggressivemode {
		printer.Danger("Unfortunately no user was found.")
	} else {
		printer.Danger("Unfortunately no user was found. Try to use agressive mode: --agressive-mode")
	}

	if pluginsEnumerateP := enumerate.PluginsEnumeratePassive(); len(pluginsEnumerateP) > 0 && !aggressivemode {
		for plugin, version := range pluginsEnumerateP {
			printer.Println()
			if version != "" {
				printer.Done("Plugin:", plugin, "(Enumerate Passive Mode)")
				printer.List("Location:", InfosWprecon.Target+"wp-content/plugins/"+plugin+"/")
				printer.List("Version:", version)
				pluginvulnenum(plugin, version)
			} else {
				printer.Done("Plugin:", plugin, "(Enumerate Passive Mode)")
				printer.List("Location:", InfosWprecon.Target+"wp-content/plugins/"+plugin+"/")
				printer.List("Version: Unidentified version")
			}
		}
	} else if pluginsEnumerateA := enumerate.PluginsEnumerateAgressive(); len(pluginsEnumerateA) > 0 && aggressivemode {
		for plugin, version := range pluginsEnumerateA {
			printer.Println()
			if version != "" {
				printer.Done("Plugin:", plugin, "(Enumerate Aggressive Mode)")
				printer.List("Location:", InfosWprecon.Target+"wp-content/plugins/"+plugin+"/")
				printer.List("Version:", version)
				pluginvulnenum(plugin, version)
			} else {
				printer.Done("Plugin:", plugin, "(Enumerate Aggressive Mode)")
				printer.List("Location:", InfosWprecon.Target+"wp-content/plugins/"+plugin+"/")
				printer.List("Version: Unidentified version")
			}
		}
	} else if len(pluginsEnumerateA) <= 0 && aggressivemode {
		printer.Danger("Unfortunately I was unable to passively list any plugin.")
	} else {
		printer.Danger("Unfortunately I was unable to passively list any plugin. Try to use aggressive mode: --aggressive-mode")
	}

	if themesEnumerateP := enumerate.ThemesEnumeratePassive(); len(themesEnumerateP) > 0 && !aggressivemode {
		for theme, version := range themesEnumerateP {
			printer.Println()
			if version != "" {
				printer.Done("Theme:", theme, "(Enumerate Passive Mode)")
				printer.List("Location:", InfosWprecon.Target+"wp-content/themes/"+theme+"/")
				printer.List("Version:", version)

				if InfosWprecon.Verbose {
					printer.List("Unfortunately wprecon doesn't have vulns for themas *yet*.")
				}
			} else {
				printer.Done("Theme:", theme, "(Enumerate Passive Mode)")
				printer.List("Location:", InfosWprecon.Target+"wp-content/themes/"+theme+"/")
				printer.List("Version: Unidentified version")
			}
		}
		printer.Println()
	} else if themesEnumerateA := enumerate.ThemesEnumerateAgressive(); len(themesEnumerateA) > 0 && aggressivemode {
		for theme, version := range themesEnumerateP {
			printer.Println()
			if version != "" {
				printer.Done("Theme:", theme, "(Enumerate Aggressive Mode)")
				printer.List("Location:", InfosWprecon.Target+"wp-content/themes/"+theme+"/")
				printer.List("Version:", version)

				if InfosWprecon.Verbose {
					printer.List("Unfortunately wprecon doesn't have vulns for themas *yet*.")
				}
			} else {
				printer.Done("Theme:", theme, "(Enumerate Aggressive Mode)")
				printer.List("Location:", InfosWprecon.Target+"wp-content/themes/"+theme+"/")
				printer.List("Version: Unidentified version")
			}
		}
		printer.Println()
	} else if len(themesEnumerateA) <= 0 && aggressivemode {
		printer.Danger("Unfortunately I was unable to passively list any theme.")
		printer.Println()
	} else {
		printer.Danger("Unfortunately I was unable to passively list any theme. Try to use aggressive mode: --aggressive-mode")
		printer.Println()
	}
}

func RootOptionsPostRun(cmd *cobra.Command, args []string) {
	printer.Info("Other interesting information:")
	printer.Println()

	if len(InfosWprecon.OtherInformationsSlice["target.http.indexof"]) > 0 {
		printer.Done("Index Of's:")
		for _, indexofs := range InfosWprecon.OtherInformationsSlice["target.http.indexof"] {
			printer.List(indexofs)
		}
		printer.Println()
	}

	if status, _ := commons.XMLRPC(); status != "false" {
		printer.Done("XML-RPC Enabled:")
		printer.List("Location:", InfosWprecon.Target+"xmlrpc.php")
		printer.List("Checked By:", InfosWprecon.OtherInformationsString["target.http.xmlrpc.php.checkedby"])
		printer.Println()
	}

	if URL := InfosWprecon.OtherInformationsString["target.http.admin-page"]; URL != "" {
		printer.Done("Admin Page Found:")
		printer.List("Location:", URL)
		printer.List("Checked by: Access")
		printer.Println()
	}

	if response := commons.Readme(); response.Response.StatusCode == 200 {
		printer.Done("WordPress Readme:")
		printer.List("Location:", response.URL.Full)
		printer.List("Checked by: Access")
		printer.Println()
	}

	printer.Done("Total requests:", fmt.Sprint(InfosWprecon.TotalRequests))
}

// Detection :: This function should be used to perform wordpress detection.
// "How does this detection work?", I decided to make a 'percentage system' where I will check if each item in a list exists... and if that item exists it will add +1 to accuracy.
// With "16.6" hits he says that wordpress is already detected. But it opens up an opportunity for you to choose whether to continue or not, because you are not 100% sure.
func wordpresscheck() float32 {
	var confidence float32
	var payloads = [4]string{
		"<meta name=\"generator content=\"WordPress",
		"<a href=\"http://www.wordpress.com\">Powered by WordPress</a>",
		"<link rel=\"https://api.wordpress.org/",
		"<link rel=\"https://api.w.org/\""}

	if has, _ := commons.AdminPage(); has == "true" || has == "redirect" {
		confidence++
	}
	if response := commons.DirectoryPlugins(); response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		confidence++
	}
	if response := commons.DirectoryThemes(); response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		confidence++
	}
	if response := commons.DirectoryUploads(); response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		confidence++
	}

	for _, payload := range payloads {
		if strings.Contains(InfosWprecon.OtherInformationsString["target.http.index.raw"], payload) {
			confidence++
		}
	}

	return confidence / 8 * 100
}

func pluginvulnenum(name string, version string) {
	if vuln := extensions.GetVuln(name, version); len(vuln.Vulnerabilities) > 0 {
		printer.List("Vuln Title:", vuln.Vulnerabilities[0].Title)
		printer.List("Vuln Plublish:", vuln.Vulnerabilities[0].Published)

		for _, value := range vuln.Vulnerabilities[0].References {
			printer.List("Reference(s):", value)
		}
	} else {
		printer.List("I have not found any vulnerability for this version.")
	}
}
