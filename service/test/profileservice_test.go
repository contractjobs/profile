package main

import (
	"log"
	"testing"

	"github.com/contractjobs/profile/service"
	"github.com/stretchr/testify/assert"
)

func TestNewProfile_forSuccess(t *testing.T) {
	expected := "sanjayk@test2.com"
	theJSON := `{"Email":"sanjayk@test2.com","ProfileName":"Sanjay Kataria","Phone":0,"Company":"","Skills":["Java","MongoDB"],"CurrentContractEndDate":"01/31/2019","CurrentLocation":"Northern Virginia","PreferredRatePerHour":0,"PreferredLocation":["Virginia","Maryland","Washington DC"],"Notify":false,"NotifyMethod":{"Email":false,"Phone":false},"NotifyRules":{"NotifyOnRateMatch":false,"NotifyOnLocationMatch":false}}`
	response, err := service.NewProfileService().NewProfile(theJSON)
	assert.IsType(t, nil, err)
	assert.Equal(t, expected, response)
}

func TestFindProfileByEmail_forSuccess(t *testing.T) {
	email := "sanjayk@test1.com"
	expected := `{"Email":"sanjayk@test1.com","ProfileName":"Sanjay Kataria","Phone":0,"Company":"","Skills":["Java","MongoDB"],"CurrentContractEndDate":"01/31/2019","CurrentLocation":"Northern Virginia","PreferredRatePerHour":0,"PreferredLocation":["Virginia","Maryland","Washington DC"],"Notify":false,"NotifyMethod":{"Email":false,"Phone":false},"NotifyRules":{"NotifyOnRateMatch":false,"NotifyOnLocationMatch":false}}`
	response, err := service.NewProfileService().FindProfileByEmail(email)
	log.Println(response)
	assert.IsType(t, nil, err)
	assert.Equal(t, expected, response)
}

func TestFindAllProfiles_forSuccess(t *testing.T) {
	response, err := service.NewProfileService().FindAllProfiles()
	assert.IsType(t, nil, err)
	assert.NotNil(t, response)
}
