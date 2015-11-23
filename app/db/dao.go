package db

import (
	"gopkg.in/mgo.v2/bson"
)

func (c *Connection) LoadDays() []Day {
	var results []Day

	err := c.days().Find(bson.M{}).All(&results)

	if err != nil {
		panic(err)
	}

	return results
}

func (c *Connection) DayById(dayId int) Day {
	var day Day

	err := c.days().Find(bson.M{"_id": dayId}).One(&day)

	if err != nil {
		panic(err)
	}

	return day
}

func (c *Connection) LoadTracks(dayId int, user *User) []Track {
	var results []Track

	err := c.tracks().Find(bson.M{"dayId": dayId}).All(&results)

	if err != nil {
		panic(err)
	}

	for i, track := range results {
		track.Slots = c.loadSlots(track.Id, user)
		results[i] = track
	}

	return results
}

func (c *Connection) loadSlots(trackId int, user *User) []Slot {
	var results []Slot

	err := c.slots().Find(bson.M{"trackId": trackId}).All(&results)

	for i, slot := range results {
		vote := Vote{Vote: 0}
		c.FindVoteBySlotAndUser(slot.Id, user.ID, &vote)
		results[i].Vote = vote.Vote
	}

	if err != nil {
		panic(err)
	}

	return results
}

func (c *Connection) SaveVote(vote *Vote) error {
	_, err := c.votes().Upsert(bson.M{"slotId": vote.SlotId, "userId": vote.UserId}, vote)
	return err
}

func (c *Connection) FindByLogin(login string, userResult *User) error {
	return c.users().Find(bson.M{"login": login}).One(userResult)
}

func (c *Connection) FindSlotById(slotId int, slotResult *Slot) error {
	return c.slots().Find(bson.M{"_id": slotId}).One(slotResult)
}

func (c *Connection) FindVoteBySlotAndUser(slotId int, userId bson.ObjectId, voteResult *Vote) error {
	return c.votes().Find(bson.M{"slotId": slotId, "userId": userId}).One(voteResult)
}

func (c *Connection) CreateUser(login, password string) (User, error) {
	var result = User{}

	c.users().Insert(&User{Login: login, Password: password})
	err := c.users().Find(bson.M{"login": login}).One(&result)

	return result, err
}
