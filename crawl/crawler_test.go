package crawl_test

import (
	"github.com/arianht/meantweets/crawl"
	"github.com/arianht/meantweets/database"
	"reflect"
	"testing"
)

func TestCrawlBasicFlowWithEmptyDatabase(t *testing.T) {
	dao := database.DaoMock{Tweets: &[]database.Tweet{}}
	twitter := crawl.TwitterFacadeMock{
		Tweets: map[string][]crawl.Tweet{
			"Johnny": []crawl.Tweet{
				crawl.Tweet{0, "j_one"},
				crawl.Tweet{1, "j_two"},
			},
			"Emma": []crawl.Tweet{
				crawl.Tweet{2, "e_one"},
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
		Dao:         dao,
		Twitter:     twitter,
		Sentiment:   sentiment,
		Celebrities: []string{"Johnny", "Emma"},
	}

	twitterCrawler.Crawl()

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
