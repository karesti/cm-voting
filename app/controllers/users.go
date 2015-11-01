package controllers
import "github.com/revel/revel"

type Users struct {
	*revel.Controller
}

func (c Users) Login() revel.Result {
	return c.Render()
}