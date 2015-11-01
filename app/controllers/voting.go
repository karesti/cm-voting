package controllers
import (
	"github.com/revel/revel"
	"github.com/karesti/cm-voting/app/db"
)



type Voting struct {
	*revel.Controller
}


func (c Voting) List() revel.Result {
	slots := db.LoadSlots()
	return c.Render(slots)
}