package main

import (
    "net/http"
    "net/url"
    "strings"
    "encoding/json"
    "fmt"
    "os"
)

const (
    API_BASE     string = "https://api.twilio.com/2010-04-01/Accounts/"
    API_MESSAGES string = "/Messages.json"
)

type Client struct {
    client *http.Client
    config Config
}

type Config struct {
    accountSid string
    authToken string
    numberFrom string
}

func (c *Client) setup() {
    c.client = &http.Client{}
    c.config = Config{}
    c.config.accountSid = os.Getenv("ACCOUNT_SID")
    c.config.authToken = os.Getenv("AUTH_TOKEN")
    c.config.numberFrom = os.Getenv("NUMBER_FROM")
}

func (c *Client) createMessage(to string, body string) string {
    data := url.Values{}
    data.Set("To", to)
    data.Set("From", c.config.numberFrom)
    data.Set("Body", body)

    return strings.Replace(data.Encode(), "%2B", "+", -1)
}

func (c *Client) send(message string) {
    req, _ := http.NewRequest("POST", c.getEndpoint(API_MESSAGES), strings.NewReader(message))
    req.SetBasicAuth(c.config.accountSid, c.config.authToken)
    req.Header.Add("Accept", "application/json")
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    resp, _ := c.client.Do(req)
    if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
        var data map[string]interface{}
        decoder := json.NewDecoder(resp.Body)
        err := decoder.Decode(&data)
        if err == nil {
          fmt.Println("Success")
        }
    } else {
        fmt.Println(data["message"]);
    }
}

func (c *Client) getEndpoint(endpoint string) string {
    return API_BASE + c.config.accountSid + endpoint
}
