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
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

// Backup structure
type Backup struct {
	ID           string `json:"id"`
	Deploymentid string `json:"deployment_id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Status       string `json:"status"`
}

// backupsResponse is used to represent and remove the JSON+HAL Embedded wrapper
type backupsResponse struct {
	Embedded struct {
		Backups []Backup `json:"backups"`
	} `json:"_embedded"`
}

//GetBackupsForDeploymentJSON returns raw deployment
func (c *Client) GetBackupsForDeploymentJSON(deploymentid string) (string, []error) {
	return c.getJSON("deployments/" + deploymentid + "/backups")
}

//GetBackupsForDeployment returns deployment structure
func (c *Client) GetBackupsForDeployment(deploymentid string) (*[]Backup, []error) {
	body, errs := c.GetBackupsForDeploymentJSON(deploymentid)

	if errs != nil {
		return nil, errs
	}

	backupsResponse := backupsResponse{}
	json.Unmarshal([]byte(body), &backupsResponse)
	Backups := backupsResponse.Embedded.Backups

	return &Backups, nil
}

//StartBackupJSON sets JSON scaling and returns string respones
func (c *Client) StartBackupJSON(params ScalingsParams) (string, []error) {
	response, body, errs := gorequest.New().Post(apibase+"deployments/"+params.DeploymentID+"/backups").
		Set("Authorization", "Bearer "+c.apiToken).
		Set("Content-type", "application/json; charset=utf-8").
		Send(params).
		End()

	if response.StatusCode != 200 { // Expect Accepted on success - assume error on anything else
		myerrors := Errors{}
		err := json.Unmarshal([]byte(body), &myerrors)
		if err != nil {
			errs = append(errs, fmt.Errorf("Unable to parse error - status code %d", response.StatusCode))
		} else {
			errs = append(errs, fmt.Errorf("%v", myerrors.Error))
		}
	}

	return body, errs
}

//StartBackup sets scale and returns recipe for scaling
func (c *Client) StartBackup(scalingsParams ScalingsParams) (*Recipe, []error) {
	body, errs := c.StartBackupJSON(scalingsParams)
	if errs != nil {
		return nil, errs
	}

	recipe := Recipe{}
	json.Unmarshal([]byte(body), &recipe)

	return &recipe, nil
}
