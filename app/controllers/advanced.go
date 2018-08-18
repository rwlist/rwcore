package controllers

import (
	"net/http"

	"github.com/rwlist/rwcore/app/models/mongodb"

	"github.com/globalsign/mgo"
	"github.com/revel/revel"
)

type AdvancedController struct {
	*revel.Controller
}

type AuthenticatedController struct {
	AdvancedController

	userID   string
	username string
}

type MongoController struct {
	AdvancedController

	DbSession *mgo.Session
	DB        *mgo.Database
}

func (c AdvancedController) RenderMyErrorStatus(err interface{}, status int) revel.Result {
	c.Response.Status = status
	c.Response.ContentType = "application/json"
	return c.RenderJSON(struct {
		Error  interface{}
		Status string
	}{err, "Error"})
}

func (c AdvancedController) RenderMyError(err interface{}) revel.Result {
	return c.RenderMyErrorStatus(err, http.StatusInternalServerError)
}

func (c AdvancedController) RenderOK(description interface{}) revel.Result {
	return c.RenderJSON(struct {
		Description interface{}
		Status      string
	}{description, "OK"})
}

func (c *AuthenticatedController) before() revel.Result {
	userID, ok1 := c.Session["userID"]
	username, ok2 := c.Session["username"]
	if !ok1 || !ok2 {
		return c.RenderMyErrorStatus("Unauthorized", http.StatusUnauthorized)
	}
	c.userID = userID
	c.username = username
	return nil
}

func (c *MongoController) before() revel.Result {
	c.DbSession = mongodb.BaseSession.Clone()
	c.DB = c.DbSession.DB("rwlist")
	return nil
}

func (c *MongoController) finally() revel.Result {
	c.DbSession.Close()
	return nil
}

func init() {
	revel.InterceptMethod((*AuthenticatedController).before, revel.BEFORE)

	revel.InterceptMethod((*MongoController).before, revel.BEFORE)
	revel.InterceptMethod((*MongoController).finally, revel.FINALLY)
}
