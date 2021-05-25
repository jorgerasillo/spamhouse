package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

type GraphQLMutation struct {
	Data struct {
		Enqueue struct {
			Status string   `json:"status"`
			Errors []string `json:"errors"`
		} `json:"enqueue"`
	} `json:"data"`
}

type GraphQLQuery struct {
	Data struct {
		Getipdetails struct {
			Node []struct {
				IPAddress    string    `json:"ip_address"`
				ResponseCode string    `json:"response_code"`
				UUID         string    `json:"uuid"`
				CreatedAt    time.Time `json:"created_at"`
				UpdatedAt    time.Time `json:"updated_at"`
			} `json:"node"`
		} `json:"getIPDetails"`
	} `json:"data"`
}

func makeHttpRequest(t *testing.T, method, url string, headers map[string]string, payload *strings.Reader) (*http.Response, error) {
	switch method {
	case http.MethodPost:

		req, err := http.NewRequest(method, url, payload)

		for k, v := range headers {
			req.Header.Add(k, v)
		}

		if err != nil {
			t.Error(err, "Unable to create request")
		}
		httpClient := &http.Client{}
		resp, err := httpClient.Do(req)

		return resp, err
	default:
		return nil, nil

	}
}

func parseMutationResponse(t *testing.T, resp *http.Response) GraphQLMutation {
	if resp == nil {
		t.Fatalf("response is ni")
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	var graphResponse GraphQLMutation
	t.Logf("response: %v", string(bodyBytes))
	if err := json.Unmarshal(bodyBytes, &graphResponse); err != nil {
		t.Fatalf("unable to parse response body, err: %v", err)
	}
	return graphResponse
}

func parseQueryResponse(t *testing.T, resp *http.Response) GraphQLQuery {
	if resp == nil {
		t.Fatalf("response is ni")
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	var graphResponse GraphQLQuery
	t.Logf("query response: %v", string(bodyBytes))
	if err := json.Unmarshal(bodyBytes, &graphResponse); err != nil {
		t.Fatalf("unable to prase response body")
	}
	return graphResponse
}
