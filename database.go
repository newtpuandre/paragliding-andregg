package main

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/mongodb/mongo-go-driver/mongo"
)

//mongodb://newtpu:database1@ds239903.mlab.com:39903/paragliding

var client *mongo.Client

/*func dbInit() {
	var err error
	client, err = mongo.NewClient("mongodb://newtpu:database1@ds239903.mlab.com:39903/paragliding")
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}*/

//Inserts a track into the track collection
func insertTrack(t *Track) {
	session, err := mgo.Dial("mongodb://newtpu:database1@ds239903.mlab.com:39903/paragliding")
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	err = session.DB("paragliding").C("tracks").Insert(t)
	if err != nil {
		fmt.Println(err)
	}
}

func countTrack() int {
	session, err := mgo.Dial("mongodb://newtpu:database1@ds239903.mlab.com:39903/paragliding")
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	count, err := session.DB("paragliding").C("tracks").Count()
	if err != nil {
		fmt.Println(err)
		return -1
	}

	return count
}

//Deletes everything in the database
func deleteTrackCollection() {
	session, err := mgo.Dial("mongodb://newtpu:database1@ds239903.mlab.com:39903/paragliding")
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	err = session.DB("paragliding").C("tracks").DropCollection()
	if err != nil {
		fmt.Println(err)
	}
}
