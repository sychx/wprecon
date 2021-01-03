package banner

// Help ::
var Help = `wprecon (Wordpress Recon) is a scanner based on wpscan, only done in golang to get better performance!

Usage:
  wprecon [flags]

Flags:
  -u, --url string           Target URL (Ex: http(s)://example.com/). (Required)
      --detection-waf        I will try to detect if the target is using any WAF.
      --detection-honeypot   I will try to detect if the target is a honeypot, based on the shodan.
      --users-enumerate      Use the supplied mode to enumerate Users.
      --themes-enumerate     Use the supplied mode to enumerate Themes.
      --plugins-enumerate    Use the supplied mode to enumerate Plugins.
      --fuzzer-backup        Performs a fuzzing to try to find the backup file if it exists.
      --no-check-wp          Will skip wordpress check on target.
      --random-agent         Use randomly selected HTTP(S) User-Agent header value.
      --tor                  Use Tor anonymity network
      --disable-tls-checks   Disables SSL/TLS certificate verification.
  -h, --help                 help for wprecon
  -v, --verbose              Verbosity mode.
`
