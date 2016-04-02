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
	WriteCelebrityTweet(tweet Tweet)
	GetCelebrityTweets(celebrityName string) (tweets []Tweet)
	DeleteAllTweetsForCelebrity(celebrityName string)
}

// A DAO for interacting with App Engine's Datastore.
type DatastoreDao struct {
	Ctx context.Context // App Engine Context which can be obtained from an HTTP request.
}

// A tweet entity for storing data in the datastore.
type Tweet struct {
	CelebrityName string
	Id            int64
	Score         int32
}

// Saves the celebrityName, tweetContents pair to the datastore. Note that duplicates aren't
// caught here because of writing asynchronicity.
func (datastoreDao DatastoreDao) WriteCelebrityTweet(tweet Tweet) {
	key := datastore.NewIncompleteKey(datastoreDao.Ctx, datastoreKind, nil)
	if _, err := datastore.Put(datastoreDao.Ctx, key, &tweet); err != nil {
		fmt.Printf("Error writing to database: %v\n", err)
		return
	}
}

// Retrieves all the celebrity tweets related to a celebrity sorted with highest scores first.
// Duplicate tweets will be filtered out.
func (datastoreDao DatastoreDao) GetCelebrityTweets(celebrityName string) (tweets []Tweet) {
	q := datastore.NewQuery(datastoreKind).
		Filter("CelebrityName = ", celebrityName).
		Order("-Score")
	var results []Tweet
	if _, err := q.GetAll(datastoreDao.Ctx, &results); err != nil {
		fmt.Printf("Error reading the database: %v\n", err)
		return
	}
	tweetSet := make(map[int64]bool)
	for _, tweetResult := range results {
		if !tweetSet[tweetResult.Id] {
			tweets = append(tweets, tweetResult)
			tweetSet[tweetResult.Id] = true
		}
	}
	return
}

// Deletes all tweets for a provided celebirty name.
func (datastoreDao DatastoreDao) DeleteAllTweetsForCelebrity(celebrityName string) {
	q := datastore.NewQuery(datastoreKind).
		Filter("CelebrityName = ", celebrityName).
		KeysOnly()
	keys, err := q.GetAll(datastoreDao.Ctx, nil)
	if err != nil {
		fmt.Printf("Error reading the database: %v\n", err)
		return
	}
	datastore.DeleteMulti(datastoreDao.Ctx, keys)
}
