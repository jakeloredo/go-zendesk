package zendesk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Zendesk is the main wrapper
type Zendesk struct {
	Subdomain string
	Username  string
	Password  string
	Token     string
}

// ResponseData contains the fields that could be returned by any endpoint
type ResponseData struct {
	ExecutedView
	ErrorString  string          `json:"error,omitempty"`
	Tickets      []Ticket        `json:"tickets,omitempty"`
	Ticket       Ticket          `json:"ticket,omitempty"`
	Comments     []TicketComment `json:"comments,omitempty"`
	NextPage     string          `json:"next_page,omitempty"`
	PreviousPage string          `json:"previous_page,omitempty"`
	Count        int             `json:"count,omitempty"`
	Result       struct {
		Ticket Ticket `json:"ticket,omitempty"`
	} `json:"result,omitempty"`
	//Rows         []ViewRow         `json:"rows"`
}


func (zd *Zendesk) doRequest(endpoint string, requestType string, payload map[string]interface{}) (ResponseData, error) {

	data := ResponseData{}

	urlComponents := []string{"https://", zd.Subdomain, ".zendesk.com/api/v2/", endpoint}
	url := strings.Join(urlComponents, "")
	client := http.Client{Timeout: time.Second * 10}

	var req *http.Request
	var httpErr error
	var jsonErr error

	if payload != nil {
		var jsonPayload []byte
		jsonPayload, jsonErr = json.Marshal(payload)
		if jsonErr != nil {
			return data, jsonErr
		}
		fmt.Println(string(jsonPayload))
		req, httpErr = http.NewRequest(requestType, url, bytes.NewBuffer(jsonPayload))
	} else {
		req, httpErr = http.NewRequest(requestType, url, nil)
	}

	if httpErr != nil {
		return data, httpErr
	}

	req.Header.Set("content-type", "application/json")

	if zd.Password == "" {
		req.SetBasicAuth(strings.Join([]string{zd.Username, "/token"}, ""), zd.Token)
	} else {
		req.SetBasicAuth(zd.Username, zd.Password)
	}

	res, resErr := client.Do(req)
	//fmt.Println(res)

	if resErr != nil {
		return data, resErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return data, readErr
	}

	jsonErr = json.Unmarshal(body, &data)

	if jsonErr != nil {
		return data, jsonErr
	}

	return data, nil
}

func (zd *Zendesk) post(endpoint string, data map[string]interface{}) (ResponseData, error) {
	return zd.doRequest(endpoint, "POST", data)
}

func (zd *Zendesk) put(endpoint string, data map[string]interface{}) (ResponseData, error) {
	return zd.doRequest(endpoint, "PUT", data)
}

func (zd *Zendesk) get(endpoint string) (ResponseData, error) {
	return zd.doRequest(endpoint, "GET", nil)
}
