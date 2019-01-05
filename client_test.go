package main

import (
	"net/http"
	"testing"
)

// Assert the query is as expected
func assertQuery(t *testing.T, expected string, actual string) {
	t.Log("Asserting query is as expected")
	if expected != actual {
		t.Errorf("Expected %s,\nGot %s", expected, actual)
	}
}

// Assert the request method is as expected
func assetMethod(t *testing.T, expected string, request *http.Request) {
	t.Log("Asserting request method is as expected")
	if expected != request.Method {
		t.Errorf("Expected %s,\nGot %s", expected, request.Method)
	}
}

// Test
func TestSend(t *testing.T) {
	testClient := Client{client: &http.Client{},
		accountSid: "accountSid",
		authToken:  "authToken",
		numberFrom: "numberFrom"}

	expectedMessage := testClient.createMessage("numberTo", "Test Body")
	actualMessage := "Body=Test+Body&From=numberFrom&To=numberTo"
	assertQuery(t, expectedMessage, actualMessage)

	request := testClient.createRequest(expectedMessage)
	assetMethod(t, "POST", request)
}
