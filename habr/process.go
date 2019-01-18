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

		article := model.NewArticle(
			v.FullURL,
			v.Title + " / Хабр",
			bson.M{
				"habr": bson.M{
					"author": bson.M{
						"login":    v.Author.Login,
						"fullname": v.Author.Fullname,
					},
					"article": bson.M{
						"id":        v.ID,
						"published": v.TimePublished,
						"comments":  v.CommentsCount,
						"title":     v.Title,
						"preview":   v.PreviewHTML,
						"reading":   v.ReadingCount,
					},
				},
				"added": "auto",
			},
		)

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
