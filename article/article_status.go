package article

import "time"

type ArticleStatus struct {
	Clicks    uint       `bson:"clicks" json:"clicks"`
	LastClick *time.Time `bson:"lastClick" json:"lastClick"`

	ReadStatus       string     `bson:"readStatus" json:"readStatus"` // unopened|viewed|completed
	ReadStatusChange *time.Time `bson:"readStatusChange" json:"readStatusChange"`

	Rating int `bson:"rating" json:"rating"`
}

func EmptyArticleStatus() ArticleStatus {
	return ArticleStatus{
		ReadStatus: "unopened",
	}
}

func ValidArticleReadStatus(readStatus string) bool {
	return readStatus == "unopened" || readStatus == "viewed" || readStatus == "completed"
}
