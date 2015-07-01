package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	childrenCollection   = "CHILDREN"
)


func (m *MongoContext) insertChild(child *Child) (id string, err error) {
	mongoSession := m.session.Clone()
	defer mongoSession.Close()

	c := mongoSession.DB(Configuration.MongoDbName).C(childrenCollection)

	id = GetNewId()
	story.StoryId = id

	err = c.Insert(story)
	if mgo.IsDup(err) {
		// retry insert with new id
		m.insertStory(story)
	} else if err != nil {
		Error.Printf("Mongo insert to %s/%s returned '%s'", Configuration.MongoDbName,
			inkblotStoryCollection, err.Error())
	}
	return
}