package cmd

import (
	"fmt"
	"strings"

	"github.com/blackbinn/wprecon/internal/database"
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

	if confidence := extensions.WordpressCheck(); confidence >= 40.0 {
		confidenceString := fmt.Sprintf("%.2f%%", confidence)
		printer.Done("WordPress confirmed with", confidenceString, "confidence!").L()
	} else if confidence < 40.0 && confidence > 15.0 && !database.Memory.GetBool("Force") {
		confidenceString := fmt.Sprintf("%.2f%%", confidence)

		if q := printer.ScanQ("I'm not absolutely sure that this target is using wordpress!", confidenceString, "chance. do you wish to continue ? [Y]es | [n]o : "); q != "y" && q != "\n" {
			printer.Fatal("Exiting...")
		}
		printer.Println()
	} else if confidence < 15.0 && !database.Memory.GetBool("Force") {
		printer.Fatal("This target is not running wordpress!")
	}

	if detectionwaf || aggressivemode {
		if waf := security.WAFAggressiveDetection(); waf != nil {
			name := strings.ReplaceAll(waf.URL.Directory, database.Memory.GetString("HTTP wp-content")+"/plugins/", "")
			name = strings.ReplaceAll(name, "/", "")
			name = strings.ReplaceAll(name, "-", " ")
			name = strings.Title(name)

			printer.Done("Web Application Firewall (WAF):", name, "(Aggressive Detection)")
			printer.List("Location:", waf.URL.Full).D()
			printer.List("Status Code:", fmt.Sprint(waf.Response.Status)).D()

			if importantfile := text.GetOneImportantFile(waf.Raw); strings.Contains(importantfile, "<!-- Avoid the directory listing. -->") && importantfile != "" {
				response := gohttp.SimpleRequest(database.Memory.GetString("Target"), waf.URL.Directory+importantfile)
				if readme := text.GetVersionStableTag(response.Raw); len(readme) != 0 {
					printer.List("Version:", readme[1]).D()
				} else if changelog := text.GetVersionChangelog(response.Raw); len(changelog) != 0 {
					printer.List("Version:", changelog[1]).D()
				}
			} else {
				if response := gohttp.SimpleRequest(database.Memory.GetString("Target"), waf.URL.Directory+"readme.txt"); response.Response.StatusCode == 200 && response.Raw != "" {
					if readme := text.GetVersionStableTag(response.Raw); len(readme) != 0 {
						printer.List("Version:", readme[1]).D()
					} else if changelog := text.GetVersionChangelog(response.Raw); len(changelog) != 0 {
						printer.List("Version:", changelog[1]).D()
					}
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

	if names := database.Memory.GetString("Scripts List Names"); names != "" {
		for _, name := range strings.Split(names, ",") {
			printer.Done("Running Script:", name)

			s := scripts.NewScript()
			s.UseScript(name)
			s.Run()

			printer.Println()
		}
	}

	if wordpressVersion := enumerate.WordpressVersionPassive(); wordpressVersion != "" {
		printer.Done("WordPress Version:")
		printer.List("Version:", wordpressVersion).D().L()
	}

	p := enumerate.NewPlugins()
	t := enumerate.NewThemes()

	switch aggressivemode {
	case false:
		if users, method, URL := enumerate.UsersEnumeratePassive(); len(users) > 0 {
			printer.Done("WordPress Users:")
			for _, username := range users {
				printer.List(username, "("+method+")").D()
			}
			printer.List("All users were found at:", URL).D().L()
		} else {
			printer.Danger("Unfortunately no user was found. Try to use agressive mode: --agressive-mode").L()
		}

		if plugins := p.Passive(); len(plugins) > 0 {
			for _, plugin := range plugins {
				matchs := strings.Split(plugin[2], "ˆ")

				if plugin[1] != "" {
					printer.Done("Plugin:", plugin[0], "(Enumerate Passive Mode)")
					printer.List("Location:", database.Memory.GetString("Target")+database.Memory.GetString("HTTP wp-content")+"/plugins/"+plugin[0]+"/").D()
					printer.List("Version:", plugin[1]).D()
					printer.List(fmt.Sprint(len(matchs)) + " Match(s):").D()
					for _, match := range matchs {
						printer.List(match + ", '" + strings.Split(match, "?")[1] + "'").Prefix("  ").D()
					}
					pluginvulnenum(plugin[0], plugin[1])
				} else {
					printer.Done("Plugin:", plugin[0], "(Enumerate Passive Mode)")
					printer.List("Location:", database.Memory.GetString("Target")+database.Memory.GetString("HTTP wp-content")+"/plugins/"+plugin[0]+"/").D()
					printer.List(fmt.Sprint(len(matchs)) + " Match(s):").D()
					for _, match := range matchs {
						printer.List(match).Prefix("  ").D()
					}
					printer.List("Version: Unidentified version").D().L()
				}
			}
		} else {
			printer.Danger("Unfortunately I was unable to passively list any plugin. Try to use aggressive mode: --aggressive-mode").L()
		}

		if themesEnumerateP := t.Passive(); len(themesEnumerateP) > 0 {
			for _, theme := range themesEnumerateP {
				matchs := strings.Split(theme[2], "ˆ")

				if theme[1] != "" {
					printer.Done("Theme:", theme[0], "(Enumerate Passive Mode)")
					printer.List("Location:", database.Memory.GetString("Target")+database.Memory.GetString("HTTP wp-content")+"/themes/"+theme[0]+"/").D()
					printer.List("Version:", theme[1]).D()
					printer.List(fmt.Sprint(len(matchs)) + " Match(s):").D()
					for _, match := range matchs {
						printer.List(match + ", '" + strings.Split(match, "?")[1] + "'").Prefix("  ").D()
					}
					if database.Memory.GetBool("Verbose") {
						printer.List("Unfortunately wprecon doesn't have vulns for themas *yet*.").Warning()
					}
				} else {
					printer.Done("Theme:", theme[0], "(Enumerate Passive Mode)")
					printer.List("Location:", database.Memory.GetString("Target")+database.Memory.GetString("HTTP wp-content")+"/themes/"+theme[0]+"/").D()
					printer.List(fmt.Sprint(len(matchs)) + " Match(s):").D()
					for _, match := range matchs {
						printer.List(match).Prefix("  ").D()
					}
					printer.List("Version: Unidentified version").D()
				}
			}
			printer.Println()
		} else {
			printer.Danger("Unfortunately I was unable to passively list any theme. Try to use aggressive mode: --aggressive-mode").L()
		}

	case true:
		if sitemapResponse := commons.Sitemap(); sitemapResponse.Response.StatusCode == 200 {
			printer.Warning("Sitemap.xml found:", sitemapResponse.URL.Full).L()
		}
		if robotsResponse := commons.Robots(); robotsResponse.Response.StatusCode == 200 {
			printer.Warning("Robots.txt file text:")
			printer.Bars(robotsResponse.Raw).L()
		}

		if users, method, URL := enumerate.UsersEnumerateAgressive(); len(users) > 0 {
			printer.Done("WordPress Users:")
			for _, username := range users {
				printer.List(username, "("+method+")").D()
			}
			printer.List("All users were found at:", URL).D().L()
		} else {
			printer.Danger("Unfortunately no user was found.").L()
		}

		if plugins := p.Aggressive(); len(plugins) > 0 {
			for _, plugin := range plugins {
				matchs := strings.Split(plugin[2], "ˆ")

				if plugin[1] != "" {
					printer.Done("Plugin:", plugin[0], "(Enumerate Aggressive Mode)")
					printer.List("Location:", database.Memory.GetString("Target")+database.Memory.GetString("HTTP wp-content")+"/plugins/"+plugin[0]+"/").D()
					printer.List("Version:", plugin[1]).D()
					printer.List(fmt.Sprint(len(matchs)) + " Match(s):").D()
					for _, match := range matchs {
						printer.List(match + ", '" + strings.Split(match, "?")[1] + "'").Prefix("  ").D()
					}
					pluginvulnenum(plugin[0], plugin[1])
				} else {
					printer.Done("Plugin:", plugin[0], "(Enumerate Aggressive Mode)")
					printer.List("Location:", database.Memory.GetString("Target")+database.Memory.GetString("HTTP wp-content")+"/plugins/"+plugin[0]+"/").D()
					printer.List(fmt.Sprint(len(matchs)) + " Match(s):").D()
					for _, match := range matchs {
						printer.List(match).Prefix("  ").D()
					}
					printer.List("Version: Unidentified version").D().L()
				}
			}
		} else {
			printer.Danger("Unfortunately I was unable to passively list any plugin.").L()
		}

		if themes := t.Aggressive(); len(themes) > 0 {
			for _, theme := range themes {
				matchs := strings.Split(theme[2], "ˆ")

				if theme[1] != "" {
					printer.Done("Theme:", theme[0], "(Enumerate Aggressive Mode)")
					printer.List("Location:", database.Memory.GetString("Target")+database.Memory.GetString("HTTP wp-content")+"/themes/"+theme[0]+"/").D()
					printer.List("Version:", theme[1]).D()
					printer.List(fmt.Sprint(len(matchs)) + " Match(s):").D()
					for _, match := range matchs {
						printer.List(match).Prefix("  ").D()
					}
					if database.Memory.GetBool("Verbose") {
						printer.List("Unfortunately wprecon doesn't have vulns for themas *yet*.").Warning()
					}
				} else {
					printer.Done("Theme:", theme[0], "(Enumerate Aggressive Mode)")
					printer.List("Location:", database.Memory.GetString("Target")+database.Memory.GetString("HTTP wp-content")+"/themes/"+theme[0]+"/").D()
					printer.List(fmt.Sprint(len(matchs)) + " Match(s):").D()
					for _, match := range matchs {
						printer.List(match).Prefix("  ").D()
					}
					printer.List("Version: Unidentified version").D()
				}
			}
			printer.Println()
		} else {
			printer.Danger("Unfortunately I was unable to passively list any theme.").L()
		}
	}
}

func RootOptionsPostRun(cmd *cobra.Command, args []string) {
	printer.Info("Other interesting information:").L()

	if database.Memory.GetString("HTTP Server") != "" || database.Memory.GetString("HTTP PHP Version") != "" {
		printer.Done("Target information(s):")
		if server := database.Memory.GetString("HTTP Server"); server != "" {
			printer.List("Server:", server).D()
		}
		if version := database.Memory.GetString("HTTP PHP Version"); version != "" {
			printer.List("PHP Version:", version).D()
		}
		if version := database.Memory.GetString("HTTP WordPress Version"); version != "" {
			printer.List("WordPress Version:", version).D()
		}

		printer.Println()
	}

	if len(database.Memory.GetSlice("HTTP Index Of's")) > 0 {
		printer.Done("Index Of's:")
		for _, indexofs := range database.Memory.GetSlice("HTTP Index Of's") {
			printer.List(indexofs).D()
		}
		printer.Println()
	}

	if status, found := commons.XMLRPC(); status != "False" {
		switch found {
		case "Link tag":
			printer.Done("XML-RPC Possibly enabled:")
		default:
			printer.Done("XML-RPC Enabled:")
			printer.List("Status:", status).D()
		}

		printer.List("Location:", database.Memory.GetString("Target")+"xmlrpc.php").D()
		printer.List("Found By:", found).D().L()
	}

	if URL := database.Memory.GetString("HTTP Admin Page"); URL != "" {
		printer.Done("Admin Page Found:")
		printer.List("Location:", URL).D()
		printer.List("Found by: Access").D().L()
	}

	if response := commons.Readme(); response.Response.StatusCode == 200 {
		printer.Done("WordPress Readme:")
		printer.List("Location:", response.URL.Full).D()
		printer.List("Found by: Access").D().L()
	}

	if raw := database.Memory.GetString("HTTP wp-content/uploads Index Of Raw"); raw != "" {
		if list := extensions.FindBackupFileOrPath(raw); len(list) > 0 {
			printer.Done("File or Path backup:")
			for _, path := range list {
				printer.List(database.Memory.GetString("Target") + database.Memory.GetString("HTTP wp-content") + "/uploads/" + path).Done()
			}
			printer.Println()
		}
	}

	printer.Done("Total requests:", fmt.Sprint(database.Memory.GetInt("HTTP Total")))
}

func pluginvulnenum(name string, version string) {
	if vuln := extensions.GetVuln(name, version); len(vuln.Vulnerabilities) > 0 {
		printer.List("Vuln:", vuln.Vulnerabilities[0].Title).Warning()
		printer.List("Report date:", vuln.Vulnerabilities[0].Published).Prefix("  ").D()

		for _, value := range vuln.Vulnerabilities[0].References {
			printer.List("Reference(s):", value).Prefix("  ").D()
		}
	} else {
		printer.List("I have not found any vulnerability for this version.").Danger()
	}

	printer.Println()
}
