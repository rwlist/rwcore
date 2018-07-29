package controllers

import (
	"math/rand"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting := "Aloha World"
	return c.Render(greeting)
}

func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("Your name is required!")
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(myName)
}

func (c App) UserPage() revel.Result {
	userID := c.AuthorizedUserID()
	return c.Render(userID)
}

func (c App) LoginPage() revel.Result {
	return c.Render()
}

func (c App) RegisterPage() revel.Result {
	return c.Render()
}

func (c App) AuthorizedUserID() string {
	if rand.Intn(100) < 30 {
		return "fake_user_id"
	}
	return ""
}
