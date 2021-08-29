package slack

import (
	"context"
	"fmt"
	slackAPI "github.com/slack-go/slack"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func (app *App) OauthHandler(writer http.ResponseWriter, request *http.Request) {
	code, ok := request.URL.Query()["code"]

	if !ok || len(code) != 1 {
		log.Fatalf("The URL(%v) does not contain ?code=XXXXX", request.URL.String())
	}

	oAuthV2Response, err := slackAPI.GetOAuthV2Response(app.HttpClient, app.config.SlackClientID, app.config.SlackClientSecret, code[0], "")
	if err != nil {
		log.Fatalf("OauthV2Response has failed. See %v", err)
	}

	log.Printf("Received token => %+v", oAuthV2Response)

	client := slackAPI.New(oAuthV2Response.AccessToken)
	teamInfo, err := client.GetTeamInfo()
	if err != nil {
		log.Fatalf("Failed to get team info %v", err)
	}
	log.Printf("Received a teamInfo => %+v", teamInfo)

	http.Redirect(writer, request, fmt.Sprintf("https://%v.slack.com", teamInfo.Domain), http.StatusTemporaryRedirect)

	// app.MongoClient
	oauthCollection := app.GetOauthCollection()
	_, err = oauthCollection.UpdateOne(context.Background(), bson.M{"_id": teamInfo.ID}, bson.M{"$set": oAuthV2Response}, options.Update().SetUpsert(true))
	if err != nil {
		log.Fatalf("Failed to upsert OAuthInformation => %v", err)
	}
}
