package qq_miniprogram

import "fmt"

// https://q.qq.com/wiki/develop/miniprogram/server/

// Response 基础数据
type CommonResponse struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

type LoginResponse struct {
	CommonResponse
	Unionid    string `json:"unionid,omitempty"`
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

type AccessToken struct {
	CommonResponse
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// https://q.qq.com/wiki/develop/miniprogram/server/open_port/port_module.html#sendtemplatemessage

type CustomMessage struct {
	ToUser           string             `json:"touser"`
	TemplateId       string             `json:"template_id"`
	Page             string             `json:"page"`
	FormId           string             `json:"form_id"`
	Data             map[string]Keyword `json:"data"`
	Emphasis_keyword string             `json:"emphasis_keyword"`
}

type Keyword struct {
	Value string `json:"value"`
}

func buildCustomMessage(toUser string, templateId string, page string, formId string, keywords []string, emphasisWordIndex int) CustomMessage {
	data := map[string]Keyword{}
	for key, value := range keywords {
		data[fmt.Sprintf("keyword%d", key+1)] = Keyword{
			Value: value,
		}
	}

	return CustomMessage{
		ToUser:           toUser,
		TemplateId:       templateId,
		Page:             page,
		FormId:           formId,
		Data:             data,
		Emphasis_keyword: fmt.Sprintf("keyword%d.DATA", emphasisWordIndex+1),
	}
}
