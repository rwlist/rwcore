package client

import (
	"github.com/rwlist/rwcore/habr/mhabr"
	"log"
	"time"
)

var (
	waitDuration = 5 * time.Minute
)

type ReaderDailyTop struct {
	client Client
	posts  chan mhabr.ArticleDaily
	exit   chan struct{}
}

func NewReaderDailyTop(client Client) *ReaderDailyTop {
	reader := &ReaderDailyTop{
		client: client,
		posts:  make(chan mhabr.ArticleDaily),
		exit:   make(chan struct{}),
	}
	go reader.run()
	return reader
}

func (r *ReaderDailyTop) Read() <-chan mhabr.ArticleDaily {
	return r.posts
}

func (r *ReaderDailyTop) Stop() {
	close(r.exit)
}

func (r *ReaderDailyTop) run() {
	log.Println("Daily reader started")
	defer close(r.posts)
	for {
		log.Println("Read all daily article")
		r.readAll(r.posts)
		select {
		case <-r.exit:
			log.Println("Daily reader received exit signal")
			return
		case <-time.After(waitDuration):
		}
	}
}

func (r *ReaderDailyTop) readAll(ch chan<- mhabr.ArticleDaily) {
	maxPage := 100
	for page := 1; page <= maxPage; page++ {
		result, err := r.client.FetchPageDaily(page)
		if err != nil {
			log.Println("Error while fetching article.", err)
			break
		}
		if !result.Success {
			log.Println("No success in fetching article.", result)
		}
		for _, v := range result.Data.Articles {
			ch <- v
		}
		maxPage = result.Data.Pages
	}
}
