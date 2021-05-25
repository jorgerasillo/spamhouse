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

	// r3, err := MakeHttpRequest(t, "POST", "http://localhost:8080/", validHeaders, queryPayload)

}

func TestEnqueue(t *testing.T) {
	validHeaders := map[string]string{
		"Authorization": "Basic c2VjdXJld29ya3M6c3VwZXJzZWNyZXQ=",
		"Content-Type":  "application/json",
	}

	query := `{"query":"mutation{\n  enqueue(input: [\"1.2.3.8888888\"]){\n    status\n errors\n  }\n}"}`
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

// func TestInvalidEnqueue(t *testing.T) {
// 	validHeaders := map[string]string{
// 		"Authorization": "Basic c2VjdXJld29ya3M6c3VwZXJzZWNyZXQ=",
// 		"Content-Type":  "application/json",
// 	}
// 	enqueuePayload := map[string]string{
// 		"query": `
// 	 	  mutation{
// 			enqueue(input: ["1.2.3.888888"]){
// 			  status
// 			  errors
// 			}
// 		  }`,
// 	}
// 	r2, err := makeHttpRequest(t, "POST", "http://spamhouse:8080/", validHeaders, enqueuePayload)
// 	if err != nil {
// 		t.Fatalf("invalid request, err: %v", err)
// 	}

// 	if r2 == nil {
// 		t.Fatalf("response is nil")
// 	}

// 	if r2.StatusCode != http.StatusOK {
// 		t.Fatalf("invalid status code,: %v, body: %v", r2.StatusCode, r2.Body)
// 	}

// 	gql := parseMutationResponse(t, r2)
// 	if gql.Data.Enqueue.Status != "Failure" {
// 		t.Fatalf("errors in response, response: %v", gql)
// 	}

// 	if len(gql.Data.Enqueue.Errors) < 1 {
// 		t.Fatalf("we were expecting errors, but didn't get any: %v", gql)
// 	}

// }

// func TestEnqueuWithQuery(t *testing.T) {
// 	validHeaders := map[string]string{
// 		"Authorization": "Basic c2VjdXJld29ya3M6c3VwZXJzZWNyZXQ=",
// 		"Content-Type":  "application/json",
// 	}
// 	enqueuePayload := map[string]string{
// 		"query": `
// 	 	  mutation{
// 			enqueue(input: ["1.2.3.8"]){
// 			  status
// 			  errors
// 			}
// 		  }`,
// 	}
// 	r2, err := makeHttpRequest(t, "POST", "http://spamhouse:8080/", validHeaders, enqueuePayload)
// 	if err != nil {
// 		t.Fatalf("invalid request, err: %v", err)
// 	}

// 	if r2 == nil {
// 		t.Fatalf("response is nil")
// 	}

// 	if r2.StatusCode != http.StatusOK {
// 		t.Fatalf("invalid status code,: %v, body: %v", r2.StatusCode, r2.Body)
// 	}

// 	gql := parseMutationResponse(t, r2)
// 	if gql.Data.Enqueue.Status != "Success" {
// 		t.Fatalf("errors in response, response: %v", gql)
// 	}

// 	queryPayload := map[string]string{
// 		"query": `
//             {
//                 getIPDetails {
// 					node{
// 						ip_address
// 						response_code
// 						uuid
// 						created_at
// 						updated_at
// 					  }
//                 }
//             }
//         `,
// 	}
// 	queryResponse, err := makeHttpRequest(t, "POST", "http://spamhouse:8080/", validHeaders, queryPayload)
// 	if err != nil {
// 		t.Fatalf("invalid request, err: %v", err)
// 	}

// 	if queryResponse == nil {
// 		t.Fatalf("response is nil")
// 	}

// 	if queryResponse.StatusCode != http.StatusOK {
// 		t.Fatalf("invalid status code,: %v, body: %v", r2.StatusCode, r2.Body)
// 	}

// 	gql2 := parseQueryResponse(t, queryResponse)
// 	if len(gql2.Data.Getipdetails.Node) < 1 {
// 		t.Fatalf("errors in response, response: %v", gql2)
// 	}

// }
