package main

import (
    "net/http"
    "net/url"
    "strings"
    "encoding/json"
    "fmt"
    "os"
    "io"
)

// Twilio base api url
const ApiBase string = "https://api.twilio.com/2010-04-01/Accounts/"

// Twilio SMS message endpoint
const ApiMessages string = "/Messages.json"

// AccountSidEnvVar is the environment variable name where the Twilio account sid should be read from
const AccountSidEnvVar string = "ACCOUNT_SID"

// AuthTokenEnvVar is the environment variable name where the Twilio auth token should be read from
const AuthTokenEnvVar string = "AUTH_TOKEN"

// NumberFromEnvVar is the environment variable name where the Twilio phone number should be read from
const NumberFromEnvVar string = "NUMBER_FROM"

// Client stores Twilio specific data & http client to send requests
type Client struct {
    // HTTP Client to use
    client *http.Client

    // Twilio account sid
    accountSid string

    // Twilio auth token
    authToken string

    // Twilio phone number
    numberFrom string
}

// Creates a new Client, this function requires the following environment variables to be set:
// ACCOUNT_SID, AUTH_TOKEN, NUMBER_FROM
func NewClient() *Client {
    return &Client{
        client: &http.Client{},
        accountSid: os.Getenv(AccountSidEnvVar),
        authToken: os.Getenv(AuthTokenEnvVar),
        numberFrom: os.Getenv(NumberFromEnvVar)}
}

// Creates an encoded message, contains:
// To phone number
// From phone number
// Message to be sent
func (c *Client) createMessage(to string, body string) string {
    data := url.Values{}
    data.Set("To", to)
    data.Set("From", c.numberFrom)
    data.Set("Body", body)

    return data.Encode()
}

// Sends an encoded string message to the Twilio SMS api
func (c *Client) send(message string) {
    req, _ := http.NewRequest("POST", c.getEndpoint(ApiMessages), strings.NewReader(message))
    req.SetBasicAuth(c.accountSid, c.authToken)
    req.Header.Add("Accept", "application/json")
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    resp, _ := c.client.Do(req)

    data, err := decodeJSON(resp.Body)

    if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
        if err == nil {
          fmt.Println("Success")
        }
    } else {
        fmt.Println(fmt.Sprintf("Status: %v", data["status"]));
        fmt.Println(fmt.Sprintf("Message: %v", data["message"]));
    }
}

// Returns the full Twilio endpoint as a string
// from the given endpoint slug
func (c *Client) getEndpoint(endpoint string) string {
    return ApiBase + c.accountSid + endpoint
}

// Decodes a json body, returns a map of data
func decodeJSON(body io.ReadCloser) (map[string]interface{}, error) {
    defer body.Close()

    var data map[string]interface{}
    decoder := json.NewDecoder(body)
    err := decoder.Decode(&data)

    return data, err
}
