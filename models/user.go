package models

type User struct {
	GUID string `json:"guid"`
}

type Auth struct {
	AccessToken string `json:"access"`
	Refresh     string `json:"refresh"`
	LifeTime    string `json:"lifetime"`
	Status      string `json:"status"`
}

type Refresh struct {
	Refresh string `json:"refresh"`
}
