package banner

import (
	"fmt"

	"runtime"

	"github.com/blackbinn/wprecon/pkg/printer"
	color "github.com/logrusorgru/aurora"
)

func green(s string) string {
	if runtime.GOOS == "windows" {
		return s
	}

	return color.Green(s).String()
}

//  -t, --threads int            Max number of concurrent.

// HelpMain ::
var HelpMain = fmt.Sprintf(`wprecon (Wordpress Recon) is a tool for wordpress exploration!

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
      --scripts %s       Auxiliary scripts.
  -v, --verbose              Verbosity mode.

Example:
  wprecon -u "https://xxxxxxxx.com" --detection-waf
  wprecon -u "https://xxxxxxxx.com" --aggressive-mode
  wprecon -u "https://xxxxxxxx.com" -A --scripts script1,script2,script3
`, green("[flags]"), green("[target]"), printer.Required, green("[dir]"), green("[seconds]"), green("[list]"))

// HelpFuzzer ::
var HelpFuzzer = fmt.Sprintf(`wprecon (Wordpress Recon) is a tool for wordpress exploration!

Usage:
  wprecon fuzzer %s

Flags:
      --backup-file            Performs a fuzzing to try to find the backup file if it exists.
  -U, --usernames %s       Set usernames attack passwords.
  -P, --passwords %s  Set wordlist attack passwords.
  -M, --method-attack %s Forces the use of a non-standard attack method (XML-RPC). Available methods: xml-rpc, wp-login

Global Flags:
  -h, --help                   help for wprecon.
  -u, --url %s           Target URL (Ex: http(s)://example.com/). %s
  -f, --force                  Forces wprecon to not check if the target is running WordPress and forces other executions.
      --random-agent           Use randomly selected HTTP(S) User-Agent header value.
      --tor                    Use Tor anonymity network.
      --disable-tls-checks     Disables SSL/TLS certificate verification.
      --http-sleep %s   You can make each request slower, if there is a WAF, it can make it difficult for it to block you. (default: 0)
      --scripts %s         Auxiliary scripts.
  -v, --verbose                Verbosity mode.

Example:
  wprecon fuzz -u "https://xxxxxxxx.com" -U user -P $HOME/wordlist/rockyou.txt
  wprecon fuzz -u "https://xxxxxxxx.com" -U user1,user2,user3 -P $HOME/wordlist/rockyou.txt
  wprecon fuzz -u "https://xxxxxxxx.com" --backup-file --random-agent
`, green("[flags]"), green("[list]"), green("[file-path]"), green("[attack]"), green("[target]"), printer.Required, green("[seconds]"), green("[list]"))
