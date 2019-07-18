package qq_miniprogram

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"strings"
)

func (c *Client) Login(appId, appSecret, code string) (lr LoginResponse, err error) {
	api, err := CodeToURL(appId, appSecret, code)
	if err != nil {
		return lr, err
	}

	res, err := c.client.Get(api)

	if err != nil {
		return lr, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return lr, ErrConnectTencentServer
	}

	err = json.NewDecoder(res.Body).Decode(&lr)
	if err != nil {
		return lr, err
	}

	if lr.Errcode != 0 {
		return lr, errors.New(lr.Errmsg)
	}

	return lr, nil
}

func (c *Client) SendCustomMessage(accessToken string, toUser string, templateId string, page string, formId string, keywords []string, emphasisWordIndex int) (cr CommonResponse, err error) {
	api, err := SendCustomMessageURL(accessToken)
	if err != nil {
		return cr, err
	}

	msg := buildCustomMessage(toUser, templateId, page, formId, keywords, emphasisWordIndex)

	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)

	err = jsonEncoder.Encode(msg)
	if err != nil {
		return cr, err
	}

	log.Println(bf.String())

	res, err := c.client.Post(api, "application/json", strings.NewReader(bf.String()))

	if err != nil {
		return cr, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return cr, ErrConnectTencentServer
	}

	err = json.NewDecoder(res.Body).Decode(&cr)
	if err != nil {
		return cr, err
	}

	if cr.Errcode != 0 {
		return cr, errors.New(cr.Errmsg)
	}

	return cr, nil
}
