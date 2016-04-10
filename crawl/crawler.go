package crawl

import (
	"fmt"
	"sort"

	"github.com/arianht/meantweets/database"
	"github.com/kurrik/twittergo"
	"golang.org/x/net/context"
)

const (
	tweetsPerRequest       uint   = 100
	apiCredentialsFilename string = "../credentials.json"
)

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
func (tweets ByScore) Less(i, j int) bool { return tweets[i].Score < tweets[j].Score }

// Crawl crawls Twitter for mean tweets for a given celebrity and writes the results to
// the database. The maxTweetsPerCelebrity specifies the maximum amount of tweets a
// celebrity could have in the database.
func (crawler TwitterCrawler) Crawl(celebrities []string, maxTweetsPerCelebrity int) {
	for _, celebrity := range celebrities {
		databaseTweets, err := crawler.Dao.GetCelebrityTweets(celebrity)
		if err != nil {
			fmt.Printf("Could not get tweets for celebrity %v from database: %v\n", celebrity, err)
		}
		tweetIds := map[int64]bool{}
		for _, tweet := range databaseTweets {
			tweetIds[tweet.Id] = true
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
		databaseTweets = append(databaseTweets, getDatabaseTweetsFromTwitterTweets(twitterTweets,
			celebrity, crawler.Sentiment, tweetIds)...)
		sortAndWriteTweets(databaseTweets, crawler.Dao, maxTweetsPerCelebrity)
	}
}

func getDatabaseTweetsFromTwitterTweets(tweets []twittergo.Tweet, celebrity string,
	sentimentAnalyzer SentimentAnalyzer, tweetIds map[int64]bool) (databaseTweets []database.Tweet) {
	for _, tweet := range tweets {
		if !tweetIds[int64(tweet.Id())] {
			databaseTweets = append(databaseTweets, database.Tweet{
				CelebrityName: celebrity,
				Id:            int64(tweet.Id()),
				Score:         sentimentAnalyzer.GetScoreForTweet(tweet.Text()),
			})
		}
	}
	return
}

func sortAndWriteTweets(tweets []database.Tweet, dao database.Dao, maxTweetsPerCelebrity int) {
	sort.Sort(ByScore(tweets))

	maxIndex := maxTweetsPerCelebrity
	if maxIndex > len(tweets) {
		maxIndex = len(tweets)
	}
	dao.WriteCelebrityTweets(tweets[:maxIndex])
}

func NewTwitterCrawler(ctx context.Context) (crawler TwitterCrawler, err error) {
	dao := database.DatastoreDao{ctx}
	twitter, err := NewTwitterFacade(apiCredentialsFilename)
	if err != nil {
		return
	}
	sentiment, err := NewSentimentAnalyzer()
	if err != nil {
		return
	}
	crawler = TwitterCrawler{
		Dao:       dao,
		Twitter:   twitter,
		Sentiment: sentiment,
	}
	return
}
