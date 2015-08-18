#!/bin/bash
go run mocksteam/server.go &
killfinish()
{
	ps ax | grep server | grep go | gawk '{print $1}' | xargs kill -9
}
trap ' killfinish ' INT
trap ' killfinish ' EXIT
cd auth
go test