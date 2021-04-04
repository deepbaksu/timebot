# timebot

<a href="https://slack.com/oauth/v2/authorize?client_id=288745980535.515443770183&scope=channels:history,channels:join,channels:read,chat:write,chat:write.customize,chat:write.public,commands,dnd:read,emoji:read,groups:history,groups:read,groups:write,im:history,im:read,im:write,links:read,links:write,mpim:history,mpim:read,mpim:write,pins:read,pins:write,reactions:read,reactions:write,reminders:read,reminders:write,remote_files:read,remote_files:share,remote_files:write,team:read,usergroups:read,usergroups:write,users.profile:read,users:read,users:read.email,users:write"><img alt="Add to Slack" height="40" width="139" src="https://platform.slack-edge.com/img/add_to_slack.png" srcset="https://platform.slack-edge.com/img/add_to_slack.png 1x, https://platform.slack-edge.com/img/add_to_slack@2x.png 2x"></a>

[![Build Status](https://travis-ci.com/deepbaksu/timebot.svg?branch=master)](https://travis-ci.com/deepbaksu/timebot)
[![codecov](https://codecov.io/gh/deepbaksu/timebot/branch/master/graph/badge.svg)](https://codecov.io/gh/deepbaksu/timebot)
[![GoDoc](https://godoc.org/github.com/deepbaksu/timebot?status.svg)](https://godoc.org/github.com/deepbaksu/timebot)

Slack 시간 변환 봇

## Example

14:24 PST <--> 07:24 KST

## Run the server

```shell script
go run cmd/server/server.go
```

## Test

```shell script
go test -v ./...
```

## Integration Test

```shell script
docker-compose -f docker-compose.yaml -f docker-compose.test.yaml build
docker-compose -f docker-compose.yaml -f docker-compose.test.yaml run web
```
