package qq_miniprogram

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/imroc/req"
)

func Login(appId, appSecret, code string) (lr LoginResponse, err error) {
	api, err := CodeToURL(appId, appSecret, code)
	if err != nil {
		return lr, err
	}

	r, err := req.Get(api)
	if err != nil {
		return lr, err
	}

	if r.Response().StatusCode != 200 {
		return lr, ErrConnectTencentServer
	}

	err = r.ToJSON(&lr)
	if err != nil {
		return lr, err
	}

	if lr.Errcode != 0 {
		return lr, errors.New(lr.Errmsg)
	}

	return lr, nil
}

func SendCustomMessage(accessToken string, toUser string, templateId string, page string, formId string, keywords []string, emphasisWordIndex int) (cr CommonResponse, err error) {
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

	r, err := req.Post(api, req.BodyJSON(bf.String()))
	if err != nil {
		return cr, err
	}

	if r.Response().StatusCode != 200 {
		return cr, ErrConnectTencentServer
	}

	err = r.ToJSON(&cr)
	if err != nil {
		return cr, err
	}

	if cr.Errcode != 0 {
		return cr, errors.New(cr.Errmsg)
	}

	return cr, nil
}
