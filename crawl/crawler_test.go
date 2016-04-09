package crawl

import (
	"reflect"
	"testing"

	"github.com/arianht/meantweets/database"
	"github.com/kurrik/twittergo"
)

func TestCrawl(t *testing.T) {
	testCases := []struct {
		initialDatabase     []database.Tweet
		celebrities         []string
		maxTweets           int
		tweets              map[string][]twittergo.Tweet
		scores              map[string]int32
		expectedEndDatabase []database.Tweet
	}{
		{ // Test Case: Empty Initial Database
			initialDatabase: []database.Tweet{},
			celebrities:     []string{"Johnny", "Emma"},
			maxTweets:       2,
			tweets: map[string][]twittergo.Tweet{
				"Johnny": []twittergo.Tweet{
					twittergo.Tweet{"id_str": "0", "text": "j_one"},
					twittergo.Tweet{"id_str": "1", "text": "j_two"},
				},
				"Emma": []twittergo.Tweet{
					twittergo.Tweet{"id_str": "2", "text": "e_one"},
				},
			},
			scores: map[string]int32{
				"j_one": 2,
				"j_two": 100,
				"e_one": 872,
			},
			expectedEndDatabase: []database.Tweet{
				database.Tweet{
					CelebrityName: "Johnny",
					Id:            0,
					Score:         2,
				},
				database.Tweet{
					CelebrityName: "Johnny",
					Id:            1,
					Score:         100,
				},
				database.Tweet{
					CelebrityName: "Emma",
					Id:            2,
					Score:         872,
				},
			},
		},
		{ // Test Case: Too many tweets for Johnny
			initialDatabase: []database.Tweet{
				database.Tweet{
					CelebrityName: "Johnny",
					Id:            3,
					Score:         300,
				},
			},
			celebrities: []string{"Johnny", "Emma"},
			maxTweets:   2,
			tweets: map[string][]twittergo.Tweet{
				"Johnny": []twittergo.Tweet{
					twittergo.Tweet{"id_str": "0", "text": "j_one"},
					twittergo.Tweet{"id_str": "1", "text": "j_two"},
				},
				"Emma": []twittergo.Tweet{
					twittergo.Tweet{"id_str": "2", "text": "e_one"},
				},
			},
			scores: map[string]int32{
				"j_one": 2,
				"j_two": 100,
				"e_one": 872,
			},
			expectedEndDatabase: []database.Tweet{
				database.Tweet{
					CelebrityName: "Johnny",
					Id:            0,
					Score:         2,
				},
				database.Tweet{
					CelebrityName: "Johnny",
					Id:            1,
					Score:         100,
				},
				database.Tweet{
					CelebrityName: "Emma",
					Id:            2,
					Score:         872,
				},
			},
		},
	}

	for _, testCase := range testCases {
		dao := database.DaoMock{Tweets: &testCase.initialDatabase}
		twitter := TwitterFacadeMock{Tweets: testCase.tweets}
		sentiment := SentimentAnalyzerMock{Scores: testCase.scores}
		twitterCrawler := TwitterCrawler{
			Dao:       dao,
			Twitter:   twitter,
			Sentiment: sentiment,
		}

		twitterCrawler.Crawl(testCase.celebrities, testCase.maxTweets)

		if !reflect.DeepEqual(*dao.Tweets, testCase.expectedEndDatabase) {
			t.Errorf("Expected tweet database is %v, but was %v", testCase.expectedEndDatabase,
				*dao.Tweets)
		}
	}
}
