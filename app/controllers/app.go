package controllers

import (
	"github.com/karesti/cm-voting/app/db"
	"github.com/karesti/cm-voting/app/routes"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller

	db *db.Connection
}

func (c *App) Index() revel.Result {
	if c.connected() != nil {
		return c.Redirect(routes.Voting.List())
	}
	return c.Render()
}

func (c *App) Before() revel.Result {
	c.db = db.CreateConnection()

	if user := c.connected(); user != nil {
		c.RenderArgs["user"] = user
	}
	return nil
}

func (c *App) After() revel.Result {
	if c.db != nil {
		c.db.Close()
	}
	return nil
}

func (c *App) connected() *db.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*db.User)
	}
	if username, ok := c.Session["user"]; ok {
		return c.getUser(username)
	}
	return nil
}

func (c *App) getUser(username string) *db.User {
	var user = db.User{}
	err := c.db.FindByLogin(username, &user)
	if err != nil {
		panic(err)
	}
	return &user
}

func (c *App) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.App.Index())
}
