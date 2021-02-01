package config

var InfosWprecon informationsWPRecon

type informationsWPRecon struct {
	Target        string
	TotalRequests int

	Verbose bool
	Force   bool

	TimeStart, TimeFinish string
	TimeSleepRequests     int
	TimeOutRequests       int

	WPContent string

	OtherInformationsString    map[string]string
	OtherInformationsInt       map[string]int
	OtherInformationsSlice     map[string][]string
	OtherInformationsBool      map[string]bool
	OtherInformationsMapString map[string]map[string]string
}

func init() {
	InfosWprecon.OtherInformationsSlice = map[string][]string{}
	InfosWprecon.OtherInformationsString = map[string]string{}
	InfosWprecon.OtherInformationsInt = map[string]int{}
	InfosWprecon.OtherInformationsBool = map[string]bool{}
	InfosWprecon.OtherInformationsMapString = map[string]map[string]string{}

	InfosWprecon.OtherInformationsMapString["target.http.plugins.versions"] = map[string]string{}
	InfosWprecon.OtherInformationsMapString["target.http.themes.versions"] = map[string]string{}
}
