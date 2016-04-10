package crawl

import (
	"fmt"
	"testing"
	"time"

	"github.com/arianht/meantweets/database"
	"google.golang.org/appengine/aetest"
)

func TestRealCrawl(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	dao := database.DatastoreDao{ctx}
	twitter, err := NewTwitterFacade("../credentials.json")
	if err != nil {
		t.Fatal(err)
	}
	sentiment, err := NewSentimentAnalyzer()
	if err != nil {
		t.Fatal(err)
	}
	twitterCrawler := TwitterCrawler{
		Dao:       dao,
		Twitter:   twitter,
		Sentiment: sentiment,
	}

	celebrities := []string{"Johnny Depp", "Jennifer Lopez", "Justin Bieber", "Taylor Swift"}
	maxTweets := 5
	twitterCrawler.Crawl(celebrities, maxTweets)

	// Sadly, App Engine Datastore takes time to fully write. Without this sleep,
	// the write won't be done in time for the read.
	time.Sleep(2 * time.Second)

	fmt.Println("Printing out results from the Crawler Integration Test")
	for _, celebrity := range celebrities {
		tweets, err := dao.GetCelebrityTweets(celebrity)
		if err != nil {
			t.Fatal(err)
		}
		// Make sure we crawled tweets and wrote the top ones to the database.
		if len(tweets) != maxTweets {
			t.Errorf("Expected %d tweets for celebrity %v in database.", maxTweets, celebrity)
		}
		printCelebrityResults(celebrity, tweets)
	}
	fmt.Println("Crawler integration test finished. Manually verify results.")
}

func printCelebrityResults(celebrity string, tweets []database.Tweet) {
	fmt.Printf("================================\n%v mean tweets: \n", celebrity)
	for _, tweet := range tweets {
		fmt.Printf("Score: %v - Link: http://twitter.com/random/status/%v\n", tweet.Score, tweet.Id)
	}
	fmt.Printf("\n")
}
