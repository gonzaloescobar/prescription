package dao

import (
	"log"

	. "github.com/gonzaloescobar/prescriptions/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PrescriptionsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "prescriptions"
)

// Establish a connection to database
func (m *PrescriptionsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of prescriptions
func (m *PrescriptionsDAO) FindAll() ([]Prescription, error) {
	var prescriptions []Prescription
	err := db.C(COLLECTION).Find(bson.M{}).All(&prescriptions)
	return prescriptions, err
}

// Find a prescription by its id
func (m *PrescriptionsDAO) FindById(id string) (Prescription, error) {
	var prescription Prescription
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&prescription)
	return prescription, err
}

// Insert a prescription into database
func (m *PrescriptionsDAO) Insert(prescription Prescription) error {
	err := db.C(COLLECTION).Insert(&prescription)
	return err
}

// Delete an existing prescription
func (m *PrescriptionsDAO) Delete(prescription Prescription) error {
	err := db.C(COLLECTION).Remove(&prescription)
	return err
}

// Update an existing prescription
func (m *PrescriptionsDAO) Update(prescription Prescription) error {
	err := db.C(COLLECTION).UpdateId(prescription.ID, &prescription)
	return err
}
