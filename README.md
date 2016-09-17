# Mean Tweets about Celebrities
[![Build Status](https://travis-ci.org/arianht/meantweets.svg?branch=master)](https://travis-ci.org/arianht/meantweets)

meantweets crawls Twitter and brings you the meanest tweets about celebrities. Inspired by Jimmy Kimmel's *Celebrities Read Mean Tweets* [series](https://www.youtube.com/playlist?list=PLs4hTtftqnlAkiQNdWn6bbKUr-P1wuSm0).

## Setting up Development Environment

### Get the Repository
`go get github.com/arianht/meantweets`

### Setting up App Engine and Datastore
TODO

### Run the tests
1. Open up two terminals in root directory of the project
2. In one terminal, run `./start_db.sh`
3. In the other, run `source export_datastore_env_vars.sh` and `./run_all_tests.sh`

### Run locally
1. Open up two terminals in root directory of the project
2. In one terminal, run `./start_db.sh`
3. In the other, run `source export_datastore_env_vars.sh`, `./build_ui.sh`, `cd server`, and `go run server.go`
