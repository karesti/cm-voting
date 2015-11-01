package controllers

import (
	"github.com/revel/revel"
	"github.com/karesti/cm-voting/app/db"
	"gopkg.in/mgo.v2/bson"
)

type App struct {
	*revel.Controller
}


func (c App) Index() revel.Result {
	db.Init()
	db.ImportReferentData()

	var track db.Track

	db.Tracks.Find(bson.M{}).One(&track)

	return c.Render(track)
}
