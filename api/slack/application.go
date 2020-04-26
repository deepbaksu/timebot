package slack

// App manages Global Application State
type App struct {
	SigningToken        string
	BotOAuthAccessToken string

	SlackClientId string
	SlackClientSecret string
	// if it's true, the request verification does not happen
	TestMode bool
}

// New creates a slack API application
func New(slackSigningToken, botOAuthAccessToken, slackClientId, slackClientSecret string) App {
	return App{
		SigningToken:        slackSigningToken,
		BotOAuthAccessToken: botOAuthAccessToken,

		SlackClientId: slackClientId,
		SlackClientSecret: slackClientSecret,

		TestMode:            false,
	}
}
