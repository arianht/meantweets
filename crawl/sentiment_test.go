package crawl

import (
	"testing"
)

func TestGetScoreForTweet(t *testing.T) {
	testCases := []struct {
		expectedScore int32
		tweet         string
	}{
		{
			expectedScore: 11,
			tweet:         "Your mother is an awful lady",
		},
		{
			expectedScore: 0,
			tweet:         "You're a terrible person. You shouldn't exist.",
		},
		{
			expectedScore: 66,
			tweet:         "Wow! You are the most amazing person I've ever met!",
		},
		{
			expectedScore: 44,
			tweet:         "Katy Perry is a great person!",
		},
	}

	for _, testCase := range testCases {
		analyzer, err := NewSentimentAnalyzer()
		if err != nil {
			t.Fatalf("Expected error %v to be nil", err)
		}
		if score := analyzer.GetScoreForTweet(testCase.tweet); score != testCase.expectedScore {
			t.Errorf("Expected sentiment score for tweet %v to be %d, but was %d",
				testCase.tweet, testCase.expectedScore, score)
		}
	}
}
