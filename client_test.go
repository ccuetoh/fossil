package fossil

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

//***** Testing *****//

func TestClientCredentials_GetServers(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		// The response is provided by the Pterodactyl API Documentation. The meta.pagination.links parameter
		// has been modified from [] to {} since all analyzed responses do not, respond with an array but an
		// empty object. See: https://github.com/parkervcp/crocgodyl/issues/8
		res := `{
		   "object":"list",
		   "data":[
			  {
				 "object":"server",
				 "attributes":{
					"server_owner":true,
					"identifier":"d3aac109",
					"uuid":"d3aac109-e5a0-4331-b03e-3454f7e136dc",
					"name":"Survival",
					"description":"",
					"limits":{
					   "memory":1024,
					   "swap":0,
					   "disk":5000,
					   "io":500,
					   "cpu":200
					},
					"feature_limits":{
					   "databases":5,
					   "allocations":5
					}
				 }
			  }
		   ],
		   "meta":{
			  "pagination":{
				 "total":1,
				 "count":1,
				 "per_page":25,
				 "current_page":1,
				 "total_pages":1,
				 "links": { } 
			  }
		   }
		}`

		return []byte(res), nil
	}

	c := NewClient("", "")

	expect := []*ClientServer{
		{
			ID:          "d3aac109",
			Name:        "Survival",
			Description: "",
			Limits: Limits{
				Memory:      1024,
				Swap:        0,
				Disk:        5000,
				IO:          500,
				CPU:         200,
				Databases:   5,
				Allocations: 5,
			},
			IsOwner: true,
		},
	}

	got, err := c.GetServers()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if cmp.Equal(got, expect) {
		t.Error("Unexpected response")
	}
}

func TestClientCredentials_GetServer(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		res := `{
		   "object":"server",
		   "attributes":{
			  "server_owner":true,
			  "identifier":"d3aac109",
			  "uuid":"d3aac109-e5a0-4331-b03e-3454f7e136dc",
			  "name":"Survival",
			  "description":"",
			  "limits":{
				 "memory":1024,
				 "swap":0,
				 "disk":5000,
				 "io":500,
				 "cpu":200
			  },
			  "feature_limits":{
				 "databases":5,
				 "allocations":5
			  }
		   }
		}`

		return []byte(res), nil
	}

	c := NewClient("", "")

	expect := &ClientServer{
		ID:          "d3aac109",
		Name:        "Survival",
		Description: "",
		Limits: Limits{
			Memory:      1024,
			Swap:        0,
			Disk:        5000,
			IO:          500,
			CPU:         200,
			Databases:   5,
			Allocations: 5,
		},
		IsOwner: true,
	}

	got, err := c.GetServer("")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if cmp.Equal(got, expect) {
		t.Error("Unexpected response")
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

	if cmp.Equal(got, expect) {
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
