package crawl

import "github.com/cdipaolo/sentiment"

// SentimentAnalyzer defines an interface for retrieving the sentiment score of a tweet.
type SentimentAnalyzer interface {
	GetScoreForTweet(tweet string) int32
}

type simpleSentimentAnalyzer struct {
	model sentiment.Models
}

func calculateSentimentScore(analysis *sentiment.Analysis) int32 {
	numOfWords := len(analysis.Words)
	numOfSentences := len(analysis.Sentences)
	totalAvg := 0.0

	// Calculate average score of words.
	if numOfWords > 0 {
		wordsAvg := 0.0
		for _, word := range analysis.Words {
			wordsAvg += float64(word.Score)
		}
		wordsAvg /= float64(numOfWords)
		totalAvg += wordsAvg
	}

	// Calculate average score of sentences.
	if numOfSentences > 0 {
		sentencesAvg := 0.0
		for _, sentence := range analysis.Sentences {
			sentencesAvg += float64(sentence.Score)
		}
		sentencesAvg /= float64(len(analysis.Sentences))
		totalAvg += sentencesAvg
	}

	totalAvg += float64(analysis.Score)

	// Calculate total average.
	totalAvg /= 3.0

	return int32(totalAvg * 100)
}

// GetScoreForTweet returns the sentiment score for a tweet using the simple
// sentiment analyzer. Smaller scores indicate meaner tweets.
func (analyzer *simpleSentimentAnalyzer) GetScoreForTweet(tweet string) int32 {
	analysis := analyzer.model.SentimentAnalysis(tweet, sentiment.English)
	return calculateSentimentScore(analysis)
}

// NewSentimentAnalyzer returns a new simple sentiment analyzer.
func NewSentimentAnalyzer() (analyzer SentimentAnalyzer, err error) {
	model, err := sentiment.Train()
	if err != nil {
		return
	}
	analyzer = &simpleSentimentAnalyzer{model}
	return
}

// SentimentAnalyzerMock provides a mock SentimentAnalyzer for unit tests of files
// that depend on a SentimentAnalyzer.
type SentimentAnalyzerMock struct {
	Scores map[string]int32
}

// GetScoreForTweet returns the sentiment score for a given tweet.
func (analyzer SentimentAnalyzerMock) GetScoreForTweet(tweet string) int32 {
	return analyzer.Scores[tweet]
}
