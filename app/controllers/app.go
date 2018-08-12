package controllers

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/revel/revel"
	"github.com/rwlist/rwcore/app/models/mongodb"
	"golang.org/x/crypto/bcrypt"
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
	return c.Session["userID"]
}

func (c App) AuthorizeUser(user mongodb.User) {
	c.Session["username"] = user.Username
	c.Session["userID"] = user.ID.Hex()
	c.Flash.Success("Welcome, " + user.Username)
}

type UserStatus struct {
	Status   string
	UserID   string
	Username string
}

func (c App) CurrentUser() revel.Result {
	user := c.AuthorizedUserID()
	var status string
	if user != "" {
		status = "User"
	} else {
		status = "Unknown"
	}
	return c.RenderJSON(UserStatus{
		Status:   status,
		UserID:   c.Session["userID"],
		Username: c.Session["username"],
	})
}

func (c App) DoLogin(username, password string, remember bool) revel.Result {
	s := mongodb.NewCollectionSession("users")
	defer s.Close()

	var user mongodb.User
	err := s.Session.Find(bson.M{"username": username}).One(&user)
	if err == mgo.ErrNotFound {
		goto failed
	}
	if err != nil {
		return c.jsonError(err)
	}
	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
	if err == nil {
		c.AuthorizeUser(user)
		if remember {
			c.Session.SetDefaultExpiration()
		} else {
			c.Session.SetNoExpiration()
		}
		return c.RenderJSON(map[string]string{
			"Result":   "Login successful",
			"Username": username,
		})
	}

failed:
	return c.RenderJSON(map[string]string{
		"Err": "Login failed",
	})
}

func (c App) DoRegister(username, password, verifyPassword string) revel.Result {
	c.Validation.Required(verifyPassword)
	c.Validation.Required(verifyPassword == password).
		MessageKey("Password does not match")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.RegisterPage)
	}

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		panic(err)
	}

	user := mongodb.User{
		ID:           bson.NewObjectId(),
		Username:     username,
		PasswordHash: passwordHash,
	}

	s := mongodb.NewCollectionSession("users")
	defer s.Close()

	err = s.Session.Insert(&user) // TODO: avoid equals usernames
	if err != nil {
		panic(err)
	}

	c.AuthorizeUser(user)
	return c.Redirect(App.Index)
}

func (c App) DoLogout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(App.Index)
}

func (c App) jsonError(err error) revel.Result {
	return c.RenderJSON(
		map[string]interface{}{
			"Err": err,
		},
	)
}
