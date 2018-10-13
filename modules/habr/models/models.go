package models

import "time"

type Flow struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Alias string `json:"alias"`
	URL   string `json:"url"`
	Path  string `json:"path"`
}

type Author struct {
	ID             int       `json:"id"`
	Login          string    `json:"login"`
	TimeRegistered time.Time `json:"time_registered"`
	Score          float64   `json:"score"`
	Fullname       string    `json:"fullname"`
	Specializm     string    `json:"specializm"`
	Sex            int       `json:"sex"`
	Rating         float64   `json:"rating"`
	RatingPosition int       `json:"rating_position"`
	Path           string    `json:"path"`
	Geo            struct {
		Country interface{} `json:"country"`
		Region  interface{} `json:"region"`
		City    interface{} `json:"city"`
	} `json:"geo"`
	Counters struct {
		Posts     int `json:"posts"`
		Comments  int `json:"comments"`
		Followed  int `json:"followed"`
		Followers int `json:"followers"`
	} `json:"counters"`
	Badges []struct {
		Alias       string `json:"alias"`
		Title       string `json:"title"`
		Plural      string `json:"plural"`
		Description string `json:"description"`
	} `json:"badges"`
	Avatar     string        `json:"avatar"`
	IsReadonly bool          `json:"is_readonly"`
	IsRc       bool          `json:"is_rc"`
	CommonTags []interface{} `json:"common_tags"`
	Contacts   []interface{} `json:"contacts"`
}

type Hub struct {
	ID               int         `json:"id"`
	CountPosts       int         `json:"count_posts"`
	CountSubscribers int         `json:"count_subscribers"`
	IsProfiled       bool        `json:"is_profiled"`
	Rating           float64     `json:"rating"`
	Alias            string      `json:"alias"`
	Title            string      `json:"title"`
	TagsString       string      `json:"tags_string"`
	About            string      `json:"about"`
	AboutSmall       interface{} `json:"about_small"`
	Flow             interface{} `json:"flow"`
	IsMembership     bool        `json:"is_membership"`
	IsCompany        bool        `json:"is_company"`
	Icon             string      `json:"icon"`
	Path             string      `json:"path"`
}

type ArticleDaily struct {
	ID              int         `json:"id"`
	IsTutorial      bool        `json:"is_tutorial"`
	TimePublished   time.Time   `json:"time_published"`
	TimeInteresting time.Time   `json:"time_interesting"`
	CommentsCount   int         `json:"comments_count"`
	Score           int         `json:"score"`
	VotesCount      int         `json:"votes_count"`
	FavoritesCount  int         `json:"favorites_count"`
	TagsString      string      `json:"tags_string"`
	Title           string      `json:"title"`
	PreviewHTML     string      `json:"preview_html"`
	TextCut         string      `json:"text_cut"`
	IsCommentsHide  int         `json:"is_comments_hide"`
	IsRecoveryMode  bool        `json:"is_recovery_mode"`
	Flows           []Flow      `json:"flows"`
	Hubs            []Hub       `json:"hubs"`
	ReadingCount    int         `json:"reading_count"`
	Path            string      `json:"path"`
	FullURL         string      `json:"full_url"`
	Author          *Author     `json:"author"`
	HasPolls        bool        `json:"has_polls"`
	URL             string      `json:"url"`
	PostType        int         `json:"post_type"`
	PostTypeStr     string      `json:"post_type_str"`
	Vote            interface{} `json:"vote"`
	IsCanVote       bool        `json:"is_can_vote"`
	IsHabred        bool        `json:"is_habred"`
	IsInteresting   bool        `json:"is_interesting"`
	IsFavorite      bool        `json:"is_favorite"`
	CommentsNew     int         `json:"comments_new"`
}

type PageDaily struct {
	Data struct {
		Articles []ArticleDaily `json:"articles"`
		Pages    int            `json:"pages"`
	} `json:"data"`
	Success bool `json:"success"`
}
