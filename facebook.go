package fb

import (
	"net/http"
	"time"
	"encoding/json"
	"errors"
	"fmt"
)

const (
	GRAPH_BASE_URL = "https://graph.facebook.com"
)

type Facebook struct {
	appId     string
	appSecret string
	version   string
	beta      bool
	timeout   time.Duration
}

func NewFacebook(appId string, appSecret string, version string, beta bool, timeout time.Duration) *Facebook {
	return &Facebook{
		appId:     appId,
		appSecret: appSecret,
		version:   version,
		beta:      beta,
		timeout:   timeout,
	}
}

func (f *Facebook) GetAppAccessToken() (*AccessTokenApp, error) {
	client := &http.Client{}
	var accessTokenApp *AccessTokenApp

	uri := GRAPH_BASE_URL + "/" + f.version + "/oauth/access_token?client_id=" + f.appId + "&client_secret=" + f.appSecret + "&grant_type=client_credentials"
	response, err := client.Get(uri)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("Error response")
	}

	err = json.NewDecoder(response.Body).Decode(&accessTokenApp)
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("Error response")
	}

	return accessTokenApp, nil
}
