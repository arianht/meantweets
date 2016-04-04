package crawl

// SentimentAnalyzer defines an interface for retrieving the sentiment score of a tweet.
type SentimentAnalyzer interface {
	GetScoreForTweet(tweet string) int32
}

// SentimentAnalyzerMock provides a mock SentimentAnalyzer for unit tests of files that depend on a SentimentAnalyzer.
type SentimentAnalyzerMock struct {
	Scores map[string]int32
}

// GetScoreForTweet returns the sentiment score for a given tweet.
func (analyzer SentimentAnalyzerMock) GetScoreForTweet(tweet string) int32 {
	return analyzer.Scores[tweet]
}
