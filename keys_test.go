package fossil

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

//***** Testing *****//

func TestClientCredentials_GetAPIKeys(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		res := `{
				  "object": "list",
				  "data": [
					{
					  "object": "apikey",
					  "attributes": {
						"identifier": "wwQ5DJ6X1XaFznQS",
						"description": "API Docs",
						"allowed_ips": [],
						"last_used_at": "2020-06-03T15:04:47+01:00",
						"created_at": "2020-05-18T00:12:43+01:00"
					  }
					}
				  ]
				}`

		return []byte(res), nil
	}

	c := NewClient("", "")

	used, _ := time.Parse(time.RFC3339, "2020-06-03T15:04:47+01:00")
	created, _ := time.Parse(time.RFC3339, "2020-05-18T00:12:43+01:00")

	expect := []*APIKey{
		{
			Identifier:  "wwQ5DJ6X1XaFznQS",
			Description: "API Docs",
			AllowedIPs:  []string{},
			LastUsed:    used,
			Created:     created,
		},
	}

	got, err := c.GetAPIKeys()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if !cmp.Equal(got, expect) {
		t.Error("Unexpected response")
	}
}

func TestClientCredentials_NewAPIKey(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectBody := `{"description":"testdescription"}`
		expectURL := "https://example.com/api/client/account/api-keys"

		if expectBody != string(data) {
			t.Errorf("Request data does not match expected: %s", string(data))
		}

		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		res := `{
				  "object": "apikey",
				  "attributes": {
					"identifier": "NWKMYMT2Mrav0Iq2",
					"description": "Created API key through API",
					"allowed_ips": [],
					"last_used_at": null,
					"created_at": "2020-07-10T05:21:11+01:00"
				  },
				  "meta": {
					"secret_token": "xa1tpfptoCM1ynFnTgSpxxukwkYC1LTf"
				  }
				}`

		return []byte(res), nil
	}

	c := NewClient("https://example.com", "")

	created, _ := time.Parse(time.RFC3339, "2020-07-10T05:21:11+01:00")

	expect := &APIKey{
		Identifier:  "NWKMYMT2Mrav0Iq2",
		Description: "Created API key through API",
		AllowedIPs:  []string{},
		LastUsed:    time.Time{},
		Created:     created,
	}

	gotKey, gotSecret,  err := c.NewAPIKey("testdescription")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if gotSecret != "xa1tpfptoCM1ynFnTgSpxxukwkYC1LTf"{
		t.Error(fmt.Sprintf("Unexpected response for secret: %s", gotSecret))
	}

	if !cmp.Equal(gotKey, expect) {
		t.Error("Unexpected response for APIKey object")
	}
}