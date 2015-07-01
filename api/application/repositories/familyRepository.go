package repositories

import (
	"gopkg.in/mgo.v2"
	"fmt"

	"github.com/dicknaniel/Chillowance/api/application/models"
	"github.com/dicknaniel/Chillowance/api/application/util"
)

const (
	familiesCollection = "FAMILIES"
)

type MyMongoContext util.MongoContext

func (m *MyMongoContext) RepoCreateFamily(family *models.Family) (id string, err error) {
	mongoSession := m.Session.Clone()
	defer mongoSession.Close()

	c := mongoSession.DB(util.Configuration.MongoDbName).C(familiesCollection)

	id = util.GetNewId()
	family.FamilyId = id

	err = c.Insert(family)
	if mgo.IsDup(err) {
		// retry insert with new id
		m.RepoCreateFamily(family)
	} else if err != nil {
		fmt.Errorf("Mongo insert to %s/%s returned '%s'", util.Configuration.MongoDbName,
			familiesCollection, err.Error())
	}
	return
}
