package reddit

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Listing consists the initial reddit post object
type Listing struct {
	MetaData struct {
		Modhash string         `json:"modhash"`
		Dist    int            `json:"dist"`
		Posts   []PostMetaData `json:"children"`
	} `json:"data,omitempty"`
}

// PostMetaData consists of specific meta data and posts
type PostMetaData struct {
	Post struct {
		Title      string `json:"title"`
		Directlink string `json:"url"`
		Permalink  string `json:"permalink"`
		Hint       string `json:"post_hint"`
	} `json:"data,omitempty"`
}

// GetTopPosts return the 20 top posts of /r/<subreddit>
func GetTopPosts() ([]PostMetaData, error) {
	url := "https://www.reddit.com/r/puppy/.json?"
	redditClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "reddit-posts")

	res, getErr := redditClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var data Listing
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	var post []PostMetaData

	for _, p := range data.MetaData.Posts {
		if p.Post.Hint == "image" && len(post) < 3 {
			post = append(post, p)
		}
	}

	data.MetaData.Posts = post

	return data.MetaData.Posts, nil
}
