package text

import (
	"github.com/blackcrw/wprecon/internal/models"
)

func ContainsFindersSlug(models_finders []models.FindersModel, slug string) (bool, int) {
	for i, item := range models_finders {
		if item.Slug == slug {
			return true, i
		}
	}

	return false, 0
}

func ContainsFindersName(models_finders []models.FindersModel, name string) (bool, int) {
	for i, item := range models_finders {
		if item.Name == name {
			return true, i
		}
	}

	return false, 0
}

func ContainsFindersMatch(models_finders []models.FindersModel, match string) (bool, int) {
	for _, item := range models_finders {	
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

func ContainsFindersVersion(models_finders []models.FindersModel, version string) (bool, int) {
	for _, item := range models_finders {	
		for i, others := range item.Others {
			if others.Version == version {
				return true, i 
			}
		}
	}

	return false, 0
}