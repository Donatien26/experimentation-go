package utils

import (
	"context"
	"net/http"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

// GenerateOauthConfig ... This func must be Exported, Capitalized, and comment added.
func GenerateOauthConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     viper.GetString("auth.conf.ClientID"),
		ClientSecret: "",
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  viper.GetString("auth.conf.authURL"),
			TokenURL: viper.GetString("auth.conf.tokenURL"),
		}}
	return conf
}

// GetLastToken ...
func GetLastToken() *oauth2.Token {
	token := new(oauth2.Token)
	token.AccessToken = viper.GetString("auth.accessToken")
	token.RefreshToken = viper.GetString("auth.refreshToken")
	token.Expiry = viper.GetTime("auth.expiry")
	token.TokenType = viper.GetString("tokenType")
	return token
}

// GetAuthClient ...
func GetAuthClient() *http.Client {
	conf := GenerateOauthConfig()
	ctx := context.Background()
	token := GetLastToken()
	return conf.Client(ctx, token)
}

//SaveToken ...
func SaveToken(token *oauth2.Token) {
	viper.Set("auth.accessToken", token.AccessToken)
	viper.Set("auth.expire", token.Expiry)
	viper.Set("auth.refreshToken", token.RefreshToken)
	viper.Set("auth.tokenType", token.TokenType)
	viper.WriteConfig()
 }
