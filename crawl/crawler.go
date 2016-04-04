/*
Package for crawling for tweets, getting their sentiment score, and writing them
to the database.
*/
package crawl

import (
	"fmt"
	"github.com/arianht/meantweets/database"
)

// A crawler for crawling celebrities.
type Crawler interface {
	Crawl(celebrities []string)
}

type TwitterCrawler struct {
	Dao       database.Dao
	Twitter   TwitterFacade
	Sentiment SentimentAnalyzer
}

// Crawls twitter for mean tweets for a given celebrity and writes the results to
// the database.
func (crawler TwitterCrawler) Crawl(celebrities []string) {
	for _, celebrity := range celebrities {
		tweets, err := crawler.Twitter.GetTweets(celebrity)
		if err != nil {
			fmt.Printf("Could not get tweets for celebrity %v: %v\n", celebrity, err)
			continue
		}
		for _, tweet := range tweets {
			crawler.Dao.WriteCelebrityTweet(database.Tweet{
				CelebrityName: celebrity,
				Id:            int64(tweet.Id()),
				Score:         crawler.Sentiment.GetScoreForTweet(tweet.Text()),
			})
		}
	}
}
