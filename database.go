package main

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//mongodb://newtpu:database1@ds239903.mlab.com:39903/paragliding

type DBInfo struct {
	ConnectionString        string
	DBString                string
	TrackCollectionString   string
	WebhookCollectionString string
}

var Credentials DBInfo

func dbInit() {
	Credentials.TrackCollectionString = "tracks"
	Credentials.WebhookCollectionString = "webhooks"
	Credentials.DBString = "paragliding"
	Credentials.ConnectionString = "mongodb://newtpu:database1@ds239903.mlab.com:39903/paragliding"
}

//Inserts a track into the track collection
func insertTrack(t *Track, db *DBInfo) {
	session, err := mgo.Dial(db.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	err = session.DB(db.DBString).C(db.TrackCollectionString).Insert(t)
	if err != nil {
		fmt.Println(err)
	}

	//Add ID to array for used ids
	trackID = append(trackID, lastID)

	//Remember to count up used ids
	lastID++

}

func countTrack(db *DBInfo) int {
	session, err := mgo.Dial(db.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	count, err := session.DB(db.DBString).C(db.TrackCollectionString).Count()
	if err != nil {
		fmt.Println(err)
		return -1
	}

	return count
}

func getAllTracks() []Track {
	session, err := mgo.Dial(Credentials.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	var all []Track

	err = session.DB(Credentials.DBString).C(Credentials.TrackCollectionString).Find(bson.M{}).All(&all)
	return all
}

func updateIdFromDB() {
	count := countTrack(&Credentials)
	for i := 0; i < count; i++ {
		trackID = append(trackID, lastID)
		lastID++
	}

	hooks := getWebHooks()
	for i := 0; i < len(hooks); i++ {
		webhookID = append(webhookID, hooks[i].WebhookID)
	}

	if len(hooks) > 0 {
		lastWebhookID = webhookID[len(webhookID)-1] + 1
	} else {
		lastWebhookID = 0
	}

}

//Deletes everything in the database
func deleteTrackCollection() int {
	count := countTrack(&Credentials) //Get database count

	session, err := mgo.Dial(Credentials.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	err = session.DB(Credentials.DBString).C(Credentials.TrackCollectionString).DropCollection()
	if err != nil {
		fmt.Println(err)
	}

	return count
}

func getWebHooks() []webhookStruct {
	session, err := mgo.Dial(Credentials.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	var all []webhookStruct

	err = session.DB(Credentials.DBString).C(Credentials.WebhookCollectionString).Find(bson.M{}).All(&all)
	return all
}

func countWebhook() int {
	session, err := mgo.Dial(Credentials.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	count, err := session.DB(Credentials.DBString).C(Credentials.WebhookCollectionString).Count()
	if err != nil {
		fmt.Println(err)
		return -1
	}

	return count
}

func insertWebhook(w *webhookStruct) {
	session, err := mgo.Dial(Credentials.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	err = session.DB(Credentials.DBString).C(Credentials.WebhookCollectionString).Insert(w)
	if err != nil {
		fmt.Println(err)
	}

}

func updateWebhook(w *webhookStruct) {
	session, err := mgo.Dial(Credentials.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	err = session.DB(Credentials.DBString).C(Credentials.WebhookCollectionString).Update(bson.M{"_id": w.ID}, w)
}

func deleteWebhook(w *webhookStruct) {
	session, err := mgo.Dial(Credentials.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	err = session.DB(Credentials.DBString).C(Credentials.WebhookCollectionString).Remove(bson.M{"_id": w.ID})
}
