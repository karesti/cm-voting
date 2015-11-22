package db
import (
	"gopkg.in/mgo.v2/bson"
)


func LoadDays() []Day {
	var results []Day

	err := Days.Find(bson.M{}).All(&results)

	if err != nil {
		panic(err);
	}

	return results
}

func DayById(dayId int) Day {
	var day Day

	err := Days.Find(bson.M{"_id" : dayId}).One(&day)

	if err != nil {
		panic(err);
	}

	return day
}


func LoadTracks(dayId int, user *User) []Track {
	var results []Track

	err := Tracks.Find(bson.M{"dayId" : dayId}).All(&results)

	if err != nil {
		panic(err);
	}

	for i, track := range results {
		track.Slots = loadSlots(track.Id, user)
		results[i] = track
	}

	return results
}

func loadSlots(trackId int, user *User) []Slot {
	var results []Slot

	err := Slots.Find(bson.M{"trackId" : trackId}).All(&results)

	for i, slot := range results {
		vote := Vote{Vote:0}
		FindVoteBySlotAndUser(slot.Id, user.ID, &vote)
		results[i].Vote = vote.Vote
	}

	if err != nil {
		panic(err);
	}

	return results
}


func SaveVote(vote *Vote) (error) {
	_, err := Votes.Upsert(bson.M{"slotId" : vote.SlotId, "userId" : vote.UserId}, vote)
	return err
}


func FindByLogin(login string, userResult *User) (error){
	return Users.Find(bson.M{"login": login}).One(userResult)
}
func FindSlotById(slotId int, slotResult *Slot) (error){
	return Slots.Find(bson.M{"_id": slotId}).One(slotResult)
}

func FindVoteBySlotAndUser(slotId int, userId bson.ObjectId, voteResult *Vote) (error){
	return Votes.Find(bson.M{"slotId": slotId, "userId" : userId}).One(voteResult)
}

func CreateUser(login, password string) (User, error) {

	var result = User{}

	Users.Insert(&User{Login : login, Password: password})
	err := Users.Find(bson.M{"login": login}).One(&result)

	return result, err;
}