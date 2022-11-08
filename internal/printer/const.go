//go:build !windows

package printer

const (
	RED       = "\u001b[31;1m"
	BLUE      = "\u001b[34;1m"
	GREEN     = "\u001b[32;1m"
	BLACK     = "\u001b[30;1m"
	WHITE     = "\u001b[37;1m"
	YELLOW    = "\u001b[33;1m"
	MAGENTA   = "\u001b[35;1m"
	CYAN      = "\u001b[36;1m"
	RESET     = "\u001b[0m"
	BOLD      = "\u001b[1m"
	UNDERLINE = "\u001b[4m"
	REVERSED  = "\u001b[7m"

	PREFIX_DONE    = GREEN   + "[+]" +RESET
	PREFIX_DANGER  = RED     + "[-]" +RESET
	PREFIX_FATAL   = RED     + "[!]" +RESET
	PREFIX_INFO    = MAGENTA + "[i]" +RESET
	PREFIX_WARNING = YELLOW  + "[!]" +RESET
	PREFIX_SCAN    = YELLOW  + "[?]" +RESET
	PREFIX_VERBOSE = BLACK    + "[v]" +RESET

	PREFIX_LIST_DONE    = GREEN   + "    —" +RESET
	PREFIX_LIST_DANGER  = RED     + "    —" +RESET
	PREFIX_LIST_DEFAULT = WHITE   + "    —" +RESET
	PREFIX_LIST_WARNING = YELLOW  + "    —" +RESET

	REQUIRED = RED    + "(Required)" +RESET
	WARNING  = YELLOW + "(Warning)"  +RESET
)
