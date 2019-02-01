package persistence

//DatabaseHandler interface
type DatabaseHandler interface {
	AddUpdateProfile(Profile) (string, error)
	FindProfileByEmail(string) (Profile, error)
	FindAllProfiles() ([]Profile, error)
}
