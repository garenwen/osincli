package osincli

type ClientConfig struct {
	ClientId                 string
	ClientSecret             string
	AuthorizeUrl             string
	TokenUrl                 string
	RedirectUrl              string
	Scope                    string
	ErrorsInStatusCode       bool
	SendClientSecretInParams bool
	UseGetAccessRequest      bool

	// PKCE / RFC7636 fields
	CodeChallenge       string
	CodeChallengeMethod string
	CodeVerifier        string

	//wechat
	Wechat WechatClientConfig
}

type WechatClientConfig struct {
	Appid string
	Secret string
}
