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
				"habr.author.login":      v.Author.Login,
				"habr.author.fullname":   v.Author.Fullname,
				"habr.article.id":        v.ID,
				"habr.article.published": v.TimePublished,
				"habr.article.comments":  v.CommentsCount,
				"habr.article.title":     v.Title,
				"name":                   v.Title + " / Хабр",
				"habr.article.preview":   v.PreviewHTML,
				"habr.article.reading":   v.ReadingCount,
				"added.type":             "auto",
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
