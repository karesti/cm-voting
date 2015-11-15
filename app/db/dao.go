package db
import (
	"gopkg.in/mgo.v2/bson"
	"errors"
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

func CreateUser(login, password string) (User, error) {

	var result = User{}

	err := Users.Find(bson.M{"Login": login}).One(&result)

	if(result.Login == login){
		return User{}, errors.New("Login already exists")
	}

	err = Users.Insert(&User{Login : login, Password: password})

	if err != nil {
		panic(err)
	}

	err = Users.Find(bson.M{"Login": login}).One(&result)

	return result, nil;
}