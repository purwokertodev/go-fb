package fb

type Facebook struct {
	appId     string
	appSecret string
	version   string
}

func NewFacebook(appId string, appSecret string, version string) *Facebook {
	return &Facebook{
		appId:     appId,
		appSecret: appSecret,
		version:   version,
	}
}
