package views

import (
	"fmt"

	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/printer"
)

func RootWAF(model *models.InterestingModel) {
	printer.Done("Web Application Firewall (WAF):", model.Name, "(Aggressive Detection)")
	printer.NewTopics("Location:", model.Url).Default()
	printer.NewTopics("Status Code:", fmt.Sprint(model.Status)).Default()

	if scan := printer.ScanQ("Do you wish to continue ?! [Y]es | [n]o : "); scan != "y" && scan != "\n" {
		printer.Fatal("Exiting...")
	}
}

func RootEnumerate(model models.EnumerateModel) {
	printer.Done("Name:", model.Name)
	printer.NewTopics("Found By:", model.FoundBy).Default()
	
	for _, model_other := range model.Others {
		if model_other.Version != "" {
			printer.NewTopics("Version:", model_other.Version).Default()
	
			if model_other.FoundBy != "" {
				printer.NewTopics("Found By:", model_other.FoundBy).Prefix("  ").Default()
			}
			
			if model_other.Confidence <= 30 {
				printer.NewTopics("Confiance:", fmt.Sprintf("%s%d%%", printer.RED, model_other.Confidence)).Prefix("  ").Danger()
			} else if model_other.Confidence < 70 {
				printer.NewTopics("Confiance:", fmt.Sprintf("%s%d%%", printer.YELLOW, model_other.Confidence)).Prefix("  ").Warning()
			} else if model_other.Confidence >= 70 {
				printer.NewTopics("Confiance:", fmt.Sprintf("%s%d%%", printer.GREEN, model_other.Confidence)).Prefix("  ").Done()
			} else {
				printer.NewTopics("Confiance:", fmt.Sprint(model_other.Confidence)).Prefix("  ").Default()
			}

			printer.NewTopics("Match:").Prefix("  ").Default()
			for _, match := range model_other.Match {
				printer.NewTopics(match).Prefix("    ").Default()
			}

		} else {
			printer.NewTopics("Version: No Version Match").Danger()
		}
	}

	printer.Println()
}