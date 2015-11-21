package db
import "gopkg.in/mgo.v2/bson"

type Agenda struct {
	Id   int
	Days []Day
}

type Day struct {
	Id int `bson:"_id"`
	Name string `bson:"name"`
	Tracks []Track `bson:"-"`
}

type Track struct {
	Id int  `bson:"_id"`
	Name string  `bson:"name"`
	DayId int  `bson:"dayId"`
	Slots []Slot `bson:"-"`
}

type Slot struct {
	Id int `bson:"_id"`
	Start string `bson:"start"`
	End string `bson:"end"`
	TrackId int `bson:"trackId"`
	DayId int `bson:"dayId"`
	Vote int `bson:"-"`
	Contents Content `bson:"content"`
}

type Content struct {
	Id int `bson:"-"`
	Type string `bson:"type"`
	Title string `bson:"title"`
	Description string `bson:"descrition"`
	Authors []Author `bson:"author"`

}

type Author struct {
	Id int `bson:"-"`
	Uuid string `bson:"uuid"`
	Name string `bson:"name"`
	Avatar string `bson:"avatar"`
	Description string `bson:"description"`
}

type User struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Login string `bson:"login"`
	Password string `bson:"password"`
}

type Vote struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	UserId bson.ObjectId `bson:"userId"`
	SlotId int `bson:"slotId"`
	Vote int `bson:"vote"`
}