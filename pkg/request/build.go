package request

import (
	"encoding/json"
	"errors"
	"github.com/YuanJey/goutils2/pkg/utils"
	"github.com/YuanJey/http_client/pkg/consts"
	"github.com/YuanJey/http_client/pkg/sign"
	"net/http"
	url2 "net/url"
	"strings"
)

// application/x-www-form-urlencoded

type Req struct {
	UrlencodedMap map[string]string
	Json          interface{}
}

func Build(method, url, contentType string, Req *Req, sign sign.RequestSign) (*http.Request, error) {
	switch contentType {
	case consts.ContentTypeJson:
		return buildJson(method, url, Req.Json, sign)
	case consts.ContentTypeForm:
		return buildUrlencoded(method, url, Req.UrlencodedMap, sign)
	default:
		return nil, utils.Wrap(errors.New("method err "), "invalid method")
	}
}
func buildJson(method, url string, req interface{}, httpSign sign.RequestSign) (*http.Request, error) {
	body := strings.NewReader("")
	if req != nil {
		jsonStr, err := json.Marshal(req)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(string(jsonStr))
	}
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", consts.ContentTypeJson)
	if httpSign != nil {
		httpSign.Sign(request)
	}
	return request, nil
}
func buildUrlencoded(method, url string, req map[string]string, httpSign sign.RequestSign) (*http.Request, error) {
	form := url2.Values{}
	for key, value := range req {
		form.Add(key, value)
	}
	body := strings.NewReader(form.Encode())
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", consts.ContentTypeForm)
	if httpSign != nil {
		httpSign.Sign(request)
	}
	return request, nil
}
