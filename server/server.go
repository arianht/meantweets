/*
Package server provides the main web server for Mean Tweets. Requires App Engine.

Run locally with:
goapp serve github.com/arianht/meantweets/server

It will automatically detect changes in the code.

Push to the cloud with:
appcfg.py A <project_id> -V <version> update github.com/arianht/meantweets/server
*/
package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arianht/meantweets/crawl"
	"github.com/arianht/meantweets/database"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

var celebrities []string = []string{"Johnny Depp", "Jennifer Lopez", "Justin Bieber", "Taylor Swift"}

type TwitterEmbed struct {
	Html string
}

func init() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/crawl", crawlHandler)
}

func rootHandler(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Content-Type", "Text/HTML")
	ctx := appengine.NewContext(r)
	dao := database.DatastoreDao{ctx}
	for _, celebrity := range celebrities {
		tweets, err := dao.GetCelebrityTweets(celebrity)
		if err != nil {
			fmt.Fprintf(writer, "Error reading database: %v", err)
			return
		}
		fmt.Fprintf(writer, "<h1>Mean Tweets for %v</h1>", celebrity)
		for _, tweet := range tweets {
			html, err := getTwitterHTML(ctx, tweet.Id)
			if err != nil {
				fmt.Fprintf(writer, "Error talking to Twitter: %v", err)
				return
			}
			fmt.Fprintf(writer, html)
		}
	}
}

func crawlHandler(writer http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	crawler, err := crawl.NewTwitterCrawler(ctx)
	if err != nil {
		fmt.Fprintf(writer, "Error crawling tweets: %v", err)
		return
	}
	maxTweets := 5
	crawler.Crawl(celebrities, maxTweets)
	fmt.Fprintf(writer, "Successfully crawled.")
}

func getTwitterHTML(ctx context.Context, id int64) (html string, err error) {
	client := urlfetch.Client(ctx)
	res, err := client.Get(fmt.Sprintf("https://api.twitter.com/1/statuses/oembed.json?id=%v", id))
	if err != nil {
		return
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	embeddedHTML := TwitterEmbed{}
	err = decoder.Decode(&embeddedHTML)
	if err != nil {
		return
	}
	html = embeddedHTML.Html
	return
}
