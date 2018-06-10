package datastore

import (
	"github.com/globalsign/mgo"
	"log"
)

type PeopleDAO struct {
	DBURL string
	DBNAME string
}

var db *mgo.Database

const (
	COLLECTION = "people"
)

func (p *PeopleDAO) Connect() {
	session, err := mgo.Dial(p.DBURL)

	if err != nil {
		log.Fatal("Error connecting to MongoDB", err)
	}

	db = session.DB(p.DBNAME)
}