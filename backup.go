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
)

// Deployment structure
type Backup struct {
	ID           string `json:"id"`
	DeploymentID string `json:"deployment_id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Status       string `json:"status"`
	DownloadLink string `json:"download_link"`
}

type CreateBackupParams struct {
	DeploymentID string `json:"-"`
}

type GetBackupsParams struct {
	DeploymentID string `json:"-"`
}

type GetBackupParams struct {
	DeploymentID string `json:"-"`
	BackupID     string `json:"-"`
}

type BackupRestoreParams struct {
	DeploymentID string `json:"-"`
	BackupID     string `json:"-"`
	Name         string `json:"name"`
	Cluster      string `json:"cluster,omitempty"`
	Datacenter   string `json:"datacenter,omitempty"`
	Version      string `json:"version,omitempty"`
	SSL          string `json:"ssl,omitempty"`
}

type backupsResponse struct {
	Embedded struct {
		Backups []Backup `json:"backups"`
	} `json:"_embedded"`
}

//CreateDeploymentJSON performs the call
func (c *Client) CreateBackupJSON(params CreateBackupParams) (string, []error) {
	return c.reqJSON(fmt.Sprintf("deployments/%s/backups", params.DeploymentID), "POST", nil)
}

//CreateDeployment creates a deployment
func (c *Client) CreateBackup(params CreateBackupParams) (*Backup, []error) {

	// This is a POST not a GET, so it builds its own request

	body, errs := c.CreateBackupJSON(params)

	if errs != nil {
		return nil, errs
	}

	backup := Backup{}
	json.Unmarshal([]byte(body), &backup)

	return &backup, nil
}

//GetDeploymentsJSON returns raw deployment
func (c *Client) GetBackupsJSON(params *GetBackupsParams) (string, []error) {
	return c.reqJSON(fmt.Sprintf("deployments/%s/backups", params.DeploymentID), "GET", nil)
}

//GetDeployments returns deployment structure
func (c *Client) GetBackups(params *GetBackupsParams) (*[]Backup, []error) {
	body, errs := c.GetDeploymentsJSON()

	if errs != nil {
		return nil, errs
	}

	backupResp := backupsResponse{}
	json.Unmarshal([]byte(body), &backupResp)
	backups := backupResp.Embedded.Backups

	return &backups, nil
}

//GetDeploymentsJSON returns raw deployment
func (c *Client) GetBackupJSON(params *GetBackupParams) (string, []error) {
	return c.reqJSON(fmt.Sprintf("deployments/%s/backups/%s", params.DeploymentID, params.BackupID), "GET", nil)
}

//GetDeployments returns deployment structure
func (c *Client) GetBackup(params *GetBackupParams) (*Backup, []error) {
	body, errs := c.GetBackupJSON(params)
	if errs != nil {
		return nil, errs
	}

	backup := Backup{}
	json.Unmarshal([]byte(body), &backup)

	return &backup, nil
}

func (c *Client) RestoreBackupJSON(params *BackupRestoreParams) (string, []error) {
	return c.reqJSON(fmt.Sprintf("deployments/%s/backups/%s", params.DeploymentID, params.BackupID), "POST", params)
}

func (c *Client) RestoreBackup(params *BackupRestoreParams) (*Deployment, []error) {
	body, errs := c.RestoreBackupJSON(params)
	if errs != nil {
		return nil, errs
	}

	deployment := Deployment{}
	json.Unmarshal([]byte(body), &deployment)

	return &deployment, nil
}
