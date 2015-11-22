package controllers
import (
	"github.com/revel/revel"
	"github.com/karesti/cm-voting/app/db"
	"github.com/karesti/cm-voting/app/routes"
	"strconv"
)

type Voting struct {
	App
}

func (c Voting) checkConnected() revel.Result {
	if user := c.connected(); user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Index())
	}
	return nil
}

func (c Voting) List() revel.Result {
	c.checkConnected()
	if user := c.connected(); user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Index())
	}

	days := db.LoadDays()
	return c.Render(days)
}

func (c Voting) ListDay(dayId int) revel.Result {
	c.checkConnected()
	user := c.connected()
	tracks := db.LoadTracks(dayId, user)
	day := db.DayById(dayId)
	day.Tracks = tracks
	return c.Render(day)
}

func (c Voting) VoteSlot(slotId int) revel.Result {
	c.checkConnected()
	user := c.connected()
	var slot = db.Slot{}
	err := db.FindSlotById(slotId, &slot)
	if err != nil {
		panic(err);
	}

	var vote = db.Vote{}
	db.FindVoteBySlotAndUser(slotId, user.ID, &vote)
	c.Flash.Data["vote"] = strconv.Itoa(vote.Vote)
	return c.Render(slot)
}

func (c Voting) SendVote(vote int) revel.Result {
	c.checkConnected()
	slotId, err := strconv.Atoi(c.Params.Get("slotId"))
	if (err != nil) {
		panic(err)
	}
	var slot = db.Slot{}
	err = db.FindSlotById(slotId, &slot)
	if (err != nil) {
		panic(err)
	}

	user := c.connected()

	err = db.SaveVote(&db.Vote{UserId : user.ID, SlotId: slot.Id, Vote : vote})
	if err != nil {
		panic(err);
	}

	return c.Redirect(routes.Voting.ListDay(slot.DayId))
}

