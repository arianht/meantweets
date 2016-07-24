/*
Package main provides the main web server for Mean Tweets. Requires App Engine.
*/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arianht/meantweets/crawl"
	"github.com/arianht/meantweets/database"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

var celebrities []string = []string{"Johnny Depp", "Jennifer Lopez", "Justin Bieber", "Taylor Swift"}

type TwitterEmbed struct {
	Html string
}

type ContextHandler struct {
	Handle func(context.Context, http.ResponseWriter, *http.Request)
}

func (contextHandler ContextHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()
	contextHandler.Handle(ctx, writer, request)
}

func main() {
	http.Handle("/test", ContextHandler{testHandler})
	http.Handle("/crawl", ContextHandler{crawlHandler})
	http.Handle("/get_celebrities", ContextHandler{celebritiesHandler})
	http.Handle("/get_tweets", ContextHandler{tweetsHandler})
	http.Handle("/", http.FileServer(http.Dir("dist")))
	appengine.Main()
}

func testHandler(ctx context.Context, writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Content-Type", "Text/HTML")
	dao, err := database.NewDatastoreDao(ctx)
	if err != nil {
		fmt.Fprintf(writer, "Error creating datastore dao: %v", err)
		return
	}

	for _, celebrity := range celebrities {
		tweets, err := dao.GetCelebrityTweets(celebrity)
		if err != nil {
			fmt.Fprintf(writer, "Error reading database: %v", err)
			return
		}
		fmt.Fprintf(writer, "<h1>Mean Tweets for %v</h1>", celebrity)
		for _, tweet := range tweets {
			html := getTwitterHTML(tweet.Id)
			fmt.Fprintf(writer, html)
		}
	}
}

func crawlHandler(ctx context.Context, writer http.ResponseWriter, r *http.Request) {
	crawler, err := crawl.NewTwitterCrawler(ctx)
	if err != nil {
		fmt.Fprintf(writer, "Error crawling tweets: %v", err)
		return
	}
	maxTweets := 5
	crawler.Crawl(celebrities, maxTweets)
	fmt.Fprintf(writer, "Successfully crawled.")
}

func tweetsHandler(ctx context.Context, writer http.ResponseWriter, r *http.Request) {
	celebrity := r.URL.Query().Get("celebrity")
	dao, err := database.NewDatastoreDao(ctx)
	if err != nil {
		fmt.Fprintf(writer, "Error creating datastore dao: %v", err)
		return
	}

	tweets, _ := dao.GetCelebrityTweets(celebrity)
	for i, _ := range tweets {
		tweets[i].CelebrityName = ""
	}
	result, _ := json.Marshal(tweets)
	writer.Write(result)
}

func celebritiesHandler(ctx context.Context, writer http.ResponseWriter, r *http.Request) {
	result, _ := json.Marshal(celebrities)
	writer.Write(result)
}

func getTwitterHTML(id int64) string {
	res, err := http.Get(fmt.Sprintf("https://api.twitter.com/1/statuses/oembed.json?id=%v", id))
	if err != nil {
		fmt.Println("Error talking to twitter.", err)
		return "Error"
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	embeddedHTML := TwitterEmbed{}
	err = decoder.Decode(&embeddedHTML)
	if err != nil {
		fmt.Println("Error talking to twitter.", err)
		return "Error"
	}
	return embeddedHTML.Html
}
