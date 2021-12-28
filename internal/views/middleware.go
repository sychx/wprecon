package views

import (
	"fmt"

	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/printer"
)

func MiddlewareWAFActive(model *models.MiddlewareFirewallModel) {
	printer.Danger("Firewall Active Detection:", model.Name)
	printer.NewTopics("Found By:", model.FoundBy).Default()
	printer.NewTopics("Confidence:", fmt.Sprintf("%d%%", model.Confidence)).Default()

	if model.Solve != "" {
		printer.NewTopics("Solve:", model.Solve)
	}
	
	printer.Println()

	if response := printer.ScanQ("Identified firewall, do you want to continue ? [y]es, [N]o : "); response != "y" && response != "Y" {
		printer.Fatal("Exiting...")
	}

	printer.Println()
}