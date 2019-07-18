package qq_miniprogram

import "net/url"

const (
	// BaseURL 基础URL
	baseURL           = "https://api.q.qq.com"
	codeAPI           = "/sns/jscode2session"
	tokenAPI          = "/api/getToken"
	sendCustomMessage = "/api/json/template/send"
)

func CodeToURL(appId, appSecret, code string) (s string, err error) {
	if appId == "" || appSecret == "" || code == "" {
		return s, ErrNotAllowEmptyParam
	}
	u, err := url.Parse(baseURL + codeAPI)
	if err != nil {
		return s, err
	}

	query := u.Query()

	query.Set("appid", appId)
	query.Set("secret", appSecret)
	query.Set("js_code", code)
	query.Set("grant_type", "authorization_code")

	u.RawQuery = query.Encode()

	return u.String(), nil
}

func TokenURL(appId, appSecret string) (s string, err error) {
	if appId == "" || appSecret == "" {
		return s, ErrNotAllowEmptyParam
	}
	u, err := url.Parse(baseURL + tokenAPI)
	if err != nil {
		return s, err
	}

	query := u.Query()

	query.Set("appid", appId)
	query.Set("secret", appSecret)
	query.Set("grant_type", "client_credential")

	u.RawQuery = query.Encode()

	return u.String(), nil
}

func SendCustomMessageURL(accessToken string) (s string, err error) {
	if accessToken == "" {
		return s, ErrNotAllowEmptyParam
	}
	u, err := url.Parse(baseURL + sendCustomMessage)
	if err != nil {
		return s, err
	}

	query := u.Query()

	query.Set("access_token", accessToken)

	u.RawQuery = query.Encode()

	return u.String(), nil
}
