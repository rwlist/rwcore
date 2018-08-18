package controllers

import (
	"github.com/globalsign/mgo/bson"
	"github.com/revel/revel"
	"github.com/rwlist/rwcore/app/models/stree"
)

type STree struct {
	AuthenticatedController
	session stree.Session
}

func (c *STree) before() revel.Result {
	c.session = stree.NewSession(bson.ObjectIdHex(c.userID))
	return nil
}

func (c *STree) finally() revel.Result {
	c.session.Close()
	return nil
}

func (c STree) GetRoot() revel.Result {
	root, err := c.session.GetRoot()
	if err != nil {
		return c.RenderMyError(err)
	}
	node, err := c.session.GetNode(root.RootID)
	if err != nil {
		return c.RenderMyError(err)
	}
	return c.RenderJSON(node)
}

func (c STree) CreateDir(parentID string) revel.Result {
	var name struct {
		Name string
	}
	err := c.Params.BindJSON(&name)
	if err != nil {
		return c.RenderMyError(err)
	}
	node, err := c.session.CreateDir(bson.ObjectIdHex(parentID), name.Name)
	if err != nil {
		return c.RenderMyError(err)
	}
	return c.RenderJSON(node)
}

func (c STree) CreateFile(parentID string) revel.Result {
	name := "TODO"
	var content interface{}
	err := c.Params.BindJSON(&content)
	if err != nil {
		return c.RenderMyError(err)
	}
	file, err := c.session.CreateFile(bson.ObjectIdHex(parentID), name, content)
	if err != nil {
		return c.RenderMyError(err)
	}
	return c.RenderJSON(file)
}

func (c STree) ListDirectory(directoryID string) revel.Result {
	list, err := c.session.ListDirectory(bson.ObjectIdHex(directoryID))
	if err != nil {
		return c.RenderMyError(err)
	}
	return c.RenderJSON(list)
}

func (c STree) Delete(nodeID string) revel.Result {
	err := c.session.Delete(bson.ObjectIdHex(nodeID))
	if err != nil {
		return c.RenderMyError(err)
	}
	return c.RenderOK("Successfully deleted")
}

func init() {
	revel.InterceptMethod((*STree).before, revel.BEFORE)
	revel.InterceptMethod((*STree).finally, revel.FINALLY)
}
