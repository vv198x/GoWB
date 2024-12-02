package requests

import (
	"bytes"
	"context"
	"fmt"
	"github.com/vv198x/GoWB/config"
	"io/ioutil"
	"log/slog"
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

func (r *Request) Do(ctx context.Context) ([]byte, error) {
	client := &http.Client{Timeout: time.Second * 60}
	slog.Debug("request", r.Method, r.URI, string(r.Data))

	req, err := http.NewRequestWithContext(ctx, r.Method, r.URI, bytes.NewBuffer(r.Data))
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

func (r *Request) DoWithRetries(ctx context.Context) (responseData []byte, err error) {
	for i := 0; i < config.Get().Retries; i++ {

		responseData, err = r.Do(ctx)
		if err == nil {
			return responseData, nil
		}
		time.Sleep(time.Duration(config.Get().RetriesTime) * time.Millisecond)
	}
	return nil, fmt.Errorf("after %d attempts, last error: %w", config.Get().Retries, err)
}
