package banner

// Help ::
var Help = `wprecon (Wordpress Recon) is a scanner based on wpscan, only done in golang to get better performance!

Usage:
  wprecon [flags]

Flags:
  -u, --url string           Target URL (Ex: http(s)://example.com/). (Required)
      --users-enumerate      Use the supplied mode to enumerate Users.
      --plugins-enumerate    Use the supplied mode to enumerate Plugins.
      --themes-enumerate    Use the supplied mode to enumerate Themes.
  -d, --detection-waf        I will try to detect if the target is using any WAF.
      --no-check-wp          Will skip wordpress check on target.
      --random-agent         Use randomly selected HTTP(S) User-Agent header value.
      --tor                  Use Tor anonymity network
      --disable-tls-checks   Disables SSL/TLS certificate verification.
  -h, --help                 help for wprecon
  -v, --verbose              Verbosity mode.
`
