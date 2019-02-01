package service

import (
	"encoding/json"

	"github.com/contractjobs/profile/lib/persistence"
	"github.com/contractjobs/profile/lib/persistence/dblayer"
)

type profileService struct {
	dbhandler persistence.DatabaseHandler
}

//NewProfileService get profile service
func NewProfileService() *profileService {
	databaseHandler, err := dblayer.NewPersistenceLayer("dynamodb", "us-east-1")
	if err != nil {
		return nil
	}
	return &profileService{
		dbhandler: databaseHandler,
	}
}

//FindAllProfiles
func (ps *profileService) FindAllProfiles() (string, error) {
	profile, err := ps.dbhandler.FindAllProfiles()
	if err != nil {
		return "", err
	}
	bytes, _ := json.Marshal(&profile)
	return string(bytes), nil
}

//FindProfileByEmail
func (ps *profileService) FindProfileByEmail(email string) (string, error) {
	profile, err := ps.dbhandler.FindProfileByEmail(email)
	if err != nil {
		return "", err
	}
	bytes, _ := json.Marshal(&profile)
	return string(bytes), nil
}

//NewProfile
func (ps *profileService) NewProfile(theJson string) (string, error) {

	var profile persistence.Profile
	err := json.Unmarshal([]byte(theJson), &profile)
	if err != nil {
		panic(err)
		return "", err
	}
	return ps.dbhandler.AddUpdateProfile(profile)
}
