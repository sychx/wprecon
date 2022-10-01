//go:build windows

package banner

const (
	BAR   = "——————————————————————————————————————————————————————————————————————"
	ASCII = "___       ______________________________________________   __\n__ |     / /__  __ \\__  __ \\__  ____/_  ____/_  __ \\__  | / /\n__ | /| / /__  /_/ /_  /_/ /_  __/  _  /    _  / / /_   |/ / \n__ |/ |/ / _  ____/_  _, _/_  /___  / /___  / /_/ /_  /|  /  \n____/|__/  /_/     /_/ |_| /_____/  \\____/  \\____/ /_/ |_/  "

	BANNER_HELP_ROOT = `WPRecon, is a tool for the recognition of vulnerabilities and blackbox information for wordpress.

	Usage:
	  wprecon [flags]
	
	Flags:
	  -h, --help                 help for wprecon.
	  -u, --url [target]         Target URL (Ex: http(s)://example.com/). (REQUIRED)
	  -f, --force                Forces wprecon to not check if the target is running WordPress and forces other executions.
	  -A, --aggressive-mode      Activates the aggressive mode of wprecon.
		  --detection-waf        I will try to detect if the target is using any WAF.
		  --random-agent         Use randomly selected HTTP(S) User-Agent header value.
		  --tor                  Use Tor anonymity network.
		  --wp-content-dir [dir] In case the wp-content directory is customized. (default: wp-content)
		  --http-sleep [seconds] You can make each request slower, if there is a WAF, it can make it difficult for it to block you. (default: 0)
		  --disable-tls-checks   Disables SSL/TLS certificate verification.
	  -v, --verbose              Verbosity mode.
	
	Example:
	  wprecon -u "https://xxxxxxxx.com" --detection-waf
	  wprecon -u "https://xxxxxxxx.com" --aggressive-mode
	`
	
)