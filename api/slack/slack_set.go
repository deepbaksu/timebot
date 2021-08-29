package slack

import (
	"github.com/deepbaksu/timebot/api/slack/security"
	"github.com/google/wire"
	"net/http"
)

var ProdSet = wire.NewSet(
	ProvideSlackApp,
	security.ProdSet,
	wire.Bind(new(SlackRequestVerifier), new(*security.Security)),
	wire.InterfaceValue(new(HttpClient), http.DefaultClient),
)
