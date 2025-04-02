package oauth

import (
	"github.com/jfraska/golang-app/internal/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetGoogleOauthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.Cfg.Oauth.Google.ClientID,
		ClientSecret: config.Cfg.Oauth.Google.ClientSecret,
		RedirectURL:  config.Cfg.Oauth.Google.CallbackURL,
		Scopes:       []string{"profile", "email"},
		Endpoint:     google.Endpoint,
	}
}
