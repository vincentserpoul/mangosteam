#!/bin/sh

for pkg in auth inventory tradeoffer user;do
	cd $pkg
	go test
	gocov test | gocov report
	cd ..
done