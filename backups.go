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

//GetBackupsJSON returns raw deployment
func (c *Client) GetBackupsJSON(deploymentid string) (string, []error) {
	return c.getJSON("deployments/" + deploymentid + "/backups")
}

//GetBackups returns deployment structure
func (c *Client) GetBackups(deploymentid string) (*[]Backup, []error) {
	body, errs := c.GetBackupsJSON(deploymentid)

	if errs != nil {
		return nil, errs
	}

	backupsResponse := backupsResponse{}
	json.Unmarshal([]byte(body), &backupsResponse)
	Backups := backupsResponse.Embedded.Backups

	return &Backups, nil
}
