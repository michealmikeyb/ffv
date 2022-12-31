package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	models "github.com/michealmikeyb/ffv/models"
	tags_pb "github.com/michealmikeyb/ffv/tags"
	users_pb "github.com/michealmikeyb/ffv/users"
	utils "github.com/michealmikeyb/ffv/utils"
)

const userkey = "user"

var secret = []byte("secret")

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	tags_pb.UnimplementedTagServiceServer
	users_pb.UnimplementedUserServiceServer
}

type tagSeenData struct {
	seen   []models.Post
	source string
}

func postsContains(s []models.Post, e models.Post) bool {
	for _, a := range s {
		if a.Content == e.Content && a.Author == e.Author && a.Source == e.Source {
			return true
		}
	}
	return false
}

func DislikePost(c *gin.Context) {
	session, err := utils.GetCassandraSession()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connecting to database"})
		return
	}
	defer session.Close()
	var dislikedPost models.Post
	if err := c.BindJSON(&dislikedPost); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't marshal json"})
		return
	}
	for _, tag := range dislikedPost.Tags {
		var (
			tag_name string
			weight   int
			source   string
			seen     []models.Post
		)

		err := session.Query(`SELECT tag_name, weight, source, seen FROM ffv.tag_list WHERE user_id = ? AND tag_name = ? AND source = ?`, c.Param("user_id"), tag, dislikedPost.Source).Scan(&tag_name, &weight, &source, &seen)
		if err == gocql.ErrNotFound {
			tag_name = tag
			weight = 0
			source = dislikedPost.Source
			seen = make([]models.Post, 0)
		} else if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error finding tag"})
			return
		}
		if weight > 0 {
			weight = weight - 1
		}
		current_post := models.Post{
			Url:     dislikedPost.Url,
			Source:  dislikedPost.Source,
			Tags:    dislikedPost.Tags,
			Author:  dislikedPost.Author,
			Likes:   int(dislikedPost.Likes),
			Content: dislikedPost.Content,
		}
		seen = append(seen, current_post)
		err = session.Query(`INSERT INTO ffv.tag_list (tag_name, weight, source, seen, user_id) VALUES (?, ?, ?, ?, ?)`, tag_name, weight, source, seen, c.Param("user_id")).Consistency(gocql.One).Exec()
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error finding tag"})
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{"success": "disliked post"})
}

func LikePost(c *gin.Context) {
	session, err := utils.GetCassandraSession()
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connecting to database"})
		return
	}
	defer session.Close()
	var likedPost models.Post
	if err := c.BindJSON(&likedPost); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't marshal json"})
		return
	}
	for _, tag := range likedPost.Tags {
		var (
			tag_name string
			weight   int
			source   string
			seen     []models.Post
		)

		err := session.Query(`SELECT tag_name, weight, source, seen FROM ffv.tag_list WHERE user_id = ? AND tag_name = ? AND source = ?`, c.Param("user_id"), tag, likedPost.Source).Scan(&tag_name, &weight, &source, &seen)
		if err == gocql.ErrNotFound {
			tag_name = tag
			weight = 0
			source = likedPost.Source
			seen = make([]models.Post, 0)
		} else if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't get tag"})
			return
		}

		weight = weight + 1
		current_post := models.Post{
			Url:     likedPost.Url,
			Source:  likedPost.Source,
			Tags:    likedPost.Tags,
			Author:  likedPost.Author,
			Likes:   int(likedPost.Likes),
			Content: likedPost.Content,
		}
		seen = append(seen, current_post)
		err = session.Query(`INSERT INTO ffv.tag_list (tag_name, weight, source, seen, user_id) VALUES (?, ?, ?, ?, ?)`, tag_name, weight, source, seen, c.Param("user_id")).Consistency(gocql.One).Exec()
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't save tag"})
			return
		}

		err = tags_pb.UpdateBuffer(tag, likedPost.Source)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't update buffer"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": "disliked post"})
}

func GetAddUser(c *gin.Context) {
	session, err := utils.GetCassandraSession()
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connecting to database"})
		return
	}
	defer session.Close()
	mastodon_id := c.Query("mastodon_id")
	var user_id string

	err = session.Query(`SELECT  user_id FROM ffv.user WHERE mastodon_id = ?`, mastodon_id).Scan(&user_id)
	if err == gocql.ErrNotFound {
		user_id = gocql.TimeUUID().String()
		err = session.Query(`INSERT INTO ffv.user (user_id, mastodon_id, mastodon_username) VALUES (?, ?, ?)`, user_id, mastodon_id, c.DefaultQuery("mastodon_id", "")).Exec()
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't get "})
			return
		}
		err = session.Query(`INSERT INTO ffv.tag_list (user_id, tag_name, weight, source) VALUES (?, ?, ?, ?)`, user_id, "popular", 20, "mastodon").Exec()
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't get "})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user_id": user_id})
		return
	} else if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't get "})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user_id": user_id})
	return

}

func GetPost(c *gin.Context) {
	session, err := utils.GetCassandraSession()
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connecting to database"})
		return
	}
	defer session.Close()
	user_id := c.Query("user_id")
	tag_list_scanner := session.Query(`SELECT tag_name, weight, source, seen FROM ffv.tag_list WHERE user_id = ?`, user_id).Iter().Scanner()
	selection_list := []string{}
	tagSeenMap := make(map[string]tagSeenData)
	log.Print("Found tags")
	for tag_list_scanner.Next() {
		var (
			tag_name string
			weight   int
			source   string
			seen     []models.Post
		)
		err := tag_list_scanner.Scan(&tag_name, &weight, &source, &seen)
		tagSeenMap[tag_name] = tagSeenData{seen: seen, source: source}
		for i := 0; i < int(weight); i++ {
			selection_list = append(selection_list, tag_name)
		}
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting tag list"})
			return
		}
	}
	if len(selection_list) < 1 {
		log.Fatal(fmt.Errorf("No tags match"))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error no tags match"})
		return
	}
	selection_index := rand.Intn(len(selection_list))
	selected_tag := selection_list[selection_index]
	log.Printf("Selected tag %s", selected_tag)
	var received_tag string

	var buffer []models.Post
	err = session.Query(`SELECT buffer, name FROM ffv.tag WHERE name = ? LIMIT 1`, selected_tag).Consistency(gocql.One).Scan(&buffer, &received_tag)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error no tag record found"})
		return
	}
	i := 0
	selected_post := buffer[i]
	for postsContains(tagSeenMap[selected_tag].seen, selected_post) {
		log.Printf("Seen %s", selected_post.Content)
		i = i + 1
		selected_post = buffer[i]
	}
	if entry, ok := tagSeenMap[selected_tag]; ok {
		log.Printf("adding to seen")
		entry.seen = append(entry.seen, selected_post)
		tagSeenMap[selected_tag] = entry
	}
	err = session.Query("UPDATE ffv.tag_list SET seen = ? WHERE user_id = ? AND tag_name = ? AND source = ?", tagSeenMap[selected_tag].seen, user_id, selected_tag, tagSeenMap[selected_tag].source).Exec()
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting tag list"})
		return
	}
	log.Print("finished adding")
	c.IndentedJSON(http.StatusOK, selected_post)
}
func main() {
	r := engine()
	r.Use(gin.Logger())
	if err := engine().Run(":80"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}

func engine() *gin.Engine {
	r := gin.New()

	// // Setup the cookie store for session management
	// r.Use(sessions.Sessions("mysession", sessions.NewCookieStore(secret)))

	// // Login and logout routes
	// r.POST("/login", login)
	// r.GET("/logout", logout)

	// // Private group, require authentication to access
	// private := r.Group("/private")
	// private.Use(AuthRequired)
	// {
	// 	private.GET("/me", me)
	// 	private.GET("/status", status)
	// }
	r.POST("/:user_id/dislike_post", DislikePost)
	r.POST("/:user_id/like_post", LikePost)
	r.GET("/get_add_user", GetAddUser)
	r.GET("/get_post", GetPost)

	return r
}

// AuthRequired is a simple middleware to check the session
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}

// login is a handler that parses a form and checks for specific data
func login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check for username and password match, usually from a database
	if username != "hello" || password != "itsme" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// Save the username in the session
	session.Set(userkey, username) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(userkey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func me(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}
