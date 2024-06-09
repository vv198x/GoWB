package requests

import (
	"bytes"
	"fmt"
	"github.com/vv198x/GoWB/config"
	"io/ioutil"
	"net/http"
	"time"
)

type Request struct {
	Method string
	URI    string
	Data   []byte
}

func New(method, uri string, data []byte) *Request {
	return &Request{
		Method: method,
		URI:    uri,
		Data:   data,
	}
}

func (r *Request) Do() ([]byte, error) {
	client := &http.Client{Timeout: time.Second * 60}

	req, err := http.NewRequest(r.Method, r.URI, bytes.NewBuffer(r.Data))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Get().ApiKeyAD))

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return responseData, nil
}

func (r *Request) DoWithRetries() (responseData []byte, err error) {
	for i := 0; i < config.Get().Retries; i++ {

		responseData, err = r.Do()
		if err == nil {
			return responseData, nil
		}
		time.Sleep(time.Duration(config.Get().RetriesTime) * time.Millisecond)
	}
	return nil, fmt.Errorf("after %d attempts, last error: %w", config.Get().Retries, err)
}
