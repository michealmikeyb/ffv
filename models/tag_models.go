package models

import (
	"time"
)

type Post struct {
	Url     string   `cql:"url"`
	Tags    []string `cql:"tags"`
	Source  string   `cql:"source"`
	Author  string   `cql:"author"`
	Content string   `cql:"content"`
	Likes   int      `cql:"likes"`
}

type MastodonResponse struct {
	ID                 string      `json:"id"`
	CreatedAt          time.Time   `json:"created_at"`
	InReplyToID        interface{} `json:"in_reply_to_id"`
	InReplyToAccountID interface{} `json:"in_reply_to_account_id"`
	Sensitive          bool        `json:"sensitive"`
	SpoilerText        string      `json:"spoiler_text"`
	Visibility         string      `json:"visibility"`
	Language           string      `json:"language"`
	URI                string      `json:"uri"`
	URL                string      `json:"url"`
	RepliesCount       int         `json:"replies_count"`
	ReblogsCount       int         `json:"reblogs_count"`
	FavouritesCount    int         `json:"favourites_count"`
	EditedAt           interface{} `json:"edited_at"`
	Content            string      `json:"content"`
	Reblog             interface{} `json:"reblog"`
	Account            struct {
		ID             string        `json:"id"`
		Username       string        `json:"username"`
		Acct           string        `json:"acct"`
		DisplayName    string        `json:"display_name"`
		Locked         bool          `json:"locked"`
		Bot            bool          `json:"bot"`
		Discoverable   bool          `json:"discoverable"`
		Group          bool          `json:"group"`
		CreatedAt      time.Time     `json:"created_at"`
		Note           string        `json:"note"`
		URL            string        `json:"url"`
		Avatar         string        `json:"avatar"`
		AvatarStatic   string        `json:"avatar_static"`
		Header         string        `json:"header"`
		HeaderStatic   string        `json:"header_static"`
		FollowersCount int           `json:"followers_count"`
		FollowingCount int           `json:"following_count"`
		StatusesCount  int           `json:"statuses_count"`
		LastStatusAt   string        `json:"last_status_at"`
		Emojis         []interface{} `json:"emojis"`
		Fields         []interface{} `json:"fields"`
	} `json:"account,omitempty"`
	MediaAttachments []struct {
		ID               string      `json:"id"`
		Type             string      `json:"type"`
		URL              string      `json:"url"`
		PreviewURL       string      `json:"preview_url"`
		RemoteURL        string      `json:"remote_url"`
		PreviewRemoteURL interface{} `json:"preview_remote_url"`
		TextURL          interface{} `json:"text_url"`
		Meta             struct {
			Original struct {
				Width  int     `json:"width"`
				Height int     `json:"height"`
				Size   string  `json:"size"`
				Aspect float64 `json:"aspect"`
			} `json:"original"`
			Small struct {
				Width  int     `json:"width"`
				Height int     `json:"height"`
				Size   string  `json:"size"`
				Aspect float64 `json:"aspect"`
			} `json:"small"`
		} `json:"meta"`
		Description interface{} `json:"description"`
		Blurhash    string      `json:"blurhash"`
	} `json:"media_attachments"`
	Mentions []interface{} `json:"mentions"`
	Tags     []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"tags"`
	Emojis      []interface{} `json:"emojis"`
	Card        interface{}   `json:"card"`
	Poll        interface{}   `json:"poll"`
	Application struct {
		Name    string `json:"name"`
		Website string `json:"website"`
	} `json:"application,omitempty"`
}
