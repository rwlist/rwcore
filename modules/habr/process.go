package habr

import (
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/rwlist/rwcore/app/model"
)

func (m *Module) process() {
	db := m.app.DB.Copy()
	defer db.Close()

	store := db.HabrDaily()

	for v := range m.reader.Read() {
		n, err := store.CountByArticleID(v.ID)
		if err != nil {
			log.Println(err)
			continue
		}

		if n != 0 {
			continue
		}

		article := model.Article{
			ID:    bson.NewObjectId(),
			URL:   v.FullURL,
			Added: time.Now(),
			Tags: map[string]interface{}{
				"habr_author_login":      v.Author.Login,
				"habr_author_fullname":   v.Author.Fullname,
				"habr_article_id":        v.ID,
				"habr_article_published": v.TimePublished,
				"habr_article_comments":  v.CommentsCount,
				"habr_article_title":     v.Title,
				"name":                   v.Title + " / Хабр",
				"habr_article_preview":   v.PreviewHTML,
				"habr_article_reading":   v.ReadingCount,
				"added_type":             "auto",
			},
		}
		err = db.Articles().InsertOne(&article)
		if err != nil {
			log.Println(err)
			continue
		}

		_, err = store.InsertOne(model.HabrDailyArticle{
			ID:        bson.NewObjectId(),
			Article:   v,
			ArticleID: v.ID,
			Added:     time.Now(),
		})
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
