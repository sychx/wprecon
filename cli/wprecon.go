package cli

import (
	"os"

	"github.com/blackcrw/wprecon/internal/pkg/banner"
	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
	"github.com/blackcrw/wprecon/tools/wordpress/fingerprint"
	"github.com/blackcrw/wprecon/tools/wordpress/scanner"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "wprecon",
	Short: "Wordpress Recon",
	Long:  `wprecon (Wordpress Recon) is a scanner based on wpscan, only done in golang to get better performance!`,
	Run: func(cmd *cobra.Command, args []string) {
		tor, _ := cmd.Flags().GetBool("tor")
		target, _ := cmd.Flags().GetString("url")
		verbose, _ := cmd.Flags().GetBool("verbose")
		nocheckwp, _ := cmd.Flags().GetBool("no-check-wp")
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
			/* WP :: Wordpress */
			WP := fingerprint.Honeypot{
				HTTP:    options,
				Verbose: verbose,
			}
			WP.Detection()
		}

		// ———————————————Wordpress Block——————————————— //
		if !nocheckwp {
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
			EP := scanner.Plugins{
				HTTP:    options,
				Verbose: verbose,
			}

			EP.Enumerate()
		}

		// ———————————————Themes Block——————————————— //
		if themeenumerate {
			/* ET :: Enumeration Theme(s) */
			ET := scanner.Themes{
				HTTP:    options,
				Verbose: verbose,
			}

			ET.Enumerate()
		}

		// ————————————————Users Block———————————————— //
		if userenumerate {
			/* EU :: Enumeration User(s) */
			EU := scanner.Users{
				HTTP:    options,
				Verbose: verbose,
			}

			EU.Enumerate()
		}
	},
}

func init() {
	cobra.OnInitialize(ibanner)

	root.PersistentFlags().StringP("url", "u", "", "Target URL (Ex: http(s)://example.com/). "+printer.Required)
	root.PersistentFlags().BoolP("users-enumerate", "", false, "Use the supplied mode to enumerate Users.")
	root.PersistentFlags().BoolP("plugins-enumerate", "", false, "Use the supplied mode to enumerate Plugins.")
	root.PersistentFlags().BoolP("themes-enumerate", "", false, "Use the supplied mode to enumerate themes.")
	root.PersistentFlags().BoolP("detection-waf", "", false, "I will try to detect if the target is using any WAF Wordpress.")
	root.PersistentFlags().BoolP("detection-honeypot", "", false, "I will try to detect if the target is a honeypot, based on the shodan.")
	root.PersistentFlags().BoolP("random-agent", "", false, "Use randomly selected HTTP(S) User-Agent header value.")
	root.PersistentFlags().BoolP("tor", "", false, "Use Tor anonymity network")
	root.PersistentFlags().BoolP("no-check-wp", "", false, "Will skip wordpress check on target.")
	root.PersistentFlags().BoolP("disable-tls-checks", "", false, "Disables SSL/TLS certificate verification.")
	root.PersistentFlags().BoolP("verbose", "v", false, "Verbosity mode.")

	root.MarkPersistentFlagRequired("url")

	root.SetHelpTemplate(banner.Help)
}

func ibanner() {
	verbose, _ := root.Flags().GetBool("verbose")
	target, _ := root.Flags().GetString("url")
	tor, _ := root.Flags().GetBool("tor")

	if isURL, err := gohttp.IsURL(target); isURL {
		banner.SBanner(target, tor, verbose)
	} else {
		banner.Banner()
		printer.Fatal(err)
	}
}

// Execute ::
func Execute() {
	if err := root.Execute(); err != nil {
		os.Exit(0)
	}
}
