package get

import (
	"encoding/json"
	"errors"
	"github.com/YuanJey/goutils2/pkg/utils"
	"github.com/YuanJey/http_client/pkg/sign"
	"io"
	"net/http"
	"strings"
	"time"
)

func Get(url string, req interface{}, resp interface{}, sign sign.RequestSign) error {
	body := strings.NewReader("")
	if req != nil {
		jsonStr, err := json.Marshal(req)
		if err != nil {
			return err
		}
		body = strings.NewReader(string(jsonStr))
	}
	request, err := http.NewRequest("GET", url, body)
	if err != nil {
		return err
	}
	if sign != nil {
		sign.Sign(request)
	}
	client := http.Client{Timeout: 5 * time.Second}
	httpResponse, err := client.Do(request)
	if err != nil {
		return err
	}
	result, err := io.ReadAll(httpResponse.Body)
	if httpResponse.StatusCode != 200 {
		return utils.Wrap(errors.New(httpResponse.Status), "status code failed "+url+string(result))
	}
	err = utils.JsonStringToStruct(string(result), &resp)
	if err != nil {
		return err
	}
	return nil
}
