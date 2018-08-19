package controllers

import (
	"strconv"

	"github.com/globalsign/mgo/bson"
	"github.com/revel/revel"
	"github.com/rwlist/rwcore/app/models/lists"
	"github.com/rwlist/rwcore/app/models/stree"
)

type Lists struct {
	AuthenticatedController
}

func (c Lists) Index() revel.Result {
	lists, err := lists.AvailableLists()
	if err != nil {
		return c.RenderMyError(err)
	}
	return c.RenderJSON(lists)
}

func (c Lists) Data(name string) revel.Result {
	data, err := lists.FullListInfo(name)
	if err != nil {
		return c.RenderMyError(err)
	}
	return c.RenderJSON(data)
}

func (c Lists) InsertOne(name string) revel.Result {
	var jsonData lists.Element
	err := c.Params.BindJSON(&jsonData)
	if err != nil {
		return c.RenderMyError(err)
	}
	insert, err := lists.InsertOne(name, jsonData)
	if err != nil {
		return c.RenderMyError(err)
	}
	return c.RenderJSON(insert)
}

func (c Lists) InsertMany(name string) revel.Result {
	var jsonData []lists.Element
	err := c.Params.BindJSON(&jsonData)
	if err != nil {
		return c.RenderMyError(err)
	}
	insert, err := lists.InsertMany(name, jsonData)
	if err != nil {
		return c.RenderMyError(err)
	}
	return c.RenderJSON(insert)
}

// TODO: Add confirmation
func (c Lists) Clear(name string) revel.Result {
	// TODO: return all data after clearing
	info, err := lists.Clear(name)
	if err != nil {
		return c.RenderMyError(err)
	}
	return c.RenderJSON(map[string]interface{}{
		"Description": "All cleared!",
		"Info":        info,
		"Status":      "OK",
	})
}

func (c Lists) Backup(name string) revel.Result {
	data, err := lists.FullListInfo(name)
	if err != nil {
		return c.RenderMyError(err)
	}
	return c.RenderJSON(data.Elements)
}

func (c Lists) CopyToDir(name, dirID string) revel.Result {
	session := stree.NewSession(bson.ObjectIdHex(c.userID))
	defer session.Close()

	dirObjectID := bson.ObjectIdHex(dirID)

	resp := struct {
		Error       interface{} `json:"Error,omitempty"`
		Status      string
		CountCopied int
	}{}

	info, err := lists.FullListInfo(name)

	if err != nil {
		resp.Error = err
		goto response
	}

	for i, v := range info.Elements {
		_, err := session.CreateFile(dirObjectID, strconv.Itoa(i), v)
		if err != nil {
			resp.Error = err
			goto response
		}
		resp.CountCopied++
	}

response:
	return c.RenderJSON(resp)
}
