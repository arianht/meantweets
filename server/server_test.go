package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/arianht/meantweets/database"
	"github.com/arianht/meantweets/util"
	"golang.org/x/net/context"
)

func TestGetTweetsEndpoint(t *testing.T) {
	ctx := context.Background()

	tweets := []database.Tweet{
		database.Tweet{"test", 0, 500},
		database.Tweet{"test", 1, 12},
	}
	tweetsWithoutName := []database.Tweet{
		database.Tweet{Id: 1, Score: 12},
		database.Tweet{Id: 0, Score: 500},
	}

	dao, err := database.NewDatastoreDao(ctx)
	if err != nil {
		t.Errorf("Failed to create datastore dao, %v", err)
	}
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
	ctx := context.Background()

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
	_, err := util.GetTwitterAPICredentialsFromFile("credentials.json")
	if err != nil {
		t.Skipf("Error getting Twitter credentials: %v. Skipping crawl endpoint test.", err)
	}

	ctx := context.Background()

	request, _ := http.NewRequest("GET", "/crawl", nil)
	recorder := httptest.NewRecorder()
	crawlHandler(ctx, recorder, request)

	if response := recorder.Body.String(); response != "Successfully crawled." {
		t.Errorf("Expected body to be OK, but was %v", response)
	}
}

func TestRootEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	recorder := httptest.NewRecorder()
	rootHandler(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status to be OK, but was %v", recorder.Code)
	}
}

func TestRandomEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/random", nil)
	recorder := httptest.NewRecorder()
	rootHandler(recorder, request)

	if recorder.Code != http.StatusSeeOther {
		t.Errorf("Expected status to be SeeOther, but was %v", recorder.Code)
	}
}
