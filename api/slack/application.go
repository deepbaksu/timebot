package slack

// App manages Global Application State
type App struct {
	SigningToken        string
	BotOAuthAccessToken string
	// if it's true, the request verification does not happen
	TestMode bool
}

// New creates a slack API application
func New(slackSigningToken, botOAuthAccessToken string) App {
	return App{
		SigningToken:        slackSigningToken,
		BotOAuthAccessToken: botOAuthAccessToken,
		TestMode:            false,
	}
}
