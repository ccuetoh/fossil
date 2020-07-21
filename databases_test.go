package fossil

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestApplicationCredentials_GetDatabases(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		res := `{
		"object": "list",
	  	"data": [
			{
			  "object": "server_database",
			  "attributes": {
				"id": 6,
				"server": 1,
				"host": 2,
				"database": "s1_test",
				"username": "u1_iff9TGoHFt",
				"remote": "%",
				"created_at": "2019-10-06T15:16:26+02:00",
				"updated_at": "2019-10-06T15:28:26+02:00"
			  }
			},
			{
			  "object": "server_database",
			  "attributes": {
				"id": 7,
				"server": 1,
				"host": 2,
				"database": "s1_db2",
				"username": "u1_tZ14uEvXan",
				"remote": "%",
				"created_at": "2019-10-06T15:28:57+02:00",
				"updated_at": "2019-10-06T15:28:57+02:00"
			  }
			}
		  ]
		}`

		return []byte(res), nil
	}

	a := NewApplication("", "")

	// u1, _ := time.Parse(time.RFC3339, "2019-10-06T15:28:26+02:00")
	// c1, _ := time.Parse(time.RFC3339, "2019-10-06T15:16:26+02:00")

	// u2, _ := time.Parse(time.RFC3339, "2019-10-06T15:28:57+02:00")
	// c2, _ := time.Parse(time.RFC3339, "2019-10-06T15:28:57+02:00")

	expect := []*Database{
		{
			ID:              "",
			Host:            nil,
			Name:            "",
			Username:        "",
			ConnectionsFrom: "",
			MaxConnections:  0,
		},
		{
			ID:              "",
			Host:            nil,
			Name:            "",
			Username:        "",
			ConnectionsFrom: "",
			MaxConnections:  0,
		},
	}

	got, err := a.GetDatabases(1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if !cmp.Equal(got, expect) {
		t.Errorf("Unexpected response: %s", cmp.Diff(got, expect))
	}
}

func TestApplicationCredentials_GetDatabase(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/application/servers/1/databases/1"
		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		res := `{
		  "object": "server_database",
		  "attributes": {
			"id": 6,
			"server": 1,
			"host": 2,
			"database": "s1_test",
			"username": "u1_iff9TGoHFt",
			"remote": "%",
			"created_at": "2019-10-06T15:16:26+02:00",
			"updated_at": "2019-10-06T15:28:26+02:00"
		  }
		}`

		return []byte(res), nil
	}

	a := NewApplication("https://example.com", "")

	// u1, _ := time.Parse(time.RFC3339, "2019-10-06T15:28:26+02:00")
	// c1, _ := time.Parse(time.RFC3339, "2019-10-06T15:16:26+02:00")

	expect := &Database{
		ID:              "",
		Host:            nil,
		Name:            "",
		Username:        "",
		ConnectionsFrom: "",
		MaxConnections:  0,
	}

	got, err := a.GetDatabase(1, 1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if !cmp.Equal(got, expect) {
		t.Errorf("Unexpected response: %s", cmp.Diff(got, expect))
	}
}
func TestApplicationCredentials_CreateDatabase(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/application/servers/1/databases"
		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		expectBody := `{"database":"mydb","remote":"%","host":2}`
		if expectBody != string(data) {
			t.Errorf("Request body does not match expected: %s", data)
		}

		return nil, nil
	}

	a := NewApplication("https://example.com", "")

	db := &Database{
		ID:              "",
		Host:            nil,
		Name:            "",
		Username:        "",
		ConnectionsFrom: "",
		MaxConnections:  0,
	}

	err := a.CreateDatabase(1, db)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestApplicationCredentials_ResetDatabasePassword(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/application/servers/1/databases/1/reset-password"
		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		return nil, nil
	}

	a := NewApplication("https://example.com", "")

	err := a.ResetDatabasePassword(1, 1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestApplicationCredentials_DeleteDatabase(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/application/servers/1/databases/1"
		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		return nil, nil
	}

	a := NewApplication("https://example.com", "")

	err := a.DeleteDatabase(1, 1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}
