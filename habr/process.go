package habr

import (
	"github.com/rwlist/rwcore/article"
	"github.com/rwlist/rwcore/habr/mhabr"
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
)

func (m *Service) Process() {
	for v := range m.reader.Read() {
		n, err := m.habrDB.CountByArticleID(v.ID)
		if err != nil {
			log.Println(err)
			continue
		}

		if n != 0 {
			continue
		}

		a := article.NewArticle(
			v.FullURL,
			v.Title+" / Хабр",
			bson.M{
				"habr": bson.M{
					"author": bson.M{
						"login":    v.Author.Login,
						"fullname": v.Author.Fullname,
					},
					"a": bson.M{
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

		err = m.articleDB.InsertOne(&a)
		if err != nil {
			log.Println(err)
			continue
		}

		_, err = m.habrDB.InsertOne(
			mhabr.ModelDaily{
				ID:        bson.NewObjectId(),
				Article:   v,
				ArticleID: v.ID,
				Added:     time.Now(),
			},
		)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
