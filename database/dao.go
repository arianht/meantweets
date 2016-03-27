/*
Package for reading and writing from the database. Perform all database
operations through this package.
*/
package database

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

const (
	datastoreKind = "Tweet"
)

// The data access object that abstracts database interactions.
type Dao interface {
	WriteCelebrityTweet(celebrtityName string, tweetContents string)
	GetCelebrityTweets(celebrtityName string) (tweets []string)
}

// A DAO for interacting with App Engine's Datastore.
type DatastoreDao struct {
	Ctx context.Context // App Engine Context which can be obtained from an HTTP request.
}

// A tweet entity for storing data in the datastore.
type tweet struct {
	CelebrityName string
	Contents      string
}

// Saves the celebrityName, tweetContents pair to the datastore. Note that duplicates aren't
// caught here because of writing asynchronicity.
func (datastoreDao DatastoreDao) WriteCelebrityTweet(celebrtityName string, tweetContents string) {
	cebTweet := &tweet{celebrtityName, tweetContents}
	key := datastore.NewIncompleteKey(datastoreDao.Ctx, datastoreKind, nil)
	if _, err := datastore.Put(datastoreDao.Ctx, key, cebTweet); err != nil {
		fmt.Printf("Error writing to database: %v\n", err)
		return
	}
}

// Retrieves all the celebrity tweets related to a celebrity. Duplicate tweets
// will be filtered out.
func (datastoreDao DatastoreDao) GetCelebrityTweets(celebrtityName string) (tweets []string) {
	q := datastore.NewQuery(datastoreKind).Filter("CelebrityName = ", celebrtityName)
	var results []tweet
	if _, err := q.GetAll(datastoreDao.Ctx, &results); err != nil {
		fmt.Printf("Error reading the database: %v\n", err)
		return
	}
	tweetSet := make(map[string]bool)
	for _, tweetResult := range results {
		if !tweetSet[tweetResult.Contents] {
			tweets = append(tweets, tweetResult.Contents)
			tweetSet[tweetResult.Contents] = true
		}
	}
	return
}
