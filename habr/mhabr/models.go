package mhabr

import "time"

type Flow struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Alias string `json:"alias"`
	URL   string `json:"url"`
	Path  string `json:"path"`
}

type Author struct {
	ID             string      `json:"id"`
	Login          string      `json:"login"`
	TimeRegistered time.Time   `json:"time_registered"`
	Score          float64     `json:"score"`
	Fullname       interface{} `json:"fullname"`
	Specializm     string      `json:"specializm"`
	Sex            string      `json:"sex"`
	Rating         float64     `json:"rating"`
	RatingPosition interface{} `json:"rating_position"`
	Path           string      `json:"path"`
	Geo            struct {
		Country interface{} `json:"country"`
		Region  interface{} `json:"region"`
		City    interface{} `json:"city"`
	} `json:"geo"`
	Counters struct {
		Posts     string `json:"posts"`
		Comments  string `json:"comments"`
		Followed  string `json:"followed"`
		Followers string `json:"followers"`
	} `json:"counters"`
	Badges []struct {
		Alias       string `json:"alias"`
		Title       string `json:"title"`
		Plural      string `json:"plural"`
		Description string `json:"description"`
	} `json:"badges"`
	Avatar       string        `json:"avatar"`
	IsReadonly   bool          `json:"is_readonly"`
	IsRc         bool          `json:"is_rc"`
	IsSubscribed bool          `json:"is_subscribed"`
	CommonTags   []interface{} `json:"common_tags"`
	Contacts     []interface{} `json:"contacts"`
}

type Hub struct {
	ID               string `json:"id"`
	CountPosts       string `json:"count_posts"`
	CountSubscribers string `json:"count_subscribers"`
	IsProfiled       bool   `json:"is_profiled"`
	Rating           string `json:"rating"`
	Alias            string `json:"alias"`
	Title            string `json:"title"`
	TagsString       string `json:"tags_string"`
	About            string `json:"about"`
	AboutSmall       string `json:"about_small"`
	Flow             Flow   `json:"flow"`
	IsMembership     bool   `json:"is_membership"`
	IsCompany        bool   `json:"is_company"`
	Icon             string `json:"icon"`
	Path             string `json:"path"`
}

type ArticleDaily struct {
	ID              string      `json:"id"`
	IsTutorial      bool        `json:"is_tutorial"`
	TimePublished   time.Time   `json:"time_published"`
	TimeInteresting time.Time   `json:"time_interesting"`
	CommentsCount   string      `json:"comments_count"`
	Score           string      `json:"score"`
	VotesCount      string      `json:"votes_count"`
	FavoritesCount  string      `json:"favorites_count"`
	Lang            string      `json:"lang"`
	TagsString      string      `json:"tags_string"`
	Title           string      `json:"title"`
	PreviewHTML     string      `json:"preview_html"`
	TextCut         string      `json:"text_cut"`
	IsCommentsHide  string      `json:"is_comments_hide"`
	Flows           []Flow      `json:"flows"`
	Hubs            []Hub       `json:"hubs"`
	ReadingCount    int         `json:"reading_count"`
	Path            string      `json:"path"`
	FullURL         string      `json:"full_url"`
	Author          Author      `json:"author"`
	HasPolls        bool        `json:"has_polls"`
	URL             string      `json:"url"`
	PostType        string      `json:"post_type"`
	PostTypeStr     string      `json:"post_type_str"`
	Vote            interface{} `json:"vote"`
	IsCanVote       bool        `json:"is_can_vote"`
	IsHabred        bool        `json:"is_habred"`
	IsInteresting   bool        `json:"is_interesting"`
	IsFavorite      bool        `json:"is_favorite"`
	IsRecoveryMode  bool        `json:"is_recovery_mode"`
	CommentsNew     interface{} `json:"comments_new"`
	SourceAuthor    string      `json:"source_author,omitempty"`
	SourceLink      string      `json:"source_link,omitempty"`
}

type PageDaily struct {
	Data struct {
		Articles []ArticleDaily `json:"articles"`
		Pages    int            `json:"pages"`
	} `json:"data"`
	Success bool `json:"success"`
}
