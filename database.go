package main

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//mongodb://newtpu:database1@ds239903.mlab.com:39903/paragliding

type DBInfo struct {
	ConnectionString string
	DBString         string
	CollectionString string
}

var Credentials DBInfo

func dbInit() {
	Credentials.CollectionString = "tracks"
	Credentials.DBString = "paragliding"
	Credentials.ConnectionString = "mongodb://newtpu:database1@ds239903.mlab.com:39903/paragliding"
}

//Inserts a track into the track collection
func insertTrack(t *Track) {
	session, err := mgo.Dial(Credentials.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	err = session.DB(Credentials.DBString).C(Credentials.CollectionString).Insert(t)
	if err != nil {
		fmt.Println(err)
	}

}

func countTrack() int {
	session, err := mgo.Dial(Credentials.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	count, err := session.DB(Credentials.DBString).C(Credentials.CollectionString).Count()
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

	err = session.DB(Credentials.DBString).C(Credentials.CollectionString).Find(bson.M{}).All(&all)
	return all
}

func loadFromDB() {
	//Load tracks from DB into memory

}

//Deletes everything in the database
func deleteTrackCollection() int {
	count := countTrack() //Get database count

	session, err := mgo.Dial(Credentials.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	err = session.DB(Credentials.DBString).C(Credentials.CollectionString).DropCollection()
	if err != nil {
		fmt.Println(err)
	}

	return count
}
