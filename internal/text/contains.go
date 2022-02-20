package text

import (
	"github.com/blackcrw/wprecon/internal/models"
)

func ContainsEnumerateSlug(enum []models.EnumerateModel, slug string) (bool, int) {
	for i, item := range enum {
		if item.Slug == slug {
			return true, i
		}
	}

	return false, 0
}

func ContainsEnumerateName(enum []models.EnumerateModel, name string) (bool, int) {
	for i, item := range enum {
		if item.Name == name {
			return true, i
		}
	}

	return false, 0
}

func ContainsEnumerateMatch(enum []models.EnumerateModel, match string) (bool, int) {
	for _, item := range enum {	
		for i, others := range item.Others {
			for _, matchs := range others.Match {
				if matchs == match {
					return true, i 
				}
			}
		}
	}

	return false, 0
}

func ContainsEnumerateVersion(enum []models.EnumerateModel, version string) (bool, int) {
	for _, item := range enum {	
		for i, others := range item.Others {
			if others.Version == version {
				return true, i 
			}
		}
	}

	return false, 0
}