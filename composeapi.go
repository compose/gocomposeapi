// Copyright 2016 Compose, an IBM Company
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package composeapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	apibase = "https://api.compose.io/2016-07/"
)

type Client struct {
	apiToken string
}

func NewClient(apiToken string) (*Client, error) {
	return &Client{
		apiToken: apiToken,
	}, nil
}

// Link structure for JSON+HAL links
type Link struct {
	HREF      string `json:"href"`
	Templated bool   `json:"templated"`
}

//Errors struct for parsing error returns
type Errors struct {
	Error map[string][]string `json:"errors,omitempty"`
}

func printJSON(jsontext string) {
	var tempholder map[string]interface{}

	if err := json.Unmarshal([]byte(jsontext), &tempholder); err != nil {
		log.Fatal(err)
	}
	indentedjson, err := json.MarshalIndent(tempholder, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(indentedjson))
}

//SetAPIToken overrides the API token
func (c *Client) SetAPIToken(newtoken string) {
	c.apiToken = newtoken
}

func (c *Client) reqJSON(endpoint, method string, body interface{}) (string, []error) {
	buffer := new(bytes.Buffer)
	err := json.NewEncoder(buffer).Encode(body)
	if err != nil {
		return "", []error{err}
	}

	req, err := http.NewRequest(method, apibase+endpoint, buffer)
	if err != nil {
		return "", []error{err}
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))
	req.Header.Add("Content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errs := []error{err}
		responseErrors := Errors{}
		err := json.NewDecoder(resp.Body).Decode(&responseErrors)
		if err != nil {
			errs = append(errs, fmt.Errorf("Unable to parse API errors - status code %d", resp.StatusCode))
		} else {
			for responseErrType, responseErrs := range responseErrors.Error {
				for _, err := range responseErrs {
					errs = append(errs, fmt.Errorf("%s: %s", responseErrType, err))
				}
			}
		}
		return "", errs
	}

	respBodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", []error{err}
	}
	return string(respBodyData), nil
}
