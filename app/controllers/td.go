package controllers

import (
	"github.com/globalsign/mgo/bson"
	"github.com/revel/revel"
	"github.com/rwlist/rwcore/app/models/treedir"
)

type TD struct {
	*revel.Controller

	session treedir.Session
}

func (c *TD) before() revel.Result {
	userID, ok := c.Session["userID"]
	if !ok {
		c.Flash.Error("Use need login to access this page")
		return c.Redirect("/login")
	}
	c.session = treedir.NewSession(bson.ObjectIdHex(userID))
	return nil
}

func (c *TD) finally() revel.Result {
	c.session.Close()
	return nil
}

func (c TD) GetRoot() revel.Result {
	root, err := c.session.GetRoot()
	if err != nil {
		return c.jsonError(err)
	}
	node, err := c.session.GetNode(root.RootID)
	if err != nil {
		return c.jsonError(err)
	}
	return c.RenderJSON(node)
}

func (c TD) CreateDir(parentID string) revel.Result {
	var name struct {
		Name string
	}
	err := c.Params.BindJSON(&name)
	if err != nil {
		return c.jsonError(err)
	}
	node, err := c.session.CreateDir(bson.ObjectIdHex(parentID), name.Name)
	if err != nil {
		return c.jsonError(err)
	}
	return c.RenderJSON(node)
}

func (c TD) CreateFile(parentID string) revel.Result {
	name := "TODO"
	var content interface{}
	err := c.Params.BindJSON(&content)
	if err != nil {
		return c.jsonError(err)
	}
	file, err := c.session.CreateFile(bson.ObjectIdHex(parentID), name, content)
	if err != nil {
		return c.jsonError(err)
	}
	return c.RenderJSON(file)
}

func (c TD) ListDirectory(directoryID string) revel.Result {
	list, err := c.session.ListDirectory(bson.ObjectIdHex(directoryID))
	if err != nil {
		return c.jsonError(err)
	}
	return c.RenderJSON(list)
}

func (c TD) Delete(nodeID string) revel.Result {
	err := c.session.Delete(bson.ObjectIdHex(nodeID))
	if err != nil {
		return c.jsonError(err)
	}
	return c.RenderJSON(map[string]string{
		"Status": "Ok",
		"Msg":    "Succefully deleted",
	})
}

func (c TD) jsonError(err error) revel.Result {
	return c.RenderJSON(
		map[string]interface{}{
			"Err": err,
		},
	)
}

func init() {
	revel.InterceptMethod((*TD).before, revel.BEFORE)
	revel.InterceptMethod((*TD).finally, revel.FINALLY)
}
