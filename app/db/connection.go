package db

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"path/filepath"
	"os"
	"github.com/revel/revel"
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
	if(revel.DevMode){
		Session, err = mgo.Dial("192.168.99.100")
		check(err)
	} else {
		Session, err = mgo.Dial(os.Getenv("MONGO_PORT_27017_TCP_ADDR"))
		check(err)
	}

	Session.SetMode(mgo.Monotonic, true)
	db:=Session.DB("cmvoting")
	Days = db.C("days")
	Tracks = db.C("tracks")
	Slots = db.C("slots")
	importReferentData()
}

func importReferentData() {
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