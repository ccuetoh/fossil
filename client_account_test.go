package fossil

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

//***** Testing *****//

func TestClientCredentials_WhoAmI(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		res := `{
			  "object": "user",
			  "attributes": {
				"id": 1,
				"admin": true,
				"username": "admin",
				"email": "example@example.com",
				"first_name": "RootAdmin",
				"last_name": "User",
				"language": "en"
			  }
			}`

		return []byte(res), nil
	}

	c := NewClient("", "")

	expect := &User{
		ID:        1,
		Username:  "admin",
		Email:     "example@example.com",
		FirstName: "RootAdmin",
		LastName:  "User",
		Language:  "en",
		RootAdmin: true,
	}

	got, err := c.WhoAmI()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if cmp.Equal(got, expect) {
		t.Error("Unexpected response")
	}
}

func TestClientCredentials_Get2FAImageURL(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		res := `{
				  "data": {
					"image_url_data": "otpauth:\/\/totp\/Pterodactyl:example%40example.com?secret=LGYOWJEGVRPPGPWATP5ZHOYC7DHAYQ6S&issuer=Pterodactyl"
				  }
				}`

		return []byte(res), nil
	}

	c := NewClient("", "")

	expect := `otpauth://totp/Pterodactyl:example%40example.com?secret=LGYOWJEGVRPPGPWATP5ZHOYC7DHAYQ6S&issuer=Pterodactyl`

	got, err := c.Get2FAImageURL()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if !cmp.Equal(got, expect) {
		t.Error("Unexpected response")
	}
}

func TestClientCredentials_Enable2FA(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectBody := `{"code":"testcode"}`
		expectURL := "https://example.com/api/client/account/two-factor"

		if expectBody != string(data) {
			t.Errorf("Request data does not match expected: %s", string(data))
		}

		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		return nil, nil
	}

	c := NewClient("https://example.com", "")

	err := c.Enable2FA("testcode")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestClientCredentials_Disable2FA(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectBody := `{"password":"testpassword"}`
		expectURL := "https://example.com/api/client/account/two-factor"

		if expectBody != string(data) {
			t.Errorf("Request data does not match expected: %s", string(data))
		}

		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		return nil, nil
	}

	c := NewClient("https://example.com", "")

	err := c.Disable2FA("testpassword")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestClientCredentials_UpdateEmail(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectBody := `{"email":"testemail","password":"testpassword"}`
		expectURL := "https://example.com/api/client/account/email"

		if expectBody != string(data) {
			t.Errorf("Request data does not match expected: %s", string(data))
		}

		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		return nil, nil
	}

	c := NewClient("https://example.com", "")

	err := c.UpdateEmail("testemail", "testpassword")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestClientCredentials_UpdatePassword(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectBody := `{"current_password":"testoldpass","password":"testnewpass","password_confirmation":"testnewpass"}`
		expectURL := "https://example.com/api/client/account/password"

		if expectBody != string(data) {
			t.Errorf("Request data does not match expected: %s", string(data))
		}

		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		return nil, nil
	}

	c := NewClient("https://example.com", "")

	err := c.UpdatePassword("testoldpass", "testnewpass")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

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

func TestClientCredentials_DeleteAPIKey(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/client/account/api-keys/NWKMYMT2Mrav0Iq2"
		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		return nil, nil
	}

	a := NewClient("https://example.com", "")

	err := a.DeleteAPIKey("NWKMYMT2Mrav0Iq2")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}
