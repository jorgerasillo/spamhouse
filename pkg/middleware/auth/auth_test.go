// +build integration

package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestAuthHandler(t *testing.T) {

	r, err := makeHttpRequest(t, "POST", "http://spamhouse:8080/", nil, strings.NewReader(""))
	if err != nil {
		t.Fatalf("invalid request, %s", err)
	}

	if r == nil {
		t.Fatalf("response is nil")
	}

	if r.StatusCode != http.StatusUnauthorized {
		t.Fatalf("invalid status code: %v, body: %v", r.StatusCode, r.Body)
	}

}

func TestValidEnqueue(t *testing.T) {
	validHeaders := map[string]string{
		"Authorization": "Basic c2VjdXJld29ya3M6c3VwZXJzZWNyZXQ=",
		"Content-Type":  "application/json",
	}

	query := `{"query":"mutation{\n  enqueue(input: [\"1.2.3.8\"]){\n    status\n errors\n  }\n}"}`
	body := strings.NewReader(query)

	r2, err := makeHttpRequest(t, "POST", "http://spamhouse:8080/graphql", validHeaders, body)
	if err != nil {
		t.Fatalf("invalid request, err: %v", err)
	}

	if r2 == nil {
		t.Fatalf("response is nil")
	}

	if r2.StatusCode != http.StatusOK {
		t.Fatalf("invalid status code,: %v, body: %v", r2.StatusCode, r2.Body)
	}

	gql := parseMutationResponse(t, r2)
	if gql.Data.Enqueue.Status != "Success" {
		t.Fatalf("errors in response, response: %v", gql)
	}

}

func TestInvalidEnqueue(t *testing.T) {
	validHeaders := map[string]string{
		"Authorization": "Basic c2VjdXJld29ya3M6c3VwZXJzZWNyZXQ=",
		"Content-Type":  "application/json",
	}
	query := `{"query":"mutation{\n  enqueue(input: [\"1.2.3.88888\"]){\n    status\n errors\n  }\n}"}`
	body := strings.NewReader(query)

	r2, err := makeHttpRequest(t, "POST", "http://spamhouse:8080/graphql", validHeaders, body)
	if err != nil {
		t.Fatalf("invalid request, err: %v", err)
	}

	if r2 == nil {
		t.Fatalf("response is nil")
	}

	if r2.StatusCode != http.StatusOK {
		t.Fatalf("invalid status code,: %v, body: %v", r2.StatusCode, r2.Body)
	}

	gql := parseMutationResponse(t, r2)
	if gql.Data.Enqueue.Status != "Failure" {
		t.Fatalf("errors in response, response: %v", gql)
	}

	// ensure errors were returned, since the ip was invalid
	if len(gql.Data.Enqueue.Errors) < 1 {
		t.Fatalf("we were expecting errors, but didn't get any: %v", gql)
	}

}

func TestEnqueueWithQuery(t *testing.T) {
	validHeaders := map[string]string{
		"Authorization": "Basic c2VjdXJld29ya3M6c3VwZXJzZWNyZXQ=",
		"Content-Type":  "application/json",
	}
	// send mutation
	query := `{"query":"mutation{\n  enqueue(input: [\"1.2.3.8\"]){\n    status\n errors\n  }\n}"}`
	body := strings.NewReader(query)

	r2, err := makeHttpRequest(t, "POST", "http://spamhouse:8080/graphql", validHeaders, body)
	if err != nil {
		t.Fatalf("invalid request, err: %v", err)
	}

	if r2 == nil {
		t.Fatalf("response is nil")
	}

	if r2.StatusCode != http.StatusOK {
		t.Fatalf("invalid status code,: %v, body: %v", r2.StatusCode, r2.Body)
	}

	// ensure successful mutation response
	gql := parseMutationResponse(t, r2)
	if gql.Data.Enqueue.Status != "Success" {
		t.Fatalf("errors in response, response: %v", gql)
	}

	// send query
	queryPayload := `{"query":"query{\n  getIPDetails(input: \"1.2.3.3\"){\n    node{\n      ip_address\n      response_code\n      uuid\n      created_at\n      updated_at\n    }\n  }\n}"}`
	queryBody := strings.NewReader(queryPayload)
	queryResponse, err := makeHttpRequest(t, "POST", "http://spamhouse:8080/graphql", validHeaders, queryBody)
	if err != nil {
		t.Fatalf("invalid request, err: %v", err)
	}

	if queryResponse == nil {
		t.Fatalf("response is nil")
	}

	if queryResponse.StatusCode != http.StatusOK {
		t.Fatalf("invalid status code,: %v, body: %v", r2.StatusCode, r2.Body)
	}

	// validate response contains nodes
	gql2 := parseQueryResponse(t, queryResponse)
	if gql2.Data.Getipdetails.Node.IPAddress != "1.2.3.3" {
		t.Fatalf("ip not found, gql2 response: %v", gql2)
	}

}
