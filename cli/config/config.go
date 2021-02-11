package config

var Database informationsWPRecon

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
	Database.OtherInformationsSlice = map[string][]string{}
	Database.OtherInformationsString = map[string]string{}
	Database.OtherInformationsInt = map[string]int{}
	Database.OtherInformationsBool = map[string]bool{}
	Database.OtherInformationsMapString = map[string]map[string]string{}

	Database.OtherInformationsMapString["target.http.plugins.versions"] = map[string]string{}
	Database.OtherInformationsMapString["target.http.themes.versions"] = map[string]string{}
}
