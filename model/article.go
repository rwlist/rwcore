package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
)

const (
	articleVersion = 1
)

type Article struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Version  uint          `bson:"version" json:"version"`
	Added    time.Time     `bson:"added" json:"added"`
	Modified *time.Time    `bson:"modified" json:"modified"`

	URL    string                 `bson:"url" json:"url"`
	Title  string                 `bson:"title" json:"title"`
	Tags   map[string]interface{} `bson:"tags" json:"tags"`
	Status ArticleStatus          `bson:"status" json:"status"`
}

func NewArticle(url, title string, tags map[string]interface{}) Article {
	return Article{
		ID:      bson.NewObjectId(),
		Version: articleVersion,
		Added:   time.Now(),
		URL:     url,
		Title:   title,
		Tags:    tags,
		Status:  EmptyArticleStatus(),
	}
}

func (a *Article) AddTag(tag string) error {
	i := strings.Index(tag, ":")
	var key, value string
	if i == -1 {
		key = tag
		value = ""
	} else {
		key = tag[:i]
		value = tag[i+1:]
	}

	_, ok := a.Tags[key]
	if ok {
		return fmt.Errorf("tag %s already exists", key)
	}

	a.Tags[key] = value
	return nil
}

func (a *Article) RemoveTag(tag string) error {
	_, ok := a.Tags[tag]
	if !ok {
		return fmt.Errorf("tag %s not found", tag)
	}
	delete(a.Tags, tag)
	return nil
}
