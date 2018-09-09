package lists

import (
	"time"

	"github.com/globalsign/mgo"

	"github.com/globalsign/mgo/bson"
)

type Info struct {
	Name string
	Size int
}

type Element map[string]interface{}

type Data struct {
	Info
	Elements []Element
}

type InsertResult struct {
	Inserted []Element
}

func GetInfo(db Database, name string) (Info, error) {
	c := db.C(name)
	size, err := c.Count()
	if err != nil {
		return Info{}, nil
	}
	return Info{name, size}, nil
}

func AvailableLists() (lists []Info, err error) {
	db := service.DB()
	defer db.Close()

	names, err := db.CollectionNames()
	if err != nil {
		return
	}

	for _, name := range names {
		var info Info
		info, err = GetInfo(db, name)
		if err != nil {
			return
		}
		lists = append(lists, info)
	}
	return
}

// FullListInfo fetches all list data, including content
// of all elements. Creates empty list, if list doesn't
// exist
func FullListInfo(name string) (*Data, error) {
	db := service.DB()
	defer db.Close()
	c := db.C(name)
	info, err := GetInfo(db, name)
	if err != nil {
		return nil, err
	}
	data := &Data{info, []Element{}}
	err = c.Find(nil).All(&data.Elements)
	return data, err
}

func InsertOne(name string, element Element) (*InsertResult, error) {
	db := service.DB()
	defer db.Close()
	c := db.C(name)
	fixElement(element)
	err := c.Insert(&element)
	if err != nil {
		return nil, err
	}
	result := &InsertResult{
		Inserted: []Element{element},
	}
	return result, nil
}

func InsertMany(name string, elements []Element) (*InsertResult, error) {
	db := service.DB()
	defer db.Close()
	c := db.C(name)
	tmp := make([]interface{}, len(elements))
	for i := range elements {
		fixElement(elements[i])
		tmp[i] = elements[i]
	}
	err := c.Insert(tmp...)
	if err != nil {
		return nil, err
	}
	result := &InsertResult{
		Inserted: elements,
	}
	return result, nil
}

func Clear(name string) (*mgo.ChangeInfo, error) {
	db := service.DB()
	defer db.Close()
	c := db.C(name)
	return c.RemoveAll(bson.M{})
}

func fixElement(element Element) {
	if element["_id"] == nil {
		element["_id"] = bson.NewObjectId()
	} else if str, ok := element["_id"].(string); ok {
		element["_id"] = bson.ObjectIdHex(str)
	}
	if element["inserted"] == nil {
		element["inserted"] = time.Now()
	}
}
