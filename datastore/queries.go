package datastore

import "github.com/globalsign/mgo/bson"

func (p *PeopleDAO) FindAll() ([]Person, error) {
	var people []Person

	err := db.C(COLLECTION).Find(bson.M{}).All(&people)
	return people, err
}

func (p *PeopleDAO) FindById(id string) (Person, error) {
	var person Person
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&person)
	return person, err
}
func (p *PeopleDAO) Insert(person Person) error {
	err := db.C(COLLECTION).Insert(&person)
	return err
}

func (p *PeopleDAO) Delete(person Person) error {
	err := db.C(COLLECTION).Remove(&person)
	return err
}

func (p *PeopleDAO) Update(person Person) error {
	err := db.C(COLLECTION).UpdateId(person.ID, &person)
	return err
}