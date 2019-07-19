package swan_miniprogram

import (
	"encoding/json"
	"errors"
)

func (c *Client) GetAccessToken(appKey, appSecret string) (ak AccessToken, err error) {
	api, err := TokenURL(appKey, appSecret)
	if err != nil {
		return ak, err
	}

	res, err := c.client.Get(api)

	if err != nil {
		return ak, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return ak, ErrConnectBaiduServer
	}

	err = json.NewDecoder(res.Body).Decode(&ak)
	if err != nil {
		return ak, err
	}

	if ak.Errcode != 0 {
		return ak, errors.New(ak.Errmsg)
	}

	return ak, nil
}
