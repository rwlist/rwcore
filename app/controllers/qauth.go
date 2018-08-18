package controllers

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/revel/revel"
	"github.com/rwlist/rwcore/app/models/mongodb"
	"golang.org/x/crypto/bcrypt"
)

type QAuth struct {
	MongoController
}

type UserStatus struct {
	Status   string
	UserID   string
	Username string
}

func (c QAuth) CurrentUser() revel.Result {
	userID := c.Session["userID"]
	username := c.Session["username"]

	var status string
	if userID != "" {
		status = "User"
	} else {
		status = "Unknown"
	}

	return c.RenderJSON(
		UserStatus{
			Status:   status,
			UserID:   userID,
			Username: username,
		},
	)
}

func (c QAuth) DoLogin(username, password string, remember bool) revel.Result {
	users := c.MongoController.DB.C("users")

	var user mongodb.User
	err := users.Find(bson.M{"username": username}).One(&user)
	if err == mgo.ErrNotFound {
		goto failed
	}
	if err != nil {
		return c.RenderMyError(err)
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
			"Description": "Login successful",
			"Username":    username,
			"Status":      "OK",
		})
	}

failed:
	return c.RenderMyError("Login failed")
}

func (c QAuth) AuthorizeUser(user mongodb.User) {
	c.Session["username"] = user.Username
	c.Session["userID"] = user.ID.Hex()
}

func (c QAuth) DoRegister(username, password, verifyPassword string) revel.Result {
	c.Validation.Required(verifyPassword)
	c.Validation.Required(verifyPassword == password).
		MessageKey("Password does not match")

	if c.Validation.HasErrors() {
		return c.RenderMyError(
			struct {
				Message    string
				Validation *revel.Validation
			}{"Validation failed", c.Validation},
		)
	}

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return c.RenderMyError(err)
	}

	user := mongodb.User{
		ID:           bson.NewObjectId(),
		Username:     username,
		PasswordHash: passwordHash,
	}

	users := c.DB.C("users")

	err = users.Insert(&user) // TODO: avoid equals usernames
	if err != nil {
		return c.RenderMyError(err)
	}

	c.AuthorizeUser(user)
	return c.RenderJSON(map[string]string{
		"Description": "Registration successful",
		"Username":    username,
		"Status":      "OK",
	})
}

func (c QAuth) DoLogout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.RenderOK("Logout ok")
}
