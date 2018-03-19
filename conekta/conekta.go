package conekta

import (
	"net/http"
	"bytes"
	"encoding/json"
)

var (
	ApiKey, ApiVersion = "", "2.0.0"
)

const (
	conektaUrl = "https://api.conekta.io"
)

func request(method, url string, v interface{}) (statusCode int) {
	jsonPayload, err := json.Marshal(v)
	if err != nil {
		return
	}
	payload := bytes.NewReader(jsonPayload)
	req, _ := http.NewRequest(method, url, payload)
	req.Header.Add("accept", "application/vnd.conekta-v"+ApiVersion+"+json")
	req.SetBasicAuth(ApiKey, "")
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)
	return res.StatusCode
}
