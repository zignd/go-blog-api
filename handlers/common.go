package handlers

import (
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
)

func ConnMongoDB() (*mgo.Session, *mgo.Database, error) {
	// TODO: set up viper, and retrieve those hardcoded values from somewhere else
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to estabilish a connection")
	}
	db := session.DB("blog")
	return session, db, nil
}
