package slack

import (
	"context"
	"fmt"
	slackApi "github.com/slack-go/slack"
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

	// slackApi.GetOAuthV2Response(app.httpClient)
	oAuthV2Response, err := slackApi.GetOAuthV2Response(app.HttpClient, app.SlackClientId, app.SlackClientSecret, code[0], "")
	if err != nil {
		log.Fatalf("OauthV2Response has failed. See %v", err)
	}

	log.Printf("Received token => %+v", oAuthV2Response)

	client := slackApi.New(oAuthV2Response.AccessToken)
	teamInfo, err := client.GetTeamInfo()
	if err != nil {
		log.Fatalf("Failed to get team info %v", err)
	}
	log.Printf("Received a teamInfo => %+v", teamInfo)

	http.Redirect(writer, request, fmt.Sprintf("https://%v.slack.com", teamInfo.Domain), http.StatusTemporaryRedirect)

	// app.MongoClient
	oauthCollection := app.GetOauthCollection()
	oAuthV2ResponseInBson, err := bson.Marshal(oAuthV2Response)
	if err != nil {
		log.Fatalf("Failed to Marshal oAuthV2Response(%+v). See %v", oAuthV2Response, err)
	}
	_, err = oauthCollection.UpdateOne(context.Background(), bson.M{"_id": teamInfo.ID}, bson.M{"$set": oAuthV2ResponseInBson}, options.Update().SetUpsert(true))
	if err != nil {
		log.Fatalf("Failed to upsert OAuthInformation => %v", err)
	}
}

