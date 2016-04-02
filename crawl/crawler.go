/*
Package for crawling for tweets, getting their sentiment score, and writing them
to the database.
*/
package crawl

import (
	"github.com/arianht/meantweets/database"
)

type Crawler interface {
	Crawl()
}

type TwitterCrawler struct {
	Dao         database.Dao
	Twitter     TwitterFacade
	Sentiment   SentimentAnalyzer
	Celebrities []string
}

// Crawls twitter for mean tweets for a given celebrity and writes the results to
// the database.
func (crawler TwitterCrawler) Crawl() {
	for _, celebrity := range crawler.Celebrities {
		tweets := crawler.Twitter.GetTweets(celebrity)
		for _, tweet := range tweets {
			crawler.Dao.WriteCelebrityTweet(database.Tweet{
				CelebrityName: celebrity,
				Id:            tweet.Id,
				Score:         crawler.Sentiment.GetScoreForTweet(tweet.Content),
			})
		}
	}
}
