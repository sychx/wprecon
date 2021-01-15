package banner

import (
	"fmt"

	"github.com/blackcrw/wprecon/pkg/printer"
)

// HelpMain ::
var HelpMain = fmt.Sprintf(`wprecon (Wordpress Recon) is a tool for wordpress exploration!

Usage:
  wprecon [flags]

Subcommands:
  fuzzer, fuzz               Fuzzing directory or Passwords.

Flags:
  -u, --url string           Target URL (Ex: http(s)://example.com/). %s
  -f, --force                Forces wprecon to not check if the target is running WordPress and forces other executions.
      --aggressive-mode      Activates the aggressive mode of wprecon.
      --detection-waf        I will try to detect if the target is using any WAF.
      --random-agent         Use randomly selected HTTP(S) User-Agent header value.
      --tor                  Use Tor anonymity network.
      --disable-tls-checks   Disables SSL/TLS certificate verification.
      --scripts string       Auxiliary scripts.
  -h, --help                 help for wprecon.
  -v, --verbose              Verbosity mode.

Example:
  wprecon -u "https://xxxxxxxx.com" --detection-waf
  wprecon -u "https://xxxxxxxx.com" --detection-waf --aggressive-mode
  wprecon --url "https://xxxxxxxx.com" --detection-waf --scripts honeypot -v
`, printer.Required)

// HelpFuzzer ::
var HelpFuzzer = fmt.Sprintf(`wprecon (Wordpress Recon) is a tool for wordpress exploration!

Usage:
  wprecon fuzzer [flags]

Flags:
      --backup-file          Performs a fuzzing to try to find the backup file if it exists.
  -U, --usernames string     Set usernames attack passwords.
  -P, --passwords wordlist   Set wordlist attack passwords.

Global Flags:
  -u, --url string           Target URL (Ex: http(s)://example.com/). %s
  -f, --force                Forces wprecon to not check if the target is running WordPress and forces other executions.
      --random-agent         Use randomly selected HTTP(S) User-Agent header value.
      --tor                  Use Tor anonymity network.
      --disable-tls-checks   Disables SSL/TLS certificate verification.
      --scripts string       Auxiliary scripts.
  -v, --verbose              Verbosity mode.
  -h, --help                 help for wprecon.

Example:
  wprecon fuzzer --url "https://xxxxxxxx.com" --backup-file
  wprecon fuzzer --url "https://xxxxxxxx.com" --usernames xxx,yyy,zzz --passwords $HOME/wordlist/rockyou.txt
  wprecon fuzzer --url "https://xxxxxxxx.com" --backup-file --random-agent -v
`, printer.Required)
