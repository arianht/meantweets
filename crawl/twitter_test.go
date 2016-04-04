package crawl_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/arianht/meantweets/crawl"
	"github.com/kurrik/twittergo"
)

type TwitterClientMock struct {
	responses map[string]*twittergo.APIResponse
}

func (twitterClient *TwitterClientMock) SendRequest(req *http.Request) (resp *twittergo.APIResponse, err error) {
	resp, exists := twitterClient.responses[req.URL.String()]
	if !exists {
		fmt.Printf("No response for request %v with URI %v\n", req, req.URL.String())
	}

	return
}

func TestGetTweets(t *testing.T) {
	expectedReqURL := "/1.1/search/tweets.json?count=100&lang=en&q=fake+name&result_type=mixed"
	expectedTweets := []twittergo.Tweet{
		twittergo.Tweet{
			"id_str": "716102930609209345",
			"text":   "this is a test tweet",
		},
		twittergo.Tweet{
			"id_str": "id_str:711700393533640704",
			"text":   "this is another test tweet",
		},
	}
	expectedTweetsJSON, _ := json.Marshal(&twittergo.SearchResults{"statuses": expectedTweets})
	expectedResp := &twittergo.APIResponse{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewReader(expectedTweetsJSON)),
	}
	twitterClient := TwitterClientMock{map[string]*twittergo.APIResponse{expectedReqURL: expectedResp}}

	twitterFacade := crawl.NewTwitterFacadeWithClient(&twitterClient)
	tweets, err := twitterFacade.GetTweets("fake name")

	if err != nil {
		t.Fatalf("Expected nil error, but was %v", err)
	}
	if !reflect.DeepEqual(tweets, expectedTweets) {
		t.Errorf("Expected tweets are %v, but were %v", expectedTweets, tweets)
	}
}
