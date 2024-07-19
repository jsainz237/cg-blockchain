package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

func CallRPC(url string, serviceMethod string, args interface{}) (*http.Response, *interface{}, error) {
	var response interface{}

	reqBody, _ := json.Marshal(map[string]interface{}{
		"id":      uuid.New().String(),
		"jsonrpc": "2.0",
		"method":  serviceMethod,
		"params":  []interface{}{args},
	})

	resp, err := http.Post(url+"/rpc", "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		defer resp.Body.Close()
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return nil, nil, fmt.Errorf("RPC call failed: %s", string(body))
	}

	defer resp.Body.Close()
	return resp, &response, nil
}
