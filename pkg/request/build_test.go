package request

import (
	"github.com/YuanJey/http_client/pkg/client"
	"github.com/YuanJey/http_client/pkg/consts"
	"testing"
)

func TestBuild(t *testing.T) {
	req := Req{
		UrlencodedMap: nil,
		Json: &struct {
			Name string `json:"name"`
		}{
			Name: "111",
		},
	}
	resp := struct {
		Code int `json:"code"`
	}{
		Code: 0,
	}
	request, err := Build(consts.MethodGet, "http://www.baidu.com", consts.ContentTypeJson, &req, nil)
	if err != nil {
		t.Error(err)
	}
	err = client.Request(request, &resp)
	if err != nil {
		return
	}
}
