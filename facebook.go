package fb

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"strings"
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
	accessToken string
	version     string
	redirectUrl string
	beta        bool
	timeout     time.Duration
	client      facebookClient
}

func NewFacebook(appId, appSecret, accessToken, version, redirectUrl string, beta bool, timeout time.Duration) *Facebook {
	return &Facebook{
		appId:       appId,
		appSecret:   appSecret,
		accessToken: accessToken,
		version:     version,
		redirectUrl: redirectUrl,
		beta:        beta,
		timeout:     timeout,
		client:      NewFacebookClient(),
	}
}

func (f *Facebook) call(method, path string, body io.Reader, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = path + "/"
	}
	var betaApp string
	if f.beta {
		betaApp = "beta."
	}
	path = GRAPH_BASE_URL + betaApp + "facebook.com/" + f.version + path
	return f.client.call(method, path, body, v)
}

//Generate Application Access Token
//retun struct of AccessTokenApp
func (f *Facebook) GetAppAccessToken() (*AccessTokenApp, error) {
	var accessTokenApp *AccessTokenApp
	path := "/oauth/access_token?client_id=" + f.appId + "&client_secret=" + f.appSecret + "&grant_type=client_credentials"
	err := f.call("GET", path, nil, &accessTokenApp)
	if err != nil {
		return nil, err
	}
	return accessTokenApp, nil
}

//Generate an app secret proof to sign a request to Graph.
//return string
func (f *Facebook) getSecretProof() string {
	key := []byte(f.appSecret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(f.accessToken))
	return hex.EncodeToString(h.Sum(nil))
}

func (f *Facebook) GetUserProfile(scope string) (map[string]interface{}, error) {
	var result interface{}
	appSecretProof := f.getSecretProof()

	path := "/me?fields=" + scope + "&access_token=" + f.accessToken + "&appsecret_proof=" + appSecretProof

	err := f.call("GET", path, nil, &result)
	if err != nil {
		return nil, errors.New("Error response")
	}
	p, ok := result.(map[string]interface{})
	if !ok {
		return nil, errors.New("Error response")
	}

	return p, nil
}
