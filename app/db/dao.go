package db
import (
	"gopkg.in/mgo.v2/bson"
)

func LoadTracks() []Track{
	var results []Track

	err := Tracks.Find(bson.M{}).All(&results)

	if err != nil {
		panic(err);
	}

	return results
}

func LoadSlots() []Slot{
	var results []Slot

	err := Slots.Find(bson.M{}).All(&results)

	if err != nil {
		panic(err);
	}

	return results
}