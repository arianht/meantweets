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
	johnTweetOneNoName := Tweet{Id: 0, Score: 500}
	johnTweetTwo := Tweet{"John", 1, 12}
	johnTweetTwoNoName := Tweet{Id: 1, Score: 12}
	jenTweetOne := Tweet{"Jen", 2, 750}
	jenTweetOneNoName := Tweet{Id: 2, Score: 750}
	jenTweetTwo := Tweet{"Jen", 3, 500}
	jenTweetTwoNoName := Tweet{Id: 3, Score: 500}
	jenTweetThree := Tweet{"Jen", 4, 300}
	jenTweetThreeNoName := Tweet{Id: 4, Score: 300}
	jenTweetFour := Tweet{"Jen", 5, 751}
	jenTweetFourNoName := Tweet{Id: 5, Score: 751}

	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	dao := DatastoreDao{ctx}
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
	time.Sleep(2 * time.Second)

	johnTweets, err := dao.GetCelebrityTweets("John")
	if err != nil {
		t.Errorf("Expected err %v to be nil.", err)
	}

	if tweetCount, expected := len(johnTweets), 2; tweetCount != expected {
		t.Errorf("Expected tweet count for John is %d, but was %d", expected, tweetCount)
	}
	if tweets, expected := johnTweets, []Tweet{johnTweetTwoNoName, johnTweetOneNoName}; !reflect.DeepEqual(tweets, expected) {
		t.Errorf("Expected tweet content for John is %v, but was %v", expected, tweets)
	}

	jenTweets, err := dao.GetCelebrityTweets("Jen")
	if err != nil {
		t.Errorf("Expected err %v to be nil.", err)
	}

	if tweetCount, expected := len(jenTweets), 4; tweetCount != expected {
		t.Errorf("Expected tweet count for Jen is %d, but was %d", expected, tweetCount)
	}

	if tweets, expected := jenTweets, []Tweet{jenTweetThreeNoName, jenTweetTwoNoName, jenTweetOneNoName, jenTweetFourNoName}; !reflect.DeepEqual(tweets, expected) {
		t.Errorf("Expected tweet content for Jen is %v, but was %v", expected, tweets)
	}

	dao.DeleteAllTweetsForCelebrity("John")
	time.Sleep(2 * time.Second)
	johnTweetsAfterDelete, err := dao.GetCelebrityTweets("John")
	if err != nil {
		t.Errorf("Expected err %v to be nil.", err)
	}

	if tweetCount, expected := len(johnTweetsAfterDelete), 0; tweetCount != expected {
		t.Errorf("Expected tweet count for John after delete is %d, but was %d", expected, tweetCount)
	}
}
