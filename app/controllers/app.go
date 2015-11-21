package controllers

import (
	"github.com/revel/revel"
	"github.com/karesti/cm-voting/app/routes"
	"github.com/karesti/cm-voting/app/db"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	if (c.connected() != nil) {
		return c.Redirect(routes.Voting.List())
	}
	return c.Render();
}
func (c App) AddUser() revel.Result {
	if user := c.connected(); user != nil {
		c.RenderArgs["user"] = user
	}
	return nil
}

func (c App) connected() *db.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*db.User)
	}
	if username, ok := c.Session["user"]; ok {
		return c.getUser(username)
	}
	return nil
}

func (c App) getUser(username string) *db.User {
	var user = db.User{}
	err := db.FindByLogin(username, &user)
	if err != nil {
		panic(err)
	}
	return &user
}

func (c App) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.App.Index())
}
