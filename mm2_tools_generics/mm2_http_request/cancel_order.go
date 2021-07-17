package mm2_http_request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kpango/glg"
	"io/ioutil"
	http2 "mm2_client/http"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"net/http"
)

func NewCancelOrderRequest(uuid string) *mm2_data_structure.CancelOrderRequest {
	genReq := http2.NewGenericRequest("cancel_order")
	req := &mm2_data_structure.CancelOrderRequest{Userpass: genReq.Userpass, Method: genReq.Method, Uuid: uuid}
	return req
}

func CancelOrder(uuid string) (*mm2_data_structure.CancelOrderAnswer, error) {
	req := NewCancelOrderRequest(uuid).ToJson()
	resp, err := http.Post(http2.GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
	if err != nil {
		glg.Errorf("Err: %v", err)
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		var answer = &mm2_data_structure.CancelOrderAnswer{}
		decodeErr := json.NewDecoder(resp.Body).Decode(answer)
		if decodeErr != nil {
			glg.Errorf("Err: %v", decodeErr)
			return nil, decodeErr
		}
		return answer, nil
	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		errStr := fmt.Sprintf("err: %s\n", bodyBytes)
		return nil, errors.New(errStr)
	}
}