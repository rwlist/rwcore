package habr

import (
	"log"
	"time"

	"github.com/rwlist/rwcore/modules/habr/models"
)

var (
	waitDuration = 5 * time.Minute
)

type ReaderDailyTop struct {
	client Client
	posts  chan models.ArticleDaily
	exit   chan struct{}
}

func NewReaderDailyTop() *ReaderDailyTop {
	reader := &ReaderDailyTop{
		client: NewClient(),
		posts:  make(chan models.ArticleDaily),
		exit:   make(chan struct{}),
	}
	// TODO: enable later
	// go reader.run()
	return reader
}

func (r *ReaderDailyTop) Read() <-chan models.ArticleDaily {
	return r.posts
}

func (r *ReaderDailyTop) Stop() {
	close(r.exit)
}

func (r *ReaderDailyTop) run() {
	log.Println("Daily reader started")
	for {
		r.readAll(r.posts)
		select {
		case <-r.exit:
			log.Println("Daily reader received exit signal")
			return
		case <-time.After(waitDuration):
		}
	}
	close(r.posts)
}

func (r *ReaderDailyTop) readAll(ch chan<- models.ArticleDaily) {
	maxPage := 100
	for page := 1; page <= maxPage; page++ {
		result, err := r.client.FetchPageDaily(page)
		if err != nil {
			log.Println("Error while fetching articles.", err)
			break
		}
		if !result.Success {
			log.Println("Unsuccess while fetching articles.", result)
		}
		for _, v := range result.Data.Articles {
			ch <- v
		}
		maxPage = result.Data.Pages
	}
}