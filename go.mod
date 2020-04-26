// +heroku goVersion go1.11
// +heroku install ./cmd/...

module github.com/dl4ab/timebot

require (
	github.com/gorilla/mux v1.7.1
	github.com/slack-go/slack v0.6.4
	github.com/stretchr/testify v1.3.0
	go.mongodb.org/mongo-driver v1.3.2
)

go 1.13
