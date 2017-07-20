package fb

import (
	"net/http"
	"time"
	"io/ioutil"
	"errors"
)

const(
	GRAPH_BASE_URL = "https://graph.facebook.com"
)

type Facebook struct {
	appId     string
	appSecret string
	version   string
	beta      bool
	timeout	time.Duration
}

func NewFacebook(appId string, appSecret string, version string, beta bool, timeout time.Duration) *Facebook {
	return &Facebook{
		appId:     appId,
		appSecret: appSecret,
		version:   version,
		beta:      beta,
		timeout: timeout,
	}
}

//https://graph.facebook.com/v2.10/me?fields=id,name,picture{height,is_silhouette,url,width}&access_token=spsB4NEZD

func (f *Facebook) GetAppAccessToken() (interface{}, error) {
	client := &http.Client{
		Timeout: f.timeout,
	}

	uri := GRAPH_BASE_URL+"/"+f.version+"/oauth/access_token?client_id="+f.appId+"&client_secret="+f.appSecret+"&grant_type=client_credentials"
	response, err := client.Get(uri)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("Error response")
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("Error response")
	}
	bodyString := string(bodyBytes)

	///oauth/access_token?client_id={app-id}&client_secret={app-secret}&grant_type=client_credentials

	return struct{
		AccessToken string
	}{
		AccessToken: bodyString,
	}, nil
}
