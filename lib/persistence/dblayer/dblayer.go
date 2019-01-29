package dblayer

import (
	"github.com/contractjobs/profile/lib/persistence"
	"github.com/contractjobs/profile/lib/persistence/dynamolayer"
)

type DBTYPE string

const (
	DYNAMODB DBTYPE = "dynamodb"
)

func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {

	switch options {
	case DYNAMODB:
		return dynamolayer.NewDynamoDBLayerByRegion(connection)
	}
	return nil, nil
}
