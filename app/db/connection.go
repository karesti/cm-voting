package db

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"log"	
        "net"
	"io/ioutil"
	"encoding/json"
	"path/filepath"
	"os"
	"github.com/revel/revel"
)

const DOMAIN = "marathon.mesos"
var Session *mgo.Session
var Days    *mgo.Collection
var Users   *mgo.Collection
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
		_, srv, err := net.LookupSRV(os.Getenv("SERVICE_NAME"), "tcp", DOMAIN)
		if err != nil {
 			log.Fatal(err)
			check(err)
		} else {
			port := srv[0].Port
			url := fmt.Sprintf("%s.%s:%d", os.Getenv("SERVICE_NAME"), DOMAIN, port)
			fmt.Println("URL : " + url)
			Session, err = mgo.Dial(url)
		}
	}


	Session.SetMode(mgo.Monotonic, true)
	db:=Session.DB("cmvoting")
	Days = db.C("days")
	Tracks = db.C("tracks")
	Slots = db.C("slots")
	Users = db.C("users")
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
