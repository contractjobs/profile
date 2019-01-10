package persistence

type DatabaseHandler interface {
	AddUpdateProfile(Profile) ([]byte, error)
	FindProfileByEmail(string) (Profile, error)
}
