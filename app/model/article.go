package model

import (
	"fmt"
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

type ArticleStatus struct {
	Clicks    uint       `bson:"clicks" json:"clicks"`
	LastClick *time.Time `bson:"lastClick" json:"lastClick"`

	ReadStatus       string     `bson:"readStatus" json:"readStatus"` // unopened|viewed|completed
	ReadStatusChange *time.Time `bson:"readStatusChange" json:"readStatusChange"`

	Rating int `bson:"rating" json:"rating"`
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

func EmptyArticleStatus() ArticleStatus {
	return ArticleStatus{
		ReadStatus: "unopened",
	}
}

func ValidArticleReadStatus(readStatus string) bool {
	return readStatus == "unopened" || readStatus == "viewed" || readStatus == "completed"
}

func (a *Article) BumpVersion() (bool, error) {
	updated := a.Version != articleVersion
	var err error

	for a.Version != articleVersion {
		switch a.Version {
		case 0:
			err = a.upgrade0()
		default:
			return false, fmt.Errorf("can't upgrade version %d of Article", a.Version)
		}

		if err != nil {
			return false, err
		}
	}

	return updated, nil
}

func (a *Article) upgrade0() error {
	a.Version = 1
	now := time.Now()

	clicks, ok := a.Tags["clicks"].(uint)
	if ok {
		a.Status.Clicks = clicks
		a.Status.LastClick = &now
	}
	delete(a.Tags, "clicks")
	delete(a.Tags, "lastClicked")

	name, ok := a.Tags["name"].(string)
	if ok {
		a.Title = name
	} else {
		a.Title = "Untitled"
	}
	delete(a.Tags, "name")

	status, ok := a.Tags["status"].(string)
	if ok {
		switch status {
		case "IRL":
			a.Status.ReadStatus = "completed"
			a.Status.ReadStatusChange = &now
			a.Status.Rating = 1
		case "AWSM":
			a.Status.ReadStatus = "viewed"
			a.Status.ReadStatusChange = &now
			a.Status.Rating = 1
		case "DSMS":
			a.Status.ReadStatus = "viewed"
			a.Status.ReadStatusChange = &now
			a.Status.Rating = 0
		case "REM":
			a.Status.ReadStatus = "unopened"
			a.Status.Rating = -1
		default:
			a.Status.ReadStatus = "unopened"
		}
	} else {
		a.Status.ReadStatus = "unopened"
	}
	delete(a.Tags, "status")
	delete(a.Tags, "statusChanged")

	return nil
}
