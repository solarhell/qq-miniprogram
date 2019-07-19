package swan_miniprogram

import "net/url"

const (
	// BaseURL 基础URL
	baseURL  = "https://spapi.baidu.com"
	codeAPI  = "/oauth/jscode2sessionkey"
	tokenAPI = "https://openapi.baidu.com/oauth/2.0/token"
)

func CodeToURL(appKey, appSecret, code string) (s string, err error) {
	if appKey == "" || appSecret == "" || code == "" {
		return s, ErrNotAllowEmptyParam
	}
	u, err := url.Parse(baseURL + codeAPI)
	if err != nil {
		return s, err
	}

	query := u.Query()

	query.Set("client_id", appKey)
	query.Set("sk", appSecret)
	query.Set("code", code)

	u.RawQuery = query.Encode()

	return u.String(), nil
}

func TokenURL(appKey, appSecret string) (s string, err error) {
	if appKey == "" || appSecret == "" {
		return s, ErrNotAllowEmptyParam
	}
	u, err := url.Parse(tokenAPI)
	if err != nil {
		return s, err
	}

	query := u.Query()

	query.Set("client_id", appKey)
	query.Set("client_secret", appSecret)
	query.Set("scope", "smartapp_snsapi_base")
	query.Set("grant_type", "client_credential")

	u.RawQuery = query.Encode()

	return u.String(), nil
}
