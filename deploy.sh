#!/bin/bash
cd server
go get google.golang.org/appengine/cmd/aedeploy
aedeploy gcloud app deploy
