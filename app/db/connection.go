package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/revel/revel"
	"gopkg.in/mgo.v2"
)

const DOMAIN = "marathon.mesos"

var session *mgo.Session

type Connection struct {
	session *mgo.Session
}

func (c *Connection) Close() {
	c.session.Close()
}

func Init() {
	var err error
	if revel.DevMode {
		session, err = mgo.Dial("192.168.99.100")
		check(err)
	} else {
		session, err = mgo.Dial(os.Getenv("MONGO_PORT_27017_TCP_ADDR"))

		if err != nil {
			// try mesos
			_, srv, err := net.LookupSRV(os.Getenv("SERVICE_NAME"), "tcp", DOMAIN)
			if err != nil {
				log.Fatal(err)
				check(err)
			} else {
				port := srv[0].Port
				url := fmt.Sprintf("%s.%s:%d", os.Getenv("SERVICE_NAME"), DOMAIN, port)
				fmt.Println("URL : " + url)
				session, err = mgo.Dial(url)
			}
		}

	}

	session.SetMode(mgo.Monotonic, true)
	importReferentData()
}

func CreateConnection() *Connection {
	return &Connection{session.Clone()}
}

func importReferentData() {
	conn := CreateConnection()
	defer conn.Close()

	absPath, _ := filepath.Abs("app/db/agenda.json")
	dat, err := ioutil.ReadFile(absPath)
	check(err)
	var parsedAgenda Agenda
	json.Unmarshal(dat, &parsedAgenda)

	conn.days().DropCollection()
	conn.tracks().DropCollection()
	conn.slots().DropCollection()

	for _, day := range parsedAgenda.Days {
		fmt.Println(" Inserting tracks day " + day.Name)
		conn.days().Insert(day)
		for _, track := range day.Tracks {
			track.DayId = day.Id
			conn.tracks().Insert(track)
			for _, slot := range track.Slots {
				if slot.Contents.Title != "" && slot.Contents.Type == "TALK" {
					slot.DayId = day.Id
					conn.slots().Insert(slot)
				}
			}
		}
	}
}

func (c *Connection) collection(name string) *mgo.Collection {
	return c.session.DB("cmvoting").C(name)
}

func (c *Connection) users() *mgo.Collection {
	return c.collection("users")
}

func (c *Connection) votes() *mgo.Collection {
	return c.collection("votes")
}

func (c *Connection) days() *mgo.Collection {
	return c.collection("days")
}

func (c *Connection) tracks() *mgo.Collection {
	return c.collection("tracks")
}

func (c *Connection) slots() *mgo.Collection {
	return c.collection("slots")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
