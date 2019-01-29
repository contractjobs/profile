package persistence

// Profile Model details
type Profile struct {
	Email                  string
	ProfileName            string
	Phone                  int
	Company                string
	Skills                 []string
	CurrentContractEndDate string
	CurrentLocation        string
	PreferredRatePerHour   int
	PreferredLocation      []string
	Notify                 bool
	NotifyMethod           NotifyMethod
	NotifyRules            NotifyRules
}

// NotifyRules used to determine if match should be notified to profile
type NotifyRules struct {
	NotifyOnRateMatch     bool
	NotifyOnLocationMatch bool
}

// NotifyMethod to notify profile
type NotifyMethod struct {
	Email bool
	Phone bool
}
