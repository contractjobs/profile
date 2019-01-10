package persistence

type Profile struct {
	ID                string
	Profilename       string
	Email             string
	PhoneNo           string
	ContractEndDate   string
	Location          string
	Skills            []string
	PreferredRate     int
	PreferredLocation []string
	notifyRules       NotifyRules
}

type NotifyRules struct {
	NotifyWhenContractIsEnding       bool
	NotifyWhenRateMatch              bool
	NotifyWhenPreferredLocationMatch bool
}
