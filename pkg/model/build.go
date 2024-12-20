package model

type GlobalOpt struct {
	Auth *Auth `json:"auth"`
}

type Auth struct {
	APIKey       string `json:"apiKey"`
	APISecret    string `json:"apiSecret"`
	AccessToken  string `json:"accessToken"`
	AccessSecret string `json:"accessSecret"`
}

type BuildInfo struct {
	Version string `json:"version"`
}
