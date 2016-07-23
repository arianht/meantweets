package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/arianht/meantweets/database"
	"github.com/arianht/meantweets/util"
	"google.golang.org/appengine/aetest"
)

func TestGetTweetsEndpoint(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatalf("Could not get a context - %v", err)
	}
	defer done()

	tweets := []database.Tweet{
		database.Tweet{"test", 0, 500},
		database.Tweet{"test", 1, 12},
	}
	tweetsWithoutName := []database.Tweet{
		database.Tweet{Id: 1, Score: 12},
		database.Tweet{Id: 0, Score: 500},
	}

	dao := database.DatastoreDao{ctx}
	dao.WriteCelebrityTweets(tweets)

	// Sadly, App Engine Datastore takes time to fully write. Without this sleep,
	// the write won't be done in time for the read.
	time.Sleep(2 * time.Second)

	request, _ := http.NewRequest("GET", "/get_tweets?celebrity=test", nil)
	recorder := httptest.NewRecorder()
	tweetsHandler(ctx, recorder, request)

	expectedResponse, _ := json.Marshal(tweetsWithoutName)
	if response := recorder.Body.String(); response != string(expectedResponse) {
		t.Errorf("Expected body to be %v, but was %v", string(expectedResponse), string(response))
	}
}

func TestGetCelebritiesEndpoint(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatalf("Could not get a context - %v", err)
	}
	defer done()

	request, _ := http.NewRequest("GET", "/get_celebrities", nil)
	recorder := httptest.NewRecorder()
	celebritiesHandler(ctx, recorder, request)

	expectedResponse, _ := json.Marshal(celebrities)
	if response := recorder.Body.String(); response != string(expectedResponse) {
		t.Errorf("Expected body to be %v, but was %v", string(expectedResponse), string(response))
	}
}

func TestCrawlEndpoint(t *testing.T) {
	// Skip test if credentials are not found.
	_, err := util.GetTwitterAPICredentialsFromFile("../credentials.json")
	if err != nil {
		t.Skipf("Error getting Twitter credentials: %v. Skipping crawl endpoint test.", err)
	}

	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatalf("Could not get a context - %v", err)
	}
	defer done()

	request, _ := http.NewRequest("GET", "/crawl", nil)
	recorder := httptest.NewRecorder()
	crawlHandler(ctx, recorder, request)

	if response := recorder.Body.String(); response != "Successfully crawled." {
		t.Errorf("Expected body to be OK, but was %v", response)
	}
}
