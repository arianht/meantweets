package database

import (
	"reflect"
	"testing"
	"time"

	"golang.org/x/net/context"
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

	ctx := context.Background()

	dao, err := NewDatastoreDao(ctx)
	if err != nil {
		t.Errorf("Failed to create dao %v", err)
	}

	dao.WriteCelebrityTweets([]Tweet{
		johnTweetOne,
		johnTweetOne,
		johnTweetTwo,
		jenTweetOne,
		jenTweetTwo,
		jenTweetThree,
		jenTweetFour,
	})

	// Sadly, App Engine Datastore takes time to fully write. Without this sleep,
	// the write won't be done in time for the read.
	time.Sleep(1 * time.Second)

	johnTweets, err := dao.GetCelebrityTweets("John")
	if err != nil {
		t.Errorf("Expected err %v to be nil.", err)
	}

	if tweetCount, expected := len(johnTweets), 2; tweetCount != expected {
		t.Errorf("Expected tweet count for John is %d, but was %d", expected, tweetCount)
	}
	if tweets, expected := johnTweets, []Tweet{johnTweetTwo, johnTweetOne}; !reflect.DeepEqual(tweets, expected) {
		t.Errorf("Expected tweet content for John is %v, but was %v", expected, tweets)
	}

	jenTweets, err := dao.GetCelebrityTweets("Jen")
	if err != nil {
		t.Errorf("Expected err %v to be nil.", err)
	}

	if tweetCount, expected := len(jenTweets), 4; tweetCount != expected {
		t.Errorf("Expected tweet count for Jen is %d, but was %d", expected, tweetCount)
	}

	if tweets, expected := jenTweets, []Tweet{jenTweetThree, jenTweetTwo, jenTweetOne, jenTweetFour}; !reflect.DeepEqual(tweets, expected) {
		t.Errorf("Expected tweet content for Jen is %v, but was %v", expected, tweets)
	}

	dao.DeleteAllTweetsForCelebrity("John")
	time.Sleep(1 * time.Second)
	johnTweetsAfterDelete, err := dao.GetCelebrityTweets("John")
	if err != nil {
		t.Errorf("Expected err %v to be nil.", err)
	}

	if tweetCount, expected := len(johnTweetsAfterDelete), 0; tweetCount != expected {
		t.Errorf("Expected tweet count for John after delete is %d, but was %d", expected, tweetCount)
	}
}
