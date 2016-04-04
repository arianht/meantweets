/*
Package server provides the main web server for Mean Tweets. Requires App Engine.

Run locally with:
goapp serve github.com/arianht/meantweets/server

It will automatically detect changes in the code.

Push to the cloud with:
appcfg.py A <project_id> -V <version> update github.com/arianht/meantweets/server
*/
package server

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprint(writer, "Mean Tweets basic server!")
}
