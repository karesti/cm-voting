package db
import (
	"gopkg.in/mgo.v2/bson"
)

func LoadTracks() []Track {
	var results []Track

	err := Tracks.Find(bson.M{}).All(&results)

	if err != nil {
		panic(err);
	}

	return results
}

func LoadSlots() []Slot {
	var results []Slot

	err := Slots.Find(bson.M{}).All(&results)

	if err != nil {
		panic(err);
	}

	return results
}

func FindByLogin(login string, userResult *User) (error){

	return Users.Find(bson.M{"login": login}).One(userResult)
}

func CreateUser(login, password string) (User, error) {

	var result = User{}

	Users.Insert(&User{Login : login, Password: password})
	err := Users.Find(bson.M{"login": login}).One(&result)

	return result, err;
}