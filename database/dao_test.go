package database

import (
	"reflect"
	"testing"
	"time"

	"google.golang.org/appengine/aetest"
)

// A simple DatastoreDao test that writes data and verifies reading.
// Requires "goapp test" for execution.
func TestDatastoreDao(t *testing.T) {
	johnTweetOne := Tweet{"John", 0, 500}
	johnTweetTwo := Tweet{"John", 1, 12}
	jenTweetOne := Tweet{"Jen", 2, 750}
	jenTweetTwo := Tweet{"Jen", 3, 500}
	jenTweetThree := Tweet{"Jen", 4, 300}
	jenTweetFour := Tweet{"Jen", 5, 751}

	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	dao := DatastoreDao{ctx}
	dao.WriteCelebrityTweet(johnTweetOne)
	dao.WriteCelebrityTweet(johnTweetOne) // Write a duplicate.
	dao.WriteCelebrityTweet(johnTweetTwo)
	dao.WriteCelebrityTweet(jenTweetOne)
	dao.WriteCelebrityTweet(jenTweetTwo)
	dao.WriteCelebrityTweet(jenTweetThree)
	dao.WriteCelebrityTweet(jenTweetFour)

	// Sadly, App Engine Datastore takes time to fully write. Without this sleep,
	// the write won't be done in time for the read.
	time.Sleep(2 * time.Second)

	johnTweets := dao.GetCelebrityTweets("John")

	if tweetCount, expected := len(johnTweets), 2; tweetCount != expected {
		t.Errorf("Expected tweet count for John is %d, but was %d", expected, tweetCount)
	}
	if tweets, expected := johnTweets, []Tweet{johnTweetOne, johnTweetTwo}; !reflect.DeepEqual(tweets, expected) {
		t.Errorf("Expected tweet content for John is %v, but was %v", expected, tweets)
	}

	jenTweets := dao.GetCelebrityTweets("Jen")

	if tweetCount, expected := len(jenTweets), 4; tweetCount != expected {
		t.Errorf("Expected tweet count for Jen is %d, but was %d", expected, tweetCount)
	}

	if tweets, expected := jenTweets, []Tweet{jenTweetFour, jenTweetOne, jenTweetTwo, jenTweetThree}; !reflect.DeepEqual(tweets, expected) {
		t.Errorf("Expected tweet content for Jen is %v, but was %v", expected, tweets)
	}

	dao.DeleteAllTweetsForCelebrity("John")
	time.Sleep(2 * time.Second)
	johnTweetsAfterDelete := dao.GetCelebrityTweets("John")

	if tweetCount, expected := len(johnTweetsAfterDelete), 0; tweetCount != expected {
		t.Errorf("Expected tweet count for John after delete is %d, but was %d", expected, tweetCount)
	}
}
