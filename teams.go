// Copyright 2017 Compose, an IBM Company
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

// Team structure
type Team struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Users []User `json:"users"`
}

type teamsResponse struct {
	Embedded struct {
		Teams []Team `json:"teams"`
	} `json:"_embedded"`
}

type createTeamParams struct {
	Team TeamParams `json:"team"`
}

type patchTeamParams struct {
	Team TeamParams `json:"team"`
}

type putTeamUsersParams struct {
	UserIDs []string `json:"user_ids"`
}

// TeamParams core parameters for a new team
type TeamParams struct {
	Name string `json:"name"`
}

// CreateTeamJSON performs the call to create a team
func (c *Client) CreateTeamJSON(params TeamParams) (string, []error) {
	teamParams := createTeamParams{Team: params}

	response, body, errs := gorequest.New().Post(apibase+"teams").
		Set("Authorization", "Bearer "+c.apiToken).
		Set("Content-type", "application/json; charset=utf-8").
		Send(teamParams).
		End()

	if response.StatusCode != 201 { // Expect Created on success - assume error on anything else
		myErrors := Errors{}
		err := json.Unmarshal([]byte(body), &myErrors)
		if err != nil {
			errs = append(errs, fmt.Errorf("Unable to parse error - status code %d - body %s",
				response.StatusCode, response.Body))
		} else {
			errs = append(errs, fmt.Errorf("%v", myErrors.Error))
		}
	}
	return body, errs
}

// CreateTeam creates a team
func (c *Client) CreateTeam(params TeamParams) (*Team, []error) {
	body, errs := c.CreateTeamJSON(params)

	if errs != nil {
		return nil, errs
	}

	team := Team{}
	json.Unmarshal([]byte(body), &team)

	return &team, nil
}

// GetTeamsJSON returns raw teams
func (c *Client) GetTeamsJSON() (string, []error) {
	return c.getJSON("teams")
}

// GetTeams returns team structure
func (c *Client) GetTeams() (*[]Team, []error) {
	body, errs := c.GetTeamsJSON()

	if errs != nil {
		return nil, errs
	}

	teamResponse := teamsResponse{}
	json.Unmarshal([]byte(body), &teamResponse)
	teams := teamResponse.Embedded.Teams

	return &teams, nil
}

// GetTeamJSON returns a raw team
func (c *Client) GetTeamJSON(teamID string) (string, []error) {
	return c.getJSON("teams/" + teamID)
}

// GetTeam returns team structure
func (c *Client) GetTeam(teamID string) (*Team, []error) {
	body, errs := c.GetTeamJSON(teamID)

	if errs != nil {
		return nil, errs
	}

	team := Team{}
	json.Unmarshal([]byte(body), &team)

	return &team, nil
}

// GetTeamByName returns a team of a given name
func (c *Client) GetTeamByName(teamName string) (*Team, []error) {
	teams, errs := c.GetTeams()
	if errs != nil {
		return nil, errs
	}

	for _, team := range *teams {
		if team.Name == teamName {
			return &team, nil
		}
	}

	return nil, []error{fmt.Errorf("team not found: %s", teamName)}
}

// DeleteTeamJSON performs that call
func (c *Client) DeleteTeamJSON(teamID string) (string, []error) {
	response, body, errs := gorequest.New().Delete(apibase+"teams/"+teamID).
		Set("Authorization", "Bearer "+c.apiToken).
		Set("Content-type", "application/json; charset=utf-8").
		End()

	if response.StatusCode != 200 { // Expect OK on success - assume error on anything else
		myErrors := Errors{}
		err := json.Unmarshal([]byte(body), &myErrors)
		if err != nil {
			errs = append(errs, fmt.Errorf("Unable to parse error - status code %d - body %s",
				response.StatusCode, body))
		} else {
			errs = append(errs, fmt.Errorf("%v", myErrors.Error))
		}
	}

	return body, errs
}

// DeleteTeam deletes a team
func (c *Client) DeleteTeam(teamID string) (*Team, []error) {
	body, errs := c.DeleteTeamJSON(teamID)
	if errs != nil {
		return nil, errs
	}

	team := Team{}
	json.Unmarshal([]byte(body), &team)

	return &team, nil
}

// PatchTeamJSON changes a team's name
func (c *Client) PatchTeamJSON(teamID, teamName string) (string, []error) {
	patchParams := patchTeamParams{Team: TeamParams{Name: teamName}}

	response, body, errs := gorequest.New().Patch(apibase+"teams/"+teamID).
		Set("Authorization", "Bearer "+c.apiToken).
		Set("Content-type", "application/json; charset=utf-8").
		Send(patchParams).
		End()

	if response.StatusCode != 200 { // Expect OK on success - assume error on anything else
		myErrors := Errors{}
		err := json.Unmarshal([]byte(body), &myErrors)
		if err != nil {
			errs = append(errs, fmt.Errorf("Unable to parse error - status code %d - body %s",
				response.StatusCode, body))
		} else {
			errs = append(errs, fmt.Errorf("%v", myErrors.Error))
		}
	}

	return body, errs
}

// PatchTeam changes a team name
func (c *Client) PatchTeam(teamID, teamName string) (*Team, []error) {
	body, errs := c.PatchTeamJSON(teamID, teamName)
	if errs != nil {
		return nil, errs
	}

	team := Team{}
	json.Unmarshal([]byte(body), &team)

	return &team, nil
}

// PutTeamUsersJSON performs the call
func (c *Client) PutTeamUsersJSON(teamID string, userIDs []string) (string, []error) {
	putUsers := putTeamUsersParams{UserIDs: userIDs}

	response, body, errs := gorequest.New().
		Put(apibase+"teams/"+teamID+"/users").
		Set("Authorization", "Bearer "+c.apiToken).
		Set("Content-type", "application/json; charset=utf-8").
		Send(putUsers).
		End()

	if response.StatusCode != 200 { // Expect OK on success - assume error on anything else
		myErrors := Errors{}
		err := json.Unmarshal([]byte(body), &myErrors)
		if err != nil {
			errs = append(errs, fmt.Errorf("Unable to parse error - status code %d - body %s", response.StatusCode, body))
		} else {
			errs = append(errs, fmt.Errorf("%v", myErrors.Error))
		}
	}

	return body, errs
}

// Put TeamUsers adds users to the given team
func (c *Client) PutTeamUsers(teamID string, userIDs []string) (*Team, []error) {
	body, errs := c.PutTeamUsersJSON(teamID, userIDs)
	if errs != nil {
		return nil, errs
	}

	team := Team{}
	json.Unmarshal([]byte(body), &team)

	return &team, nil
}
