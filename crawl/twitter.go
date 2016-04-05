/*
Package crawl provides a means for crawling for tweets, getting their sentiment score, and writing them
to the database.

To run, the environment variables TWITTER_CONSUMER_KEY and TWITTER_CONSUMER_SECRET
must be set. On Linux, set them using the following commands:

export TWITTER_CONSUMER_KEY=<consumer key>
export TWITTER_CONSUMER_SECRET=<consumer secret>

*/
package crawl

import (
	"fmt"
	"github.com/kurrik/oauth1a"
	"github.com/kurrik/twittergo"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

// TwitterFacade is an interface for retrieving celebrity tweets from Twitter.
type TwitterFacade interface {
	GetTweets(celebrity string, count uint) ([]twittergo.Tweet, error)
}

type twitterFacade struct {
	client TwitterClient
}

// TwitterClient is used to write unit tests for code based on the twittergo package.
type TwitterClient interface {
	SendRequest(req *http.Request) (resp *twittergo.APIResponse, err error)
}

func getTwitterClient() (client TwitterClient, err error) {
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	if consumerKey == "" {
		err = fmt.Errorf("The TWITTER_CONSUMER_KEY env variable must be set.")
		return
	}
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	if consumerSecret == "" {
		err = fmt.Errorf("The TWITTER_CONSUMER_SECRET env variable must be set.")
		return
	}
	config := &oauth1a.ClientConfig{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
	}
	client = twittergo.NewClient(config, nil)

	return
}

func getSearchTweetsRequest(searchQuery string, count uint) (req *http.Request, err error) {
	url := fmt.Sprintf("/1.1/search/tweets.json?%v", getURLValues(searchQuery, count).Encode())
	req, err = http.NewRequest("GET", url, nil)

	return
}

func sendSearchTweetsRequest(client TwitterClient, req *http.Request) (results *twittergo.SearchResults,
	err error) {
	resp, err := client.SendRequest(req)
	if err != nil {
		return
	}

	results = &twittergo.SearchResults{}
	err = resp.Parse(results)

	return
}

func getURLValues(searchQuery string, count uint) (query *url.Values) {
	query = &url.Values{}
	query.Set("q", searchQuery)
	query.Set("lang", "en")
	query.Set("count", strconv.FormatUint(uint64(count), 10))
	query.Set("result_type", "mixed") // Include both popular and real time results in the response

	return
}

// Searches Twitter for and returns up to count tweets that contain the given string. At the time of this writing, Twitter will not return more than 100 results per request.
func (twitter *twitterFacade) GetTweets(celebrity string, count uint) (tweets []twittergo.Tweet, err error) {
	req, err := getSearchTweetsRequest(celebrity, count)
	if err != nil {
		return
	}

	results, err := sendSearchTweetsRequest(twitter.client, req)
	if err != nil {
		return
	}
	tweets = results.Statuses()

	return
}

// NewTwitterFacade returns a Twitter facade that can be used to get tweets based on a query string.
func NewTwitterFacade() (twitter TwitterFacade, err error) {
	facade := &twitterFacade{}
	client, err := getTwitterClient()
	if err != nil {
		return
	}
	facade.client = client
	twitter = facade

	return
}

// NewTwitterFacadeWithClient returns a Twitter facade with a custom client that can be used for unit testing.
func NewTwitterFacadeWithClient(client TwitterClient) TwitterFacade {
	return &twitterFacade{client}
}

// TwitterFacadeMock provides a mock TwitterFacade for unit tests of files that depend on a TwitterFacade.
type TwitterFacadeMock struct {
	Tweets map[string][]twittergo.Tweet
}

// GetTweets implementation for the TwitterFacade mock.
func (twitter TwitterFacadeMock) GetTweets(celebrity string, count uint) ([]twittergo.Tweet, error) {
	return twitter.Tweets[celebrity], nil
}
