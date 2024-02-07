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
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetOrganization(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	c, _ := NewClient(&Config{RedashURI: "https://com.acme/", APIKey: "ApIkEyApIkEyApIkEyApIkEyApIkEy"})

	httpmock.RegisterResponder("GET", "https://com.acme/api/settings/organization",
		httpmock.NewStringResponder(200, `{"auth_saml_entity_id": "test", "auth_saml_metadata_url": "test", "auth_saml_nameid_format": "test"}`))

	organization, err := c.GetOrganization()
	if err != nil {
		panic(err.Error())
	}

	assert.Equal("test", organization.AuthSamlEntityId)
}