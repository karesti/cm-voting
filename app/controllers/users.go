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

func (c Users) Login() revel.Result {
	return c.Render()
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

	user, err := db.CreateUser(login, password)

	fmt.Errorf(user)
	
	if(err != nil){
		c.Validation.Keep()
		c.FlashParams()

		return c.Redirect(routes.Users.Signup())
	}



	c.Session["user"] = user.Login
	c.Flash.Success("Welcome, " + user.Login)
	return c.Redirect(routes.Voting.List())
}