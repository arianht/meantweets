package crawl_test

import (
	"github.com/arianht/meantweets/crawl"
	"github.com/arianht/meantweets/database"
	"github.com/kurrik/twittergo"
	"reflect"
	"testing"
)

func TestCrawlBasicFlowWithEmptyDatabase(t *testing.T) {
	dao := database.DaoMock{Tweets: &[]database.Tweet{}}
	twitter := crawl.TwitterFacadeMock{
		Tweets: map[string][]twittergo.Tweet{
			"Johnny": []twittergo.Tweet{
				twittergo.Tweet{"id_str": "0", "text": "j_one"},
				twittergo.Tweet{"id_str": "1", "text": "j_two"},
			},
			"Emma": []twittergo.Tweet{
				twittergo.Tweet{"id_str": "2", "text": "e_one"},
			},
		},
	}
	sentiment := crawl.SentimentAnalyzerMock{
		Scores: map[string]int32{
			"j_one": 2,
			"j_two": 100,
			"e_one": 872,
		},
	}
	twitterCrawler := crawl.TwitterCrawler{
		Dao:       dao,
		Twitter:   twitter,
		Sentiment: sentiment,
	}

	twitterCrawler.Crawl([]string{"Johnny", "Emma"})

	expectedTweetsDatabase := []database.Tweet{
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
	}
	if !reflect.DeepEqual(*dao.Tweets, expectedTweetsDatabase) {
		t.Errorf("Expected tweet database is %v, but was %v", expectedTweetsDatabase, *dao.Tweets)
	}
}
