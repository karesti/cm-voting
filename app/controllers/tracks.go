package controllers
import (
	"github.com/revel/revel"
	"github.com/karesti/cm-voting/app/db"
)



type Tracks struct {
	*revel.Controller
}


func (c Tracks) List() revel.Result {
	tracks := db.LoadTracks()
	return c.Render(tracks)
}