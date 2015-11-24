#!/bin/sh

go test
gocov test | gocov report

for pkg in auth inventory market tradeoffer user;do
	cd $pkg
	go test
	gocov test | gocov report
	cd ..
done
