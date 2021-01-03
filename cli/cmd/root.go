package cmd

import (
	"fmt"

	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
	"github.com/blackcrw/wprecon/tools/wordpress/enumerate"
	"github.com/blackcrw/wprecon/tools/wordpress/fingerprint"
	"github.com/spf13/cobra"
)

// RootOptionsRun ::
func RootOptionsRun(cmd *cobra.Command, args []string) {
	tor, _ := cmd.Flags().GetBool("tor")
	target, _ := cmd.Flags().GetString("url")
	verbose, _ := cmd.Flags().GetBool("verbose")
	force, _ := cmd.Flags().GetBool("force")
	detectionwaf, _ := cmd.Flags().GetBool("detection-waf")
	detectionhoneypot, _ := cmd.Flags().GetBool("detection-honeypot")
	randomuseragent, _ := cmd.Flags().GetBool("random-agent")
	userenumerate, _ := cmd.Flags().GetBool("users-enumerate")
	pluginenumerate, _ := cmd.Flags().GetBool("plugins-enumerate")
	themeenumerate, _ := cmd.Flags().GetBool("themes-enumerate")
	tlscertificateverify, _ := cmd.Flags().GetBool("disable-tls-verify")

	options := &gohttp.HTTPOptions{
		URL: gohttp.URLOptions{
			Simple: target,
		},
		Options: gohttp.Options{
			Tor:                  tor,
			RandomUserAgent:      randomuseragent,
			TLSCertificateVerify: tlscertificateverify,
		},
	}

	if detectionhoneypot {
		/* HP :: Honeypot */
		HP := fingerprint.Honeypot{
			HTTP:    options,
			Verbose: verbose,
		}
		HP.Detection()
	}

	// ———————————————Wordpress Block——————————————— //
	if !force {
		/* WP :: Wordpress */
		WP := fingerprint.Wordpress{
			HTTP:    options,
			Verbose: verbose,
		}

		WP.Detection()
	}

	// ————————WebApplicationFirewall Block———————— //
	if detectionwaf {
		/* WAF :: Web Application Firewall */
		WAF := fingerprint.WebApplicationFirewall{
			HTTP:    options,
			Verbose: verbose,
		}

		WAF.Detection()
	}

	// ———————————————Plugins Block—————————————— //
	if pluginenumerate {
		/* EP :: Enumeration Plugin(s) */
		EP := enumerate.Plugins{
			HTTP:    options,
			Verbose: verbose,
		}

		EP.Enumerate()
	}

	// ———————————————Themes Block——————————————— //
	if themeenumerate {
		/* ET :: Enumeration Theme(s) */
		ET := enumerate.Themes{
			HTTP:    options,
			Verbose: verbose,
		}

		ET.Enumerate()
	}

	// ————————————————Users Block———————————————— //
	if userenumerate {
		/* EU :: Enumeration User(s) */
		EU := enumerate.Users{
			HTTP:    options,
			Verbose: verbose,
		}

		EU.Enumerate()
	}

	printer.Done("Total requests:", fmt.Sprint(options.TotalRequests))
}
