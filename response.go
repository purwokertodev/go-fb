package fb

type AccessTokenApp struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type ProfileResponse struct {
	About string `json:"about"`
	Cover struct {
		ID      string `json:"id"`
		OffsetX int    `json:"offset_x"`
		OffsetY int    `json:"offset_y"`
		Source  string `json:"source"`
	} `json:"cover"`
	Education []struct {
		Concentration []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"concentration"`
		School struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"school"`
		Type string `json:"type"`
		Year struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"year"`
		ID string `json:"id"`
	} `json:"education"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Name       string `json:"name"`
	NameFormat string `json:"name_format"`
	ShortName  string `json:"short_name"`
	Gender     string `json:"gender"`
	Picture    struct {
		Data struct {
			IsSilhouette bool   `json:"is_silhouette"`
			URL          string `json:"url"`
		} `json:"data"`
	} `json:"picture"`
	ID string `json:"id"`
}
