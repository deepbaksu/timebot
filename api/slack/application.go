package slack

// App manages Global Application State
type App struct {
	SigningToken string
	// if it's true, the request verification does not happen
	TestMode bool
}

// New creates a slack API application
func New(slackSigningToken string) App {
	return App{slackSigningToken, false}
}
