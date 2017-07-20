package fb

import (
	"net/http"
)

type Facebook struct {
	appId     string
	appSecret string
	version   string
	beta      bool
}

func NewFacebook(appId string, appSecret string, version string, beta bool) *Facebook {
	return &Facebook{
		appId:     appId,
		appSecret: appSecret,
		version:   version,
		beta:      beta,
	}
}

//https://graph.facebook.com/v2.10/me?fields=id,name,picture{height,is_silhouette,url,width}&access_token=spsB4NEZD

func (f *Facebook) GetAccessToken(path string) {

}
