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
	"time"
)

// Deployment structure
type Deployment struct {
	ID                  string            `json:"id"`
	Name                string            `json:"name"`
	Type                string            `json:"type"`
	CreatedAt           time.Time         `json:"created_at"`
	ProvisionRecipeID   string            `json:"provision_recipe_id"`
	CACertificateBase64 string            `json:"ca_certificate_base64"`
	Connection          ConnectionStrings `json:"connection_strings"`
	Links               struct {
		ComposeWebUILink Link `json:"compose_web_ui"`
	} `json:"_links"`
}

// ConnectionStrings structure
type ConnectionStrings struct {
	Health   string   `json:"health"`
	SSH      string   `json:"ssh"`
	Admin    string   `json:"admin"`
	SSHAdmin string   `json:"ssh_admin"`
	CLI      []string `json:"cli"`
	Direct   []string `json:"direct"`
}

// DeploymentsResponse holding structure
type DeploymentsResponse struct {
	Embedded struct {
		Deployments []Deployment `json:"deployments"`
	} `json:"_embedded"`
}

//CreateDeploymentParams Parameters to be completed before creating a deployment
type CreateDeploymentParams struct {
	Name         string `json:"name"`
	AccountID    string `json:"account_id"`
	ClusterID    string `json:"cluster_id,omitempty"`
	Datacenter   string `json:"datacenter,omitempty"`
	DatabaseType string `json:"type"`
	Version      string `json:"version,omitempty"`
	Units        int    `json:"units,omitempty"`
	SSL          bool   `json:"ssl,omitempty"`
	WiredTiger   bool   `json:"wired_tiger,omitempty"`
}

//VersionTransition a struct wrapper for version transition information
type VersionTransition struct {
	Application string `json:"application"`
	Method      string `json:"method"`
	FromVersion string `json:"from_version"`
	ToVersion   string `json:"to_version"`
}

//VersionsResponse Version holding structure
type VersionsResponse struct {
	Embedded struct {
		VersionTransitions []VersionTransition `json:"transitions"`
	} `json:"_embedded"`
}
