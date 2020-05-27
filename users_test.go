package fossil

import (
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

func TestApplicationCredentials_GetUsers(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectUrl := "https://example.com/api/application/users"
		if expectUrl != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		// The response is provided by the Pterodactyl API Documentation. The meta.pagination.links parameter
		// has been modified from [] to {} since all analyzed responses do not, respond with an array but an
		// empty object. See: https://github.com/parkervcp/crocgodyl/issues/8
		res := `{
		  "object": "list",
		  "data": [
			{
			  "object": "user",
			  "attributes": {
				"id": 1,
				"external_id": null,
				"uuid": "c4022c6c-9bf1-4a23-bff9-519cceb38335",
				"username": "codeco",
				"email": "codeco@file.properties",
				"first_name": "Rihan",
				"last_name": "Arfan",
				"language": "en",
				"root_admin": true,
				"2fa": false,
				"created_at": "2018-03-18T15:15:17+00:00",
				"updated_at": "2018-10-16T21:51:21+00:00"
			  }
			},
			{
			  "object": "user",
			  "attributes": {
				"id": 4,
				"external_id": null,
				"uuid": "f253663c-5a45-43a8-b280-3ea3c752b931",
				"username": "wardledeboss",
				"email": "wardle315@gmail.com",
				"first_name": "Harvey",
				"last_name": "Wardle",
				"language": "en",
				"root_admin": false,
				"2fa": false,
				"created_at": "2018-09-29T17:59:45+00:00",
				"updated_at": "2018-10-02T18:59:03+00:00"
			  }
			},
			{
			  "object": "user",
			  "attributes": {
				"id": 5,
				"external_id": null,
				"uuid": "0d8da9a5-6ccd-4b57-9786-70a97a1a55e7",
				"username": "matthewp",
				"email": "me@matthewp.io",
				"first_name": "Matthew",
				"last_name": "Penner",
				"language": "en",
				"root_admin": true,
				"2fa": false,
				"created_at": "2018-09-29T22:39:05+00:00",
				"updated_at": "2018-09-29T22:39:27+00:00"
			  }
			},
			{
			  "object": "user",
			  "attributes": {
				"id": 6,
				"external_id": null,
				"uuid": "d006fe91-3c64-4b0c-81d1-718af2cc384e",
				"username": "rihan554rnk",
				"email": "rihan554@gmail.com",
				"first_name": "Server",
				"last_name": "Subuser",
				"language": "en",
				"root_admin": false,
				"2fa": false,
				"created_at": "2018-10-02T21:26:18+00:00",
				"updated_at": "2018-10-02T21:26:18+00:00"
			  }
			}
		  ],
		  "meta": {
			"pagination": {
			  "total": 4,
			  "count": 4,
			  "per_page": 50,
			  "current_page": 1,
			  "total_pages": 1,
			  "links": {}
			}
		  }
		}`

		return []byte(res), nil
	}

	a := NewApplication("https://example.com", "")

	u1, _ := time.Parse(time.RFC3339, "2018-10-16T21:51:21+00:00")
	c1, _ := time.Parse(time.RFC3339, "2018-03-18T15:15:17+00:00")

	u2, _ := time.Parse(time.RFC3339, "2018-10-02T18:59:03+00:00")
	c2, _ := time.Parse(time.RFC3339, "2018-09-29T17:59:45+00:00")

	u3, _ := time.Parse(time.RFC3339, "2018-09-29T22:39:27+00:00")
	c3, _ := time.Parse(time.RFC3339, "2018-09-29T22:39:05+00:00")

	u4, _ := time.Parse(time.RFC3339, "2018-10-02T21:26:18+00:00")
	c4, _ := time.Parse(time.RFC3339, "2018-10-02T21:26:18+00:00")

	expect := []*User{
		{
			ID:                      1,
			ExternalID:              "",
			UUID:                    "c4022c6c-9bf1-4a23-bff9-519cceb38335",
			Username:                "codeco",
			Email:                   "codeco@file.properties",
			FirstName:               "Rihan",
			LastName:                "Arfan",
			Language:                "en",
			RootAdmin:               true,
			TwoFactorAuthentication: false,
			CreatedAt:               c1,
			UpdatedAt:               u1,
		},
		{
			ID:                      4,
			ExternalID:              "",
			UUID:                    "f253663c-5a45-43a8-b280-3ea3c752b931",
			Username:                "wardledeboss",
			Email:                   "wardle315@gmail.com",
			FirstName:               "Harvey",
			LastName:                "Wardle",
			Language:                "en",
			RootAdmin:               false,
			TwoFactorAuthentication: false,
			CreatedAt:               c2,
			UpdatedAt:               u2,
		},
		{
			ID:                      5,
			ExternalID:              "",
			UUID:                    "0d8da9a5-6ccd-4b57-9786-70a97a1a55e7",
			Username:                "matthewp",
			Email:                   "me@matthewp.io",
			FirstName:               "Matthew",
			LastName:                "Penner",
			Language:                "en",
			RootAdmin:               true,
			TwoFactorAuthentication: false,
			CreatedAt:               c3,
			UpdatedAt:               u3,
		},
		{
			ID:                      6,
			ExternalID:              "",
			UUID:                    "d006fe91-3c64-4b0c-81d1-718af2cc384e",
			Username:                "rihan554rnk",
			Email:                   "rihan554@gmail.com",
			FirstName:               "Server",
			LastName:                "Subuser",
			Language:                "en",
			RootAdmin:               false,
			TwoFactorAuthentication: false,
			CreatedAt:               c4,
			UpdatedAt:               u4,
		},
	}

	got, err := a.GetUsers()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if !cmp.Equal(got, expect) {
		t.Errorf("Unexpected response: %s", cmp.Diff(got, expect))
	}
}

func TestApplicationCredentials_GetUser(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectUrl := "https://example.com/api/application/users/1"
		if expectUrl != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		res := `{
		  "object": "user",
		  "attributes": {
			"id": 1,
			"external_id": null,
			"uuid": "c4022c6c-9bf1-4a23-bff9-519cceb38335",
			"username": "codeco",
			"email": "codeco@file.properties",
			"first_name": "Rihan",
			"last_name": "Arfan",
			"language": "en",
			"root_admin": true,
			"2fa": false,
			"created_at": "2018-03-18T15:15:17+00:00",
			"updated_at": "2018-10-16T21:51:21+00:00"
		  }
		}`

		return []byte(res), nil
	}

	a := NewApplication("https://example.com", "")

	u, _ := time.Parse(time.RFC3339, "2018-10-16T21:51:21+00:00")
	c, _ := time.Parse(time.RFC3339, "2018-03-18T15:15:17+00:00")

	expect := &User{
		ID:                      1,
		ExternalID:              "",
		UUID:                    "c4022c6c-9bf1-4a23-bff9-519cceb38335",
		Username:                "codeco",
		Email:                   "codeco@file.properties",
		FirstName:               "Rihan",
		LastName:                "Arfan",
		Language:                "en",
		RootAdmin:               true,
		TwoFactorAuthentication: false,
		CreatedAt:               c,
		UpdatedAt:               u,
	}

	got, err := a.GetUser(1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if !cmp.Equal(got, expect) {
		t.Errorf("Unexpected response: %s", cmp.Diff(got, expect))
	}
}

func TestApplicationCredentials_GetUserExternal(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectUrl := "https://example.com/api/application/users/external/1"
		if expectUrl != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		res := `{
		  "object": "user",
		  "attributes": {
			"id": 1,
			"external_id": "1",
			"uuid": "c4022c6c-9bf1-4a23-bff9-519cceb38335",
			"username": "codeco",
			"email": "codeco@file.properties",
			"first_name": "Rihan",
			"last_name": "Arfan",
			"language": "en",
			"root_admin": true,
			"2fa": false,
			"created_at": "2018-03-18T15:15:17+00:00",
			"updated_at": "2018-10-16T21:51:21+00:00"
		  }
		}`

		return []byte(res), nil
	}

	a := NewApplication("https://example.com", "")

	u, _ := time.Parse(time.RFC3339, "2018-10-16T21:51:21+00:00")
	c, _ := time.Parse(time.RFC3339, "2018-03-18T15:15:17+00:00")

	expect := &User{
		ID:                      1,
		ExternalID:              "1",
		UUID:                    "c4022c6c-9bf1-4a23-bff9-519cceb38335",
		Username:                "codeco",
		Email:                   "codeco@file.properties",
		FirstName:               "Rihan",
		LastName:                "Arfan",
		Language:                "en",
		RootAdmin:               true,
		TwoFactorAuthentication: false,
		CreatedAt:               c,
		UpdatedAt:               u,
	}

	got, err := a.GetUserExternal("1")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if !cmp.Equal(got, expect) {
		t.Errorf("Unexpected response: %s", cmp.Diff(got, expect))
	}
}

func TestApplicationCredentials_CreateUser(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectUrl := "https://example.com/api/application/users"
		if expectUrl != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		expectBody := `{"external_id":"example_ext_id","username":"example",` +
			`"email":"example@example.com","first_name":"John",` +
			`"last_name":"Doe","password":"cat","root_admin":false,"language":"en"}`
		if expectBody != string(data) {
			t.Errorf("Request data does not match expected: %s", string(data))
		}

		return nil, nil
	}

	a := NewApplication("https://example.com", "")

	user := &User{
		ExternalID: "example_ext_id",
		Username:   "example",
		Email:      "example@example.com",
		FirstName:  "John",
		LastName:   "Doe",
		Language:   "en",
		RootAdmin:  false,
	}

	err := a.CreateUser(user, "cat")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestApplicationCredentials_UpdateUser(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectUrl := "https://example.com/api/application/users/1"
		if expectUrl != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		expectBody := `{"external_id":"codeco","username":"codeco","email":"codeco@file.properties",` +
			`"first_name":"Updated","last_name":"User","password":"betterPassword","root_admin":true,"language":"en"}`
		if expectBody != string(data) {
			t.Errorf("Request data does not match expected: %s", string(data))
		}

		return nil, nil
	}

	a := NewApplication("https://example.com", "")

	user := &User{
		ID:         1,
		ExternalID: "codeco",
		Username:   "codeco",
		Email:      "codeco@file.properties",
		FirstName:  "Updated",
		LastName:   "User",
		Language:   "en",
		RootAdmin:  true,
	}

	err := a.UpdateUser(user, "betterPassword")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestApplicationCredentials_DeleteUser(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectUrl := "https://example.com/api/application/users/1"
		if expectUrl != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		return nil, nil
	}

	a := NewApplication("https://example.com", "")

	err := a.DeleteUser(1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}
