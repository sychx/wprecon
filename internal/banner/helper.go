package banner

import (
	"fmt"

	"github.com/blackcrw/wprecon/internal/printer"
)

func format_green(in string) string {
    return printer.GREEN + in + printer.RESET
}

// BannerHelpRoot :: Root help banner
var BannerHelpRoot = fmt.Sprintf(`WPRecon, is a tool for the recognition of vulnerabilities and blackbox information for wordpress.

Usage:
  wprecon %s

Subcommands:
  fuzzer, fuzz               Fuzzing directory or Passwords.

Flags:
  -h, --help                 help for wprecon.
  -u, --url %s         Target URL (Ex: http(s)://example.com/). %s
  -f, --force                Forces wprecon to not check if the target is running WordPress and forces other executions.
  -A, --aggressive-mode      Activates the aggressive mode of wprecon.
      --detection-waf        I will try to detect if the target is using any WAF.
      --random-agent         Use randomly selected HTTP(S) User-Agent header value.
      --tor                  Use Tor anonymity network.
      --wp-content-dir %s In case the wp-content directory is customized. (Default: wp-content)
      --http-sleep %s You can make each request slower, if there is a WAF, it can make it difficult for it to block you. (default: 0)
      --disable-tls-checks   Disables SSL/TLS certificate verification.
  -v, --verbose              Verbosity mode.

Example:
  wprecon -u "https://xxxxxxxx.com" --detection-waf
  wprecon -u "https://xxxxxxxx.com" --aggressive-mode
`,format_green("[flags]"),
	format_green("[target]"),
	printer.REQUIRED,
	format_green("[dir]"),
	format_green("[seconds]"))

// BannerHelpFuzzer :: Fuzzer subcommand help banner
var BannerHelpFuzzer = fmt.Sprintf(`WPRecon, is a tool for the recognition of vulnerabilities and blackbox information for wordpress.

Usage:
  wprecon fuzzer %s

Flags:
      --backup-file            Performs a fuzzing to try to find the backup file if it exists.
  -U, --usernames %s       Set usernames attack passwords.
  -P, --passwords %s  Set wordlist attack passwords.
  -M, --method-attack %s Forces the use of a non-standard attack method (XML-RPC). Available methods: xml-rpc, wp-login
      --p-prefix %s        Sets a prefix for all passwords in the wordlist.
      --p-suffix %s        Sets a suffix for all passwords in the wordlist.

Global Flags:
  -h, --help                   help for wprecon.
  -u, --url %s           Target URL (Ex: http(s)://example.com/). %s
  -f, --force                  Forces wprecon to not check if the target is running WordPress and forces other executions.
      --random-agent           Use randomly selected HTTP(S) User-Agent header value.
      --tor                    Use Tor anonymity network.
      --disable-tls-checks     Disables SSL/TLS certificate verification.
      --http-sleep %s   You can make each request slower, if there is a WAF, it can make it difficult for it to block you. (default: 0)
  -v, --verbose                Verbosity mode.

Example:
  wprecon fuzz -u "https://xxxxxxxx.com" -U user -P $HOME/wordlist/rockyou.txt
  wprecon fuzz -u "https://xxxxxxxx.com" -U user1,user2,user3 -P $HOME/wordlist/rockyou.txt
  wprecon fuzz -u "https://xxxxxxxx.com" --backup-file --random-agent
`, format_green("[flags]"),
	format_green("[list]"),
	format_green("[file-path]"),
	format_green("[attack]"),
	format_green("[text]"),
	format_green("[text]"),
	format_green("[target]"),
	printer.REQUIRED,
	format_green("[seconds]"))
