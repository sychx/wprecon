package text

import (
	"regexp"
)

/*
   FindFilesByExtensions
     • This function searches for files by their extension, within an index of.

   This function searches for the version of the plugin or theme.
     • GetVersionByStableTag
     • GetVersionByChangelog
     • GetVersionByReleaseLog
*/

func FindImportantFiles(raw string) [][]string {
    var regex = regexp.MustCompile(`<a href=\"((?i)readme|license|changelog|release|release_log).(.*)\">.*</a>`)

    return regex.FindAllStringSubmatch(raw, -1)
}

func FindBackupFileOrPath(raw string) [][]string {
    var regex = regexp.MustCompile(`<a href=\"((?i)(back[up|wp|.*]|bkp).*)\">.*?</a>`)

    return regex.FindAllStringSubmatch(raw, -1)
}

func FindFilesByExtensions(raw string) [][]string {
    var regex = regexp.MustCompile(`<a href=\"((?i).*.zip|db|tar|tar.gz)\">.*?</a>`)

    return regex.FindAllStringSubmatch(raw, -1)
}

func GetVersionByStableTag(raw string) []string {
    var regex = regexp.MustCompile(`(?i)stable tag.*?([0-9.-]+)`)

    return regex.FindStringSubmatch(raw)
}

func GetVersionByChangelog(raw string) []string {
    var regex = regexp.MustCompile("=+\\s+(?:v(?:ersion)?\\s*)?([0-9.-]+)[ \ta-z0-9().\\-/]*=+")

    return regex.FindStringSubmatch(raw)
}

func GetVersionByReleaseLog(raw string) []string {
    regex := regexp.MustCompile(`(?i)version.*?([0-9.-]+)`)

    return regex.FindStringSubmatch(raw)
}