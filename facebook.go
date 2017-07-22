package fb

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

const (
	GRAPH_BASE_URL         = "https://graph."
	BASE_AUTHORIZATION_URL = "https://www.facebook.com"
	DEFAULT_REDIRECT_URL   = "https://www.facebook.com/connect/login_success.html"
)

type Facebook struct {
	appId       string
	appSecret   string
	version     string
	redirectUrl string
	beta        bool
	timeout     time.Duration
}

func NewFacebook(appId string, appSecret string, version string, redirectUrl string, beta bool, timeout time.Duration) *Facebook {
	return &Facebook{
		appId:       appId,
		appSecret:   appSecret,
		version:     version,
		redirectUrl: redirectUrl,
		beta:        beta,
		timeout:     timeout,
	}
}

//Generate Application Access Token
//retun struct of AccessTokenApp
func (f *Facebook) GetAppAccessToken() (*AccessTokenApp, error) {
	client := &http.Client{}
	var accessTokenApp *AccessTokenApp
	var betaApp string
	if f.beta {
		betaApp = "beta."
	}
	uri := GRAPH_BASE_URL + betaApp + "facebook.com/" + f.version + "/oauth/access_token?client_id=" + f.appId + "&client_secret=" + f.appSecret + "&grant_type=client_credentials"

	response, err := client.Get(uri)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("Error response")
	}

	err = json.NewDecoder(response.Body).Decode(&accessTokenApp)
	if err != nil {
		return nil, errors.New("Error response")
	}

	return accessTokenApp, nil
}

//Generate an app secret proof to sign a request to Graph.
//return string
func (f *Facebook) GetSecretProof(accessToken string) string {
	key := []byte(f.appSecret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(accessToken))
	return hex.EncodeToString(h.Sum(nil))
}
