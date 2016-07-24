package crawl

import (
	"fmt"
	"testing"
	"time"

	"github.com/arianht/meantweets/database"
	"github.com/arianht/meantweets/util"
	"golang.org/x/net/context"
)

func TestRealCrawl(t *testing.T) {
	// Skip test if credentials are not found.
	_, err := util.GetTwitterAPICredentialsFromFile("credentials.json")
	if err != nil {
		t.Skipf("Error getting Twitter credentials: %v. Skipping crawler integration test.", err)
	}
	ctx := context.Background()

	twitterCrawler, err := NewTwitterCrawler(ctx)
	if err != nil {
		t.Fatal(err)
	}
	celebrities := []string{"Johnny Depp", "Jennifer Lopez", "Justin Bieber", "Taylor Swift"}
	maxTweets := 5
	twitterCrawler.Crawl(celebrities, maxTweets)

	// Sadly, App Engine Datastore takes time to fully write. Without this sleep,
	// the write won't be done in time for the read.
	time.Sleep(2 * time.Second)

	fmt.Println("Printing out results from the Crawler Integration Test")
	for _, celebrity := range celebrities {
		tweets, err := twitterCrawler.Dao.GetCelebrityTweets(celebrity)
		if err != nil {
			t.Fatal(err)
		}
		// Make sure we crawled tweets and wrote the top ones to the database.
		if len(tweets) != maxTweets {
			t.Errorf("Expected %d tweets for celebrity %v in database, but was %v.", maxTweets, celebrity, len(tweets))
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
