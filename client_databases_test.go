package fossil

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

//***** Testing *****//

func TestClientCredentials_GetDatabases(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		res := `{
			  "object": "list",
			  "data": [
				{
				  "object": "server_database",
				  "attributes": {
					"id": "bEY4yAD5",
					"host": {
					  "address": "127.0.0.1",
					  "port": 3306
					},
					"name": "s5perms",
					"username": "u5QsIAp1jhvS",
					"connections_from": "%",
					"max_connections": 0
				  }
				},
				{
				  "object": "server_database",
				  "attributes": {
					"id": "E0A0Rw42",
					"host": {
					  "address": "127.0.0.1",
					  "port": 3306
					},
					"name": "s5coreprotect",
					"username": "u52jtJx1nO1d",
					"connections_from": "%",
					"max_connections": 0
				  }
				}
			  ]
			}`

		return []byte(res), nil
	}

	c := NewClient("", "")

	expect := []*Database{
		{
			ID: "bEY4yAD5",
			Host: &Host{
				Address: "127.0.0.1",
				Port:    3306,
			},
			Name:            "s5perms",
			Username:        "u5QsIAp1jhvS",
			ConnectionsFrom: "%",
			MaxConnections:  0,
		},
		{
			ID: "E0A0Rw42",
			Host: &Host{
				Address: "127.0.0.1",
				Port:    3306,
			},
			Name:            "s5coreprotect",
			Username:        "u52jtJx1nO1d",
			ConnectionsFrom: "%",
			MaxConnections:  0,
		},
	}

	got, err := c.GetDatabases(0)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if !cmp.Equal(got, expect) {
		t.Errorf("Unexpected response: %s", cmp.Diff(got, expect))
	}
}

func TestClientCredentials_CreateDatabase(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/client/servers/1a7ce997/databases"

		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		res := `{
			  "object": "serverdatabase",
			  "attributes": {
				"id": "y9YVxO4V",
				"host": {
				  "address": "127.0.0.1",
				  "port": 3306
				},
				"name": "s5punishments",
				"username": "u5aeZqbGdCM9",
				"connections_from": "%",
				"max_connections": 0,
				"relationships": {
				  "password": {
					"object": "databasepassword",
					"attributes": {
					  "password": "=lR2orDOcwfKkM=BXb.BVF.C"
					}
				  }
				}
			  }
			}`

		return []byte(res), nil
	}

	c := NewClient("https://example.com", "")

	expectDb := &Database{
		ID: "y9YVxO4V",
		Host: &Host{
			Address: "127.0.0.1",
			Port:    3306,
		},
		Name:            "s5punishments",
		Username:        "u5aeZqbGdCM9",
		ConnectionsFrom: "%",
		MaxConnections:  0,
	}

	expectPassword := "=lR2orDOcwfKkM=BXb.BVF.C"

	got, password, err := c.CreateDatabase("1a7ce997")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if expectPassword != password {
		t.Errorf("Unexpected password: %s", password)
	}

	if !cmp.Equal(got, expectDb) {
		t.Errorf("Unexpected response: %s", cmp.Diff(got, expectDb))
	}
}