package views

import (
	"fmt"

	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/printer"
)

func RootWAF(models_finders *models.InterestingModel) {
	printer.Done("Web Application Firewall (WAF):", models_finders.Name, "(Aggressive Detection)")
	printer.NewTopics("Location:", models_finders.Url).Default()
	printer.NewTopics("Status Code:", fmt.Sprint(models_finders.Status)).Default()

	if scan := printer.ScanQ("Do you wish to continue ?! [Y]es | [n]o : "); scan != "y" && scan != "\n" {
		printer.Fatal("Exiting...")
	}
}

func RootFindersUser(models_finders models.FindersModel) {
	if models_finders.Name != "" && models_finders.Slug != "" {
		printer.Done("Name:", models_finders.Name, "Slug:", models_finders.Slug)
	} else if models_finders.Name != "" && models_finders.Slug == "" { printer.Done("Name:", models_finders.Name)
	} else if models_finders.Name == "" && models_finders.Slug != "" { printer.Done("Slug:", models_finders.Slug) }

	if models_finders.FoundBy != "" { printer.NewTopics("Found By:", models_finders.FoundBy).Default() }
}

func RootFindersPluginsAndThemes(models_finders models.FindersModel) {
	printer.Done("Name:", models_finders.Name)
	printer.NewTopics("Found By:", models_finders.FoundBy).Default()
	
	for _, model_other := range models_finders.Others {
		if model_other.Version != "" {
			printer.NewTopics("Version:", model_other.Version).Default()
		}

		if model_other.FoundBy != "" {
			printer.NewTopics("Found By:", model_other.FoundBy).Prefix("  ").Default()
		}

		if model_other.Confidence != 0 { 
			if model_other.Confidence <= 30 {
				printer.NewTopics("Confiance:", fmt.Sprintf("%s%d%%", printer.RED, model_other.Confidence)).Prefix("  ").Danger()
			} else if model_other.Confidence < 70 {
				printer.NewTopics("Confiance:", fmt.Sprintf("%s%d%%", printer.YELLOW, model_other.Confidence)).Prefix("  ").Warning()
			} else if model_other.Confidence >= 70 {
				printer.NewTopics("Confiance:", fmt.Sprintf("%s%d%%", printer.GREEN, model_other.Confidence)).Prefix("  ").Done()
			} else {
				printer.NewTopics("Confiance:", fmt.Sprint(model_other.Confidence)).Prefix("  ").Default()
			}
		}

		if len(model_other.Match) != 0 {
			printer.NewTopics("Match:").Prefix("  ").Default()
			for _, match := range model_other.Match {
				printer.NewTopics(match).Prefix("    ").Default()
			}
		}
	}

	printer.Println()
}