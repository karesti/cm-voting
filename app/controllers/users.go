package controllers
import (
	"github.com/revel/revel"
	"github.com/karesti/cm-voting/app/routes"
	"github.com/karesti/cm-voting/app/db"
	"fmt"
)

type Users struct {
	*revel.Controller
}

func (c Users) Login(login, password string) revel.Result {
	c.Validation.Required(login)
	c.Validation.Required(password)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Index())
	}

	user := db.User{}

	err := db.FindByLogin(login, &user)

	fmt.Print(user)
	fmt.Print(err)
	if (err != nil) {
		c.Flash.Error("User does not exist")
		return c.Redirect(routes.App.Index())
	}

	if(user.Password != password){
		c.Flash.Error("User password does not match")
		return c.Redirect(routes.App.Index())
	}

	c.Session["user"] = login
	c.Flash.Success("Welcome, " + login)

	return c.Redirect(routes.Voting.List())
}

func (c Users) Signup() revel.Result {
	return c.Render()
}

func (c Users) SaveUser(login, password string) revel.Result {

	c.Validation.Required(login)
	c.Validation.Required(password)


	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Users.Signup())
	}

	user := db.User{}

	err := db.FindByLogin(login, &user)

	if (err != nil) {
		db.CreateUser(login, password)
		c.Session["user"] = login
		c.Flash.Success("Welcome, " + login)
		return c.Redirect(routes.Voting.List())
	}

	c.Flash.Error("User already exists")
	return c.Redirect(routes.Users.Signup())
}