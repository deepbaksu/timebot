package slack

// App manages Global Application State
type App struct {
	SigningToken string
}

// New creates a slack API application
func New(slackSigningToken string) App {
	return App{slackSigningToken}
}
