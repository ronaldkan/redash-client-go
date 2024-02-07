//
// Copyright (c) 2020 Snowplow Analytics Ltd. All rights reserved.
//
// This program is licensed to you under the Apache License Version 2.0,
// and you may not use this file except in compliance with the Apache License Version 2.0.
// You may obtain a copy of the Apache License Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the Apache License Version 2.0 is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the Apache License Version 2.0 for the specific language governing permissions and limitations there under.
//

package redash

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
)

// Organization struct
type Organization struct {
	AuthPasswordLoginEnabled	bool   	`json:"auth_password_login_enabled,omitempty"`
	AuthSamlEnabled           	bool 	`json:"auth_saml_enabled,omitempty"`
	AuthSamlType              	string 	`json:"auth_saml_type,omitempty"`
	AuthSamlEntityId         	string 	`json:"auth_saml_entity_id,omitempty"`
	AuthSamlMetadataUrl      	string 	`json:"auth_saml_metadata_url,omitempty"`
	AuthSamlNameidFormat     	string 	`json:"auth_saml_nameid_format,omitempty"`
	AuthSamlSsoUrl           	string 	`json:"auth_saml_sso_url,omitempty"`
}

type Settings struct {
	AuthPasswordLoginEnabled	bool   	`json:"settings.auth_password_login_enabled,omitempty"`
	AuthSamlEnabled           	bool 	`json:"settings.auth_saml_enabled,omitempty"`
	AuthSamlType              	string 	`json:"settings.auth_saml_type,omitempty"`
	AuthSamlEntityId         	string 	`json:"settings.auth_saml_entity_id,omitempty"`
	AuthSamlMetadataUrl      	string 	`json:"settings.auth_saml_metadata_url,omitempty"`
	AuthSamlNameidFormat     	string 	`json:"settings.auth_saml_nameid_format,omitempty"`
	AuthSamlSsoUrl           	string 	`json:"settings.auth_saml_sso_url,omitempty"`
}


// GetOrganization returns the organization settings
func (c *Client) GetOrganization() (*Settings, error) {
	path := "/api/settings/organization"

	query := url.Values{}
	response, err := c.get(path, query)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	organization := Settings{}

	err = json.Unmarshal(body, &organization)
	if err != nil {
		return nil, err
	}

	return &organization, nil
}

// UpdateOrganization updates an existing Redash organization
func (c *Client) UpdateOrganization(organization *Organization) (*Organization, error) {
	path := "/api/settings/organization"

	payload, err := json.Marshal(organization)
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	response, err := c.post(path, string(payload), query)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if body != nil && err != nil {
		return nil, err
	}

	return organization, nil
}
