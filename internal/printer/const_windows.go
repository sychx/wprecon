//go:build windows

package printer

const (
    PREFIX_DANGER  = "[-]"
    PREFIX_FATAL   = "[!]"
    PREFIX_DONE    = "[+]"
    PREFIX_WARNING = "[!]"
    PREFIX_SCAN    = "[?]"
    PREFIX_INFO    = "[i]"
    PREFIX_VERBOSE = "[v]"

    PREFIX_LIST_DONE    = "    —"
    PREFIX_LIST_DANGER  = "    —"
    PREFIX_LIST_DEFAULT = "    —"
    PREFIX_LIST_WARNING = "    —"

    REQUIRED = "(Required)"
    WARNING  = "(Warning)"
)
