package tags

import (
	"encoding/json"
	fmt "fmt"
	"log"
	"net/http"
	"sort"

	"github.com/gocql/gocql"
	models "github.com/michealmikeyb/ffv/models"
	utils "github.com/michealmikeyb/ffv/utils"
)

const mastodon_base_url = "https://mastodon.social"

type ByLikes []models.Post

func (a ByLikes) Len() int           { return len(a) }
func (a ByLikes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLikes) Less(i, j int) bool { return a[i].Likes < a[j].Likes }

func UpdateBuffer(tag_name string, tag_source string) error {
	if tag_source == "mastodon" {
		var mastodon_tag_url string
		if tag_name == "popular" {
			mastodon_tag_url = fmt.Sprintf("%s/%s", mastodon_base_url, "api/v1/trends/statuses")
		} else {
			mastodon_tag_url = fmt.Sprintf("%s/%s/%s", mastodon_base_url, "api/v1/timelines/tag", tag_name)
		}
		resp, err := http.Get(mastodon_tag_url)
		if err != nil {
			log.Printf("error getting mastodon response")
			return err
		}
		defer resp.Body.Close()

		var mas_responses []models.MastodonResponse
		err = json.NewDecoder(resp.Body).Decode(&mas_responses)
		if err != nil {
			log.Printf("error marshalling json")
			return err
		}

		var posts []models.Post
		for _, res := range mas_responses {
			var tags []string
			for _, tag := range res.Tags {
				tags = append(tags, tag.Name)
			}
			post := models.Post{
				Url:     res.URL,
				Source:  "mastodon",
				Tags:    tags,
				Author:  res.Account.Username,
				Likes:   res.FavouritesCount,
				Content: res.Content,
			}
			posts = append(posts, post)
		}
		sort.Sort(ByLikes(posts))

		session, err := utils.GetCassandraSession()
		if err != nil {
			log.Printf("error connecting to cassandra")
			return err
		}
		defer session.Close()

		err = session.Query(`INSERT INTO ffv.tag (name, source, buffer) VALUES (?, ?, ?)`, tag_name, tag_source, posts).Consistency(gocql.One).Exec()
		if err != nil {
			log.Fatal(err)
		}

		return nil

	} else {
		return fmt.Errorf("Tag source not supported")
	}
}