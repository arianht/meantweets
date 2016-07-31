#!/bin/bash
cd ui
npm i
npm run build:prod
cp -r dist ../server/
cd ..
