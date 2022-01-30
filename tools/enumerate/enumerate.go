package enumerate

import (
	"sync"

	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/text"
)

var (
	wg sync.WaitGroup
)

// In order not to have to keep repeating the same code several times and creating several "if" and "else", prefer to create this function and write the code only once.
// Its function is just to check if the version exists or if it doesn't.
// But if there is it will add +20 to the version confidence, otherwise it will add the version found in the version list.
func findings_add_version(name string, model *[]models.EnumerateModel, model_finding *models.FindingsVersionModel) {
	var model_matriz_others models.EnumerateOthersModel

	var _, int_name = text.ContainsEnumerateName(*model, name)

	if model_finding.Version != "" {
		if has_version, int_version := text.ContainsEnumerateVersion(*model, model_finding.Version); has_version {
			(*model)[int_name].Others[int_version].Confidence = format_confidence((*model)[int_name].Others[int_version].Confidence, 20)
		} else {
			model_matriz_others.Version = model_finding.Version
			model_matriz_others.FoundBy = "Findings Aggressive In Important Files"
			model_matriz_others.Match = append(model_matriz_others.Match, model_finding.Match)
			model_matriz_others.Confidence = format_confidence(model_matriz_others.Confidence, 20)
	
			(*model)[int_name].Others = append((*model)[int_name].Others, model_matriz_others)
		}
	}
}

func format_confidence(varx int8, value int8) int8 {
	if varx >= 100 { return 100 }

	return varx + value
}

func append_for(dst []models.EnumerateModel, src []models.EnumerateModel) (v []models.EnumerateModel) {
	for _, name := range src {
		v = append(dst, models.EnumerateModel{Name: name.Name, Slug: name.Slug})
	}

	return v
}
