package database_test

import (
	"github.com/arianht/meantweets/database"
	"google.golang.org/appengine/aetest"
	"reflect"
	"testing"
	"time"
)

// A simple DatastoreDao test that writes data and verifies reading.
// Requires "goapp test" for execution.
func TestDatastoreDao(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	dao := database.DatastoreDao{ctx}
	dao.WriteCelebrityTweet("name", "tweet")
	dao.WriteCelebrityTweet("name", "tweet") // Write a duplicate.
	dao.WriteCelebrityTweet("name", "tweet_two")
	dao.WriteCelebrityTweet("name_two", "tweet_three")

	// Sadly, App Engine Datastore takes time to fully write. Without this sleep,
	// the write won't be done in time for the read.
	time.Sleep(1 * time.Second)

	nameTweets := dao.GetCelebrityTweets("name")

	if tweetCount, expected := len(nameTweets), 2; tweetCount != expected {
		t.Errorf("Expected tweet count is %d, but was %d", expected, tweetCount)
	}
	if tweets, expected := nameTweets, []string{"tweet", "tweet_two"}; !checkContentsOfTweets(tweets, expected) {
		t.Errorf("Expected tweet content for name is %v, but was %v", expected, tweets)
	}

	nameTwoTweets := dao.GetCelebrityTweets("name_two")

	if tweetCount, expected := len(nameTwoTweets), 1; tweetCount != expected {
		t.Errorf("Expected tweet count is %d, but was %d", expected, tweetCount)
	}
	if tweets, expected := nameTwoTweets, []string{"tweet_three"}; !checkContentsOfTweets(tweets, expected) {
		t.Errorf("Expected tweet content for name_two is %v, but was %v", expected, tweets)
	}
}

// Returns true if tweets and expectedTweets have the same contents (order agnostic).
// False, otherwise.
func checkContentsOfTweets(tweets, expectedTweets []string) bool {
	tweetSet := make(map[string]int)
	for _, tweet := range tweets {
		tweetSet[tweet]++
	}
	expectedTweetSet := make(map[string]int)
	for _, tweet := range expectedTweets {
		expectedTweetSet[tweet]++
	}
	return reflect.DeepEqual(tweetSet, expectedTweetSet)
}
