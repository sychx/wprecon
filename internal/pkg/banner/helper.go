package banner

// HelpMain ::
var HelpMain = `wprecon (Wordpress Recon) is a scanner based on wpscan, only done in golang to get better performance!

Usage:
  wprecon [flags]

SubCommands:
  fuzzer, fuzz

Flags:
  -u, --url string           Target URL (Ex: http(s)://example.com/). (Required)
      --detection-waf        I will try to detect if the target is using any WAF.
      --detection-honeypot   I will try to detect if the target is a honeypot, based on the shodan.
      --users-enumerate      Use the supplied mode to enumerate Users.
      --themes-enumerate     Use the supplied mode to enumerate Themes.
      --plugins-enumerate    Use the supplied mode to enumerate Plugins.
      --no-check-wp          Will skip wordpress check on target.
      --random-agent         Use randomly selected HTTP(S) User-Agent header value.
      --tor                  Use Tor anonymity network.
      --disable-tls-checks   Disables SSL/TLS certificate verification.
  -h, --help                 help for wprecon.
  -v, --verbose              Verbosity mode.

Example:
  wprecon --url "https://xxxxxxxx.com" --detection-waf
  wprecon --url "https://xxxxxxxx.com" --detection-waf --detection-honeypot
  wprecon --url "https://xxxxxxxx.com" --detection-waf --detection-honeypot --users-enumerate
`

// HelpFuzzer ::
var HelpFuzzer = `wprecon (Wordpress Recon) is a scanner based on wpscan, only done in golang to get better performance!

Usage:
  wprecon fuzzer [flags]

Flags:
      --backup-file          Performs a fuzzing to try to find the backup file if it exists.
  -h, --help                 help for wprecon.

Global Flags:
  -u, --url string           Target URL (Ex: http(s)://example.com/). (Required)
      --random-agent         Use randomly selected HTTP(S) User-Agent header value.
      --tor                  Use Tor anonymity network.
      --disable-tls-checks   Disables SSL/TLS certificate verification.
  -h, --help                 help for wprecon.
  -v, --verbose              Verbosity mode.

Example:
  wprecon fuzzer --url "https://xxxxxxxx.com" --backup-file
  wprecon fuzzer --url "https://xxxxxxxx.com" --backup-file --random-agent -v
`
