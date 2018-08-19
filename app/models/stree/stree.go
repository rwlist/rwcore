package stree

import (
	"errors"
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/rwlist/rwcore/app/models/mongodb"
)

// TODO: Add mutex for single user access locks
// TODO: Add userID to every node for security

var (
	DIRECTORY = "directory"
	FILE      = "file"
)

var (
	ErrNotDirectory      = errors.New("Specified node not a directory")
	ErrNotEmptyDirectory = errors.New("Can't delete not empty directory")
	ErrRootDirectory     = errors.New("Can't delete root directory")
)

type Session struct {
	dbSession *mgo.Session
	userID    bson.ObjectId

	Nodes *mgo.Collection
	Roots *mgo.Collection
}

type Root struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	UserID bson.ObjectId `bson:"userID"`
	RootID bson.ObjectId `bson:"rootID"`
}

type Node struct {
	ID       bson.ObjectId  `bson:"_id"`
	ParentID *bson.ObjectId `bson:"parentID,omitempty"`
	Name     string         `bson:"name"`
	Type     string         `bson:"type"`
	Content  interface{}    `bson:"content"`
}

func NewSession(userID bson.ObjectId) Session {
	dbSession := mongodb.BaseSession.Clone()
	s := Session{
		dbSession,
		userID,
		dbSession.DB("rwlist").C("rwlist_nodes"),
		dbSession.DB("rwlist").C("rwlist_roots"),
	}
	log.Println("New stree session")
	return s
}

func (s Session) Close() {
	s.dbSession.Close()
	log.Println("Old stree session closed")
}

func (s Session) GetRoot() (*Root, error) {
	var root Root
	err := s.Roots.Find(bson.M{"userID": s.userID}).One(&root)
	if err == mgo.ErrNotFound {
		node := Node{
			bson.NewObjectId(),
			nil,
			"",
			DIRECTORY,
			nil,
		}
		err = s.Nodes.Insert(&node)
		if err != nil {
			return nil, err
		}
		root = Root{
			bson.NewObjectId(),
			s.userID,
			node.ID,
		}
		err = s.Roots.Insert(&root)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	return &root, nil
}

func (s Session) GetNode(nodeID bson.ObjectId) (*Node, error) {
	var node Node
	err := s.Nodes.Find(bson.M{"_id": nodeID}).One(&node)
	if err != nil {
		return nil, err
	}
	return &node, nil
}

func (s Session) AssureDirectoryExists(dirID bson.ObjectId) error {
	var dir Node
	err := s.Nodes.Find(bson.M{"_id": dirID}).One(&dir)
	if err != nil {
		return err
	}
	if dir.Type != DIRECTORY {
		return ErrNotDirectory
	}
	return nil
}

func (s Session) insertNode(node *Node) error {
	if err := s.AssureDirectoryExists(*node.ParentID); err != nil {
		return err
	}
	node.ID = bson.NewObjectId()
	if err := s.Nodes.Insert(&node); err != nil {
		return err
	}
	return nil
}

func (s Session) CreateDir(parentID bson.ObjectId, name string) (*Node, error) {
	node := &Node{
		ParentID: &parentID,
		Name:     name,
		Type:     DIRECTORY,
	}
	if err := s.insertNode(node); err != nil {
		return nil, err
	}
	return node, nil
}

func (s Session) CreateFile(parentID bson.ObjectId, name string, content interface{}) (*Node, error) {
	node := &Node{
		ParentID: &parentID,
		Name:     name,
		Type:     FILE,
		Content:  content,
	}
	if err := s.insertNode(node); err != nil {
		return nil, err
	}
	return node, nil
}

func (s Session) ListDirectory(directoryID bson.ObjectId) ([]*Node, error) {
	if err := s.AssureDirectoryExists(directoryID); err != nil {
		return nil, err
	}
	list := make([]*Node, 0)
	err := s.Nodes.Find(bson.M{"parentID": directoryID}).All(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s Session) Delete(nodeID bson.ObjectId) error {
	var node Node
	err := s.Nodes.Find(bson.M{"_id": nodeID}).One(&node)
	if err != nil {
		return err
	}
	if node.ParentID == nil {
		return ErrRootDirectory
	}
	n, err := s.Nodes.Find(bson.M{"parentID": nodeID}).Count()
	if err != nil {
		return err
	}
	if n != 0 {
		return ErrNotEmptyDirectory
	}
	err = s.Nodes.Remove(bson.M{"_id": nodeID})
	return err
}

func (s Session) Rename(nodeID bson.ObjectId, newName string) error {
	var node Node
	err := s.Nodes.Find(bson.M{"_id": nodeID}).One(&node)
	if err != nil {
		return err
	}
	if node.ParentID == nil {
		return ErrRootDirectory
	}
	node.Name = newName
	return s.Nodes.UpdateId(nodeID, node)
}
