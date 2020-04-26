package slack

func GetMockApp(httpClient HttpClient) App {
	app := New("SLACK_SIGNING_TOKEN", "BOT_ACCESS_TOKEN", "SLACK_CLIENT_ID", "SLACK_CLIENT_SECRET", httpClient, nil)
	app.TestMode = true
	return app
}

// GetBotMessageForTesting returns a mock bot message (JSON).
func GetBotMessageForTesting() string {
	return `{
  "token": "FWQU4dzFJVJw6mNUTgVgtj5f",
  "team_id": "T8GMXUUFR",
  "api_app_id": "AF5D1NN5D",
  "event": {
    "bot_id": "BF9AJTAGJ",
    "type": "message",
    "text": "2020-03-28 03:08 KST =&gt; 2020-03-27 11:08 PDT",
    "user": "UF95LP3NG",
    "ts": "1585341018.032200",
    "team": "T8GMXUUFR",
    "bot_profile": {
      "id": "BF9AJTAGJ",
      "deleted": false,
      "name": "timebot",
      "updated": 1546979303,
      "app_id": "AF5D1NN5D",
      "icons": {
        "image_36": "https://avatars.slack-edge.com/2019-01-08/520758781318_de293f408c7741c9e443_36.jpg",
        "image_48": "https://avatars.slack-edge.com/2019-01-08/520758781318_de293f408c7741c9e443_48.jpg",
        "image_72": "https://avatars.slack-edge.com/2019-01-08/520758781318_de293f408c7741c9e443_72.jpg"
      },
      "team_id": "T8GMXUUFR"
    },
    "thread_ts": "1585341010.029900",
    "parent_user_id": "U8H34MA8P",
    "channel": "C010KF3RFTN",
    "event_ts": "1585341018.032200",
    "channel_type": "channel"
  },
  "type": "event_callback",
  "event_id": "Ev0110DZDRKR",
  "event_time": 1585341018,
  "authed_users": ["UF95LP3NG"]
}`
}
