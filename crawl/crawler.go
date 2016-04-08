package crawl

import (
	"fmt"
	"sort"

	"github.com/arianht/meantweets/database"
)

const tweetsPerRequest uint = 100

// Crawler is an interface for crawling celebrities.
type Crawler interface {
	Crawl(celebrities []string)
}

// TwitterCrawler implements Crawler and is used to crawl tweets about celebrities on Twitter.
type TwitterCrawler struct {
	Dao       database.Dao
	Twitter   TwitterFacade
	Sentiment SentimentAnalyzer
}

type ByScore []database.Tweet

func (tweets ByScore) Len() int           { return len(tweets) }
func (tweets ByScore) Swap(i, j int)      { tweets[i], tweets[j] = tweets[j], tweets[i] }
func (tweets ByScore) Less(i, j int) bool { return tweets[i].Score >= tweets[j].Score }

// Crawl crawls Twitter for mean tweets for a given celebrity and writes the results to
// the database. The maxTweetsPerCelebrity specifies the maximum amount of tweets a
// celebrity could have in the database.
func (crawler TwitterCrawler) Crawl(celebrities []string, maxTweetsPerCelebrity int) {
	for _, celebrity := range celebrities {
		databaseTweets, err := crawler.Dao.GetCelebrityTweets(celebrity)
		if err != nil {
			fmt.Printf("Could not get tweets for celebrity %v from database: %v\n", celebrity, err)
		}
		err = crawler.Dao.DeleteAllTweetsForCelebrity(celebrity)
		if err != nil {
			fmt.Printf("Failed to delete tweets for celebrity %v from database: %v\n", celebrity, err)
		}
		twitterTweets, err := crawler.Twitter.GetTweets(celebrity, tweetsPerRequest)
		if err != nil {
			fmt.Printf("Could not get tweets for celebrity %v: %v\n", celebrity, err)
			continue
		}
		for _, tweet := range twitterTweets {
			databaseTweets = append(databaseTweets, database.Tweet{
				CelebrityName: celebrity,
				Id:            int64(tweet.Id()),
				Score:         crawler.Sentiment.GetScoreForTweet(tweet.Text()),
			})
		}
		sort.Sort(ByScore(databaseTweets))
		maxIndex := maxTweetsPerCelebrity
		if maxIndex > len(databaseTweets) {
			maxIndex = len(databaseTweets)
		}
		crawler.Dao.WriteCelebrityTweets(databaseTweets[:maxIndex])
	}
}
