package db

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"path/filepath"
//"gopkg.in/mgo.v2/bson"
	"os"
)

var Session *mgo.Session
var Days    *mgo.Collection
var Tracks  *mgo.Collection
var Slots   *mgo.Collection

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Init() {
	var err error
	Session, err = mgo.Dial(os.Getenv("MONGO_PORT_27017_TCP_ADDR"))
	check(err)
	Session.SetMode(mgo.Monotonic, true)
	Days = Session.DB("cmvoting").C("days")
	Tracks = Session.DB("cmvoting").C("tracks")
	Slots = Session.DB("cmvoting").C("slots")
}

func ImportReferentData() {

	absPath, _ := filepath.Abs("app/db/agenda.json")
	dat, err := ioutil.ReadFile(absPath)
	check(err)
	var parsedAgenda Agenda
	json.Unmarshal(dat, &parsedAgenda)

	Days.DropCollection()
	Tracks.DropCollection()
	Slots.DropCollection()

	for _, day := range parsedAgenda.Days {
		fmt.Println(" Inserting tracks day " + day.Name)
		Days.Insert(day)
		for _, track := range day.Tracks {
			track.DayId = day.Id
			Tracks.Insert(track)
			for _, slot := range track.Slots {
				Slots.Insert(slot)
			}
		}

	}
}