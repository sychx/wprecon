package plugins

import (
	"regexp"

	"github.com/blackcrw/wprecon/internal/text"
	"github.com/blackcrw/wprecon/tools/finders"
)

func IndexSourceCodeBody(content, raw string) *[]finders.Finders {
	var regex = regexp.MustCompile(content + `/plugins/(.*?)/.*?[.css|.js]`)
	
	for _, submatch := range regex.FindAllStringSubmatch(raw, -1) {
		if !text.ContainsAny(*finders.Plugins, "Name", submatch[1]) {
			*finders.Plugins = append(*finders.Plugins, finders.Finders{
				Name: submatch[1],
				FoundBy: "In the HTML of the index - No Version",
				Others: []finders.FindersOthers{},
			})
		}
	}

	return finders.Plugins
}

func IndexSourceCodeBodyVersion(content, raw string) *[]finders.Finders {
	var regex = regexp.MustCompile(content + `/plugins/(.*?)/.*?[css|js].*?ver=(\d{1,2}\.\d{1,2}\.\d{1,3})`)

	for _, submatch := range regex.FindAllStringSubmatch(raw, -1) {
		if !text.ContainsAny(*finders.Plugins, "Name", submatch[1]) {
			*finders.Plugins = append(*finders.Plugins, finders.Finders{
				Name: submatch[1],
				FoundBy: "In the HTML of the index",
				Others: []finders.FindersOthers{{
					Version: submatch[2],
					FoundBy:"Version in the HTML code of the index",
					Match: append([]string{}, submatch[0]),
					Confidence: text.FormatConfidence(0, 15),
				}},
			})
		} else {
			var entity = (*finders.Plugins)[text.IndexAny(*finders.Plugins, "Name", submatch[1])]

			if !text.ContainsAny(entity.Others, "Version", submatch[2]) {
				entity.FoundBy = "In the HTML of the index"

				entity.Others = append(entity.Others, finders.FindersOthers{
					Version: submatch[2],
					FoundBy:  "Version in the HTML code of the index",
					Match: append([]string{}, submatch[0]),
					Confidence: text.FormatConfidence(0, 15),
				})
			}
		}
	}

	return finders.Plugins
}