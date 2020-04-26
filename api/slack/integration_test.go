// +build integration

package slack

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"testing"
	"time"
)

func Test_GetTokenFromTeamId(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 10)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	if err != nil {
		t.Fatalf("Failed to connect to a MongoDB for testing. See %v", err)
	}

	app := GetMockApp(nil)
	app.MongoClient = client

	teamID := "T8GMXUUFR"
	accessToken := "xoxb-test"
	mockObject := `{ "_id" : "5ea5f78e3c8fad2abc2d18e4",
"accesstoken" : "xoxb-test",
"tokentype" : "bot",
"scope" : "",
"botuserid" : "UF95LP3NG",
"appid" : "AAP5D1NN5D", 
"team" : { "id" : "T8GMXUUFR", "name" : "DeeplearningForAllBakSu" }, 
"enterprise" : { "id" : "", "name" : "" }, 
"autheduser" : { "id" : "U834MA8P", "scope" : "", "accesstoken" : "", "tokentype" : "" }, 
"slackresponse" : { "ok" : true, "error" : "" } } `
	var mockObjectMap map[string]interface{}
	err = json.Unmarshal([]byte(mockObject), &mockObjectMap)
	if err != nil {
		log.Fatalf("Failed to Unmarshal mockObjectMap. See %v", err)
	}
	deleteMockData(t, app, ctx, teamID)
	_, err = app.GetOauthCollection().UpdateOne(ctx, bson.M{"_id": "5ea5f78e3c8fad2abc2d18e4"}, bson.M{"$set": mockObjectMap}, options.Update().SetUpsert(true))
	if err != nil {
		t.Fatalf("Failed to insert a temporary mock data. See %v", err)
	}

	defer deleteMockData(t, app, ctx, teamID)

	token := GetTokenFromTeamId(&app, teamID)

	if token != accessToken {
		t.Fatalf("Expected %v but Received %v", accessToken, token)
	}
}

func deleteMockData(t *testing.T, app App, ctx context.Context, teamID string) {
	func() {
		_, err := app.GetOauthCollection().DeleteMany(ctx, bson.M{"team.id": teamID})
		if err != nil {
			t.Fatalf("Failed to delete a temporary mock data. See %v", err)
		}
	}()
}
