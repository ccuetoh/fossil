package fossil

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

//***** Testing *****//

func TestClientCredentials_GetServers(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		res := `{
				  "object": "list",
				  "data": [
					{
					  "object": "server",
					  "attributes": {
						"server_owner": true,
						"identifier": "1a7ce997",
						"uuid": "1a7ce997-259b-452e-8b4e-cecc464142ca",
						"name": "Wuhu Island",
						"node": "Test",
						"sftp_details": {
						  "ip": "pterodactyl.file.properties",
						  "port": 2022
						},
						"description": "Matt from Wii Sports",
						"allocation": {
						  "ip": "45.86.168.218",
						  "port": 25565
						},
						"limits": {
						  "memory": 512,
						  "swap": 0,
						  "disk": 200,
						  "io": 500,
						  "cpu": 0
						},
						"feature_limits": {
						  "databases": 5,
						  "allocations": 5,
						  "backups": 2
						},
						"is_suspended": false,
						"is_installing": false
					  }
					}
				  ],
				  "meta": {
					"pagination": {
					  "total": 1,
					  "count": 1,
					  "per_page": 15,
					  "current_page": 1,
					  "total_pages": 1,
					  "links": {}
					}
				  }
				}`

		return []byte(res), nil
	}

	c := NewClient("", "")

	expect := []*ClientServer{
		{
			ID:          "1a7ce997",
			UUID:        "1a7ce997-259b-452e-8b4e-cecc464142ca",
			Name:        "Wuhu Island",
			Description: "Matt from Wii Sports",
			IsOwner: true,
			IsInstalling: false,
			Node:        "Test",
			SFTPDetails: SFTPDetails{
				IP:   "pterodactyl.file.properties",
				Port: 2022,
			},
			Allocation: Allocation{
				IP:   "45.86.168.218",
				Port: 25565,
			},
			Limits: Limits{
				Memory:      512,
				Disk:        200,
				IO:          500,
				CPU:         0,
				Databases:   5,
				Allocations: 5,
				Backups: 2,
			},
		},
	}

	got, err := c.GetServers()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if !cmp.Equal(got, expect) {
		t.Error("Unexpected response")
	}
}

func TestClientCredentials_GetServer(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		res := `{
			  "object": "server",
			  "attributes": {
				"server_owner": true,
				"identifier": "1a7ce997",
				"uuid": "1a7ce997-259b-452e-8b4e-cecc464142ca",
				"name": "Wuhu Island",
				"node": "Test",
				"sftp_details": {
				  "ip": "pterodactyl.file.properties",
				  "port": 2022
				},
				"description": "Matt from Wii Sports",
				"limits": {
				  "memory": 512,
				  "swap": 0,
				  "disk": 200,
				  "io": 500,
				  "cpu": 0
				},
				"feature_limits": {
				  "databases": 5,
				  "allocations": 5,
				  "backups": 2
				},
				"is_suspended": false,
				"is_installing": false,
				"relationships": {
				  "allocations": {
					"object": "list",
					"data": [
					  {
						"object": "allocation",
						"attributes": {
						  "id": 1,
						  "ip": "45.86.168.218",
						  "ip_alias": null,
						  "port": 25565,
						  "notes": null,
						  "is_default": true
						}
					  },
					  {
						"object": "allocation",
						"attributes": {
						  "id": 2,
						  "ip": "45.86.168.218",
						  "ip_alias": null,
						  "port": 25566,
						  "notes": null,
						  "is_default": false
						}
					  }
					]
				  }
				}
			  },
			  "meta": {
				"is_server_owner": true,
				"user_permissions": [
				  "*",
				  "admin.websocket.errors",
				  "admin.websocket.install"
				]
			  }
			}`

		return []byte(res), nil
	}

	c := NewClient("", "")

	expect := &ClientServerDetail{
		ID:           "1a7ce997",
		UUID:         "1a7ce997-259b-452e-8b4e-cecc464142ca",
		Name:         "Wuhu Island",
		Description:  "Matt from Wii Sports",
		IsOwner:      true,
		IsInstalling: false,
		Node:         "Test",
		SFTPDetails: SFTPDetails{
			IP:   "pterodactyl.file.properties",
			Port: 2022,
		},
		Allocations: []*AllocationDetails{
			{
				ID:        1,
				IP:        "45.86.168.218",
				Port:      25565,
				IsDefault: true,
			},
			{
				ID:        2,
				IP:        "45.86.168.218",
				Port:      25566,
				IsDefault: false,
			},
		},
		Limits: Limits{
			Memory:      512,
			Disk:        200,
			IO:          500,
			CPU:         0,
			Databases:   5,
			Allocations: 5,
			Backups:     2,
		},
		Permissions: []string{"*", "admin.websocket.errors", "admin.websocket.install"},
	}

	got, err := c.GetServer("")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if !cmp.Equal(got, expect) {
		t.Errorf("Unexpected response: %s", cmp.Diff(got, expect))
	}
}

func TestClientCredentials_GetStatus(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		res := `{
		   "object":"stats",
		   "attributes":{
			  "state":"on",
			  "memory":{
				 "current":375,
				 "limit":1024
			  },
			  "cpu":{
				 "current":1.522,
				 "cores":[
					0.033,
					0.048,
					0.04,
					0,
					0.031,
					0,
					0.021,
					0.024,
					0.249,
					0.042,
					0.007,
					0,
					0.293,
					0.003,
					0.6,
					0.131
				 ],
				 "limit":200
			  },
			  "disk":{
				 "current":119,
				 "limit":5000
			  }
		   }
		}`

		return []byte(res), nil
	}

	c := NewClient("", "")

	expect := &ServerStatus{
		State: "on",
		Memory: Memory{
			Used:  375,
			Limit: 1024,
		},
		CPU: CPU{
			Current: 1.522,
			Cores: []float32{
				0.033,
				0.048,
				0.04,
				0,
				0.031,
				0,
				0.021,
				0.024,
				0.249,
				0.042,
				0.007,
				0,
				0.293,
				0.003,
				0.6,
				0.131,
			},
			Limit: 200,
		},
		Disk: Disk{
			Used:  119,
			Limit: 5000,
		},
		Players: Players{},
	}

	got, err := c.GetServerStatus("")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if !cmp.Equal(got, expect) {
		t.Error("Unexpected response")
	}
}

func TestClientCredentials_ExecuteCommand(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectBody := `{"command":"test"}`
		expectURL := "https://example.com/api/client/servers/test_id/command"

		if expectBody != string(data) {
			t.Errorf("Request data does not match expected: %s", string(data))
		}

		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		return nil, nil
	}

	c := NewClient("https://example.com", "")

	err := c.ExecuteCommand("test_id", "test")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestClientCredentials_SetPowerState(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectBody := `{"signal":"start"}`
		expectURL := "https://example.com/api/client/servers/test_id/power"

		if expectBody != string(data) {
			t.Errorf("Request data does not match expected: %s", string(data))
		}

		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", string(data))
		}

		return nil, nil
	}

	c := NewClient("https://example.com", "")

	err := c.SetPowerState("test_id", ON)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}