package disguard // import "go.zeta.pm/disguard"

import (
	"fmt"
	"net/url"
)

var (
	oauthScope       = url.QueryEscape("identify guilds")
	authorizationURL = "https://discord.com/oauth2/authorize?response_type=code&client_id=%s&scope=%s&redirect_uri=%s"
	tokenURL         = "https://discord.com/api/oauth2/token"
)

// TokenResponse contains response from OAuth
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func (conf *OAuthSection) getRedirectURL() string {
	if conf.redirectURL != nil {
		return *conf.redirectURL
	}
	temp := url.QueryEscape(conf.RedirectURL)
	conf.redirectURL = &temp
	return temp
}

// GetAuthorizationURL returns blah
func (conf *OAuthSection) GetAuthorizationURL() string {
	if conf.authorizationURL != nil {
		return *conf.authorizationURL
	}

	temp := fmt.Sprintf(authorizationURL, conf.ClientID, oauthScope, conf.getRedirectURL())
	conf.authorizationURL = &temp
	return temp
}

// GetTokenURL returns blah
func (conf *OAuthSection) GetTokenURL() string {
	return tokenURL
}
