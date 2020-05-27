package fossil

import (
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

//***** Testing *****//

func TestApplicationCredentials_GetServers(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/application/servers?include=allocations"
		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		// The response is provided by the Pterodactyl API Documentation. The meta.pagination.links parameter
		// has been modified from [] to {} since all analyzed responses do not, respond with an array but an
		// empty object. See: https://github.com/parkervcp/crocgodyl/issues/8
		res := `{
	  "object": "list",
	  "data": [
		{
		  "object": "server",
		  "attributes": {
			"id": 2,
			"external_id": null,
			"uuid": "47a7052b-f07e-4845-989d-e876e30960f4",
			"identifier": "47a7052b",
			"name": "Eat Cows",
			"description": "",
			"suspended": false,
			"limits": {
			  "memory": 2048,
			  "swap": -1,
			  "disk": 10000,
			  "io": 500,
			  "cpu": 300
			},
			"feature_limits": {
			  "databases": 10,
			  "allocations": 0
			},
			"user": 1,
			"node": 2,
			"allocation": 3,
			"nest": 1,
			"egg": 4,
			"pack": null,
			"container": {
			  "startup_command": "java -Xms128M -Xmx{{SERVER_MEMORY}}M -jar {{SERVER_JARFILE}}",
			  "image": "quay.io/pterodactyl/core:java",
			  "installed": true,
			  "environment": {
				"SERVER_JARFILE": "server.jar",
				"VANILLA_VERSION": "latest",
				"STARTUP": "java -Xms128M -Xmx{{SERVER_MEMORY}}M -jar {{SERVER_JARFILE}}",
				"P_SERVER_LOCATION": "test",
				"P_SERVER_UUID": "47a7052b-f07e-4845-989d-e876e30960f4"
			  }
			},
			"updated_at": "2018-11-20T14:35:00+00:00",
			"created_at": "2018-09-29T22:50:16+00:00"
		  }
		},
		{
		  "object": "server",
		  "attributes": {
			"id": 6,
			"external_id": null,
			"uuid": "6d1567c5-08d4-4ecb-8d5d-0ce1ba6b0b99",
			"identifier": "6d1567c5",
			"name": "Wow",
			"description": "t",
			"suspended": false,
			"limits": {
			  "memory": 0,
			  "swap": -1,
			  "disk": 5000,
			  "io": 500,
			  "cpu": 200
			},
			"feature_limits": {
			  "databases": 0,
			  "allocations": 0
			},
			"user": 5,
			"node": 2,
			"allocation": 4,
			"nest": 1,
			"egg": 15,
			"pack": null,
			"container": {
			  "startup_command": "./parkertron",
			  "image": "quay.io/parkervcp/pterodactyl-images:parkertron",
			  "installed": true,
			  "environment": {
				"STARTUP": "./parkertron",
				"P_SERVER_LOCATION": "test",
				"P_SERVER_UUID": "6d1567c5-08d4-4ecb-8d5d-0ce1ba6b0b99"
			  }
			},
			"updated_at": "2018-11-10T19:52:13+00:00",
			"created_at": "2018-11-10T19:51:23+00:00"
		  }
		}
	  ],
	  "meta": {
		"pagination": {
		  "total": 2,
		  "count": 2,
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

	u1, _ := time.Parse(time.RFC3339, "2018-11-20T14:35:00+00:00")
	c1, _ := time.Parse(time.RFC3339, "2018-09-29T22:50:16+00:00")

	u2, _ := time.Parse(time.RFC3339, "2018-11-10T19:52:13+00:00")
	c2, _ := time.Parse(time.RFC3339, "2018-11-10T19:51:23+00:00")

	expect := []*ApplicationServer{
		{
			ID:          2,
			Name:        "Eat Cows",
			Description: "",
			Limits: Limits{
				Memory:      2048,
				Swap:        -1,
				Disk:        10000,
				IO:          500,
				CPU:         300,
				Databases:   10,
				Allocations: 0,
			},
			ExternalID: "",
			UUID:       "47a7052b-f07e-4845-989d-e876e30960f4",
			Suspended:  false,
			User:       1,
			Node:       2,
			Nest:       1,
			Egg:        4,
			Pack:       0,
			Allocation: 3,
			Container: Container{
				StartupCommand: "java -Xms128M -Xmx{{SERVER_MEMORY}}M -jar {{SERVER_JARFILE}}",
				Image:          "quay.io/pterodactyl/core:java",
				Installed:      true,
				Environment: map[string]string{
					"SERVER_JARFILE":    "server.jar",
					"VANILLA_VERSION":   "latest",
					"STARTUP":           "java -Xms128M -Xmx{{SERVER_MEMORY}}M -jar {{SERVER_JARFILE}}",
					"P_SERVER_LOCATION": "test",
					"P_SERVER_UUID":     "47a7052b-f07e-4845-989d-e876e30960f4",
				},
			},
			Updated: u1,
			Created: c1,
		},
		{
			ID:          6,
			Name:        "Wow",
			Description: "t",
			Limits: Limits{
				Memory:      0,
				Swap:        -1,
				Disk:        5000,
				IO:          500,
				CPU:         200,
				Databases:   0,
				Allocations: 0,
			},
			ExternalID: "",
			UUID:       "6d1567c5-08d4-4ecb-8d5d-0ce1ba6b0b99",
			Suspended:  false,
			User:       5,
			Node:       2,
			Nest:       1,
			Egg:        15,
			Pack:       0,
			Allocation: 4,
			Container: Container{
				StartupCommand: "./parkertron",
				Image:          "quay.io/parkervcp/pterodactyl-images:parkertron",
				Installed:      true,
				Environment: map[string]string{
					"STARTUP":           "./parkertron",
					"P_SERVER_LOCATION": "test",
					"P_SERVER_UUID":     "6d1567c5-08d4-4ecb-8d5d-0ce1ba6b0b99",
				},
			},
			Updated: u2,
			Created: c2,
		},
	}

	got, err := a.GetServers()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if !cmp.Equal(got, expect) {
		t.Errorf("Unexpected response: %s", cmp.Diff(got, expect))
	}
}

func TestApplicationCredentials_GetServer(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/application/servers/2?include=allocations"
		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		res := `{
		  "object": "server",
		  "attributes": {
			"id": 2,
			"external_id": null,
			"uuid": "47a7052b-f07e-4845-989d-e876e30960f4",
			"identifier": "47a7052b",
			"name": "Survival",
			"description": "gsk;ljgkj;hgdakl;gha",
			"suspended": false,
			"limits": {
			  "memory": 2048,
			  "swap": -1,
			  "disk": 10000,
			  "io": 500,
			  "cpu": 300
			},
			"feature_limits": {
			  "databases": 10,
			  "allocations": 0
			},
			"user": 1,
			"node": 2,
			"allocation": 3,
			"nest": 1,
			"egg": 4,
			"pack": null,
			"container": {
			  "startup_command": "java -Xms128M -Xmx{{SERVER_MEMORY}}M -jar {{SERVER_JARFILE}}",
			  "image": "quay.io/pterodactyl/core:java",
			  "installed": true,
			  "environment": {
				"SERVER_JARFILE": "server.jar",
				"VANILLA_VERSION": "latest",
				"STARTUP": "java -Xms128M -Xmx{{SERVER_MEMORY}}M -jar {{SERVER_JARFILE}}",
				"P_SERVER_LOCATION": "test",
				"P_SERVER_UUID": "47a7052b-f07e-4845-989d-e876e30960f4"
			  }
			},
			"updated_at": "2018-11-20T02:52:37+00:00",
			"created_at": "2018-09-29T22:50:16+00:00"
		  }
		}`

		return []byte(res), nil
	}

	a := NewApplication("https://example.com", "")

	u1, _ := time.Parse(time.RFC3339, "2018-11-20T02:52:37+00:00")
	c1, _ := time.Parse(time.RFC3339, "2018-09-29T22:50:16+00:00")

	expect := &ApplicationServer{
		ID:          2,
		Name:        "Survival",
		Description: "gsk;ljgkj;hgdakl;gha",
		Limits: Limits{
			Memory:      2048,
			Swap:        -1,
			Disk:        10000,
			IO:          500,
			CPU:         300,
			Databases:   10,
			Allocations: 0,
		},
		ExternalID: "",
		UUID:       "47a7052b-f07e-4845-989d-e876e30960f4",
		Suspended:  false,
		User:       1,
		Node:       2,
		Nest:       1,
		Egg:        4,
		Pack:       0,
		Allocation: 3,
		Container: Container{
			StartupCommand: "java -Xms128M -Xmx{{SERVER_MEMORY}}M -jar {{SERVER_JARFILE}}",
			Image:          "quay.io/pterodactyl/core:java",
			Installed:      true,
			Environment: map[string]string{
				"SERVER_JARFILE":    "server.jar",
				"VANILLA_VERSION":   "latest",
				"STARTUP":           "java -Xms128M -Xmx{{SERVER_MEMORY}}M -jar {{SERVER_JARFILE}}",
				"P_SERVER_LOCATION": "test",
				"P_SERVER_UUID":     "47a7052b-f07e-4845-989d-e876e30960f4",
			},
		},
		Updated: u1,
		Created: c1,
	}

	got, err := a.GetServer(2)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if !cmp.Equal(got, expect) {
		t.Errorf("Unexpected response: %s", cmp.Diff(got, expect))
	}
}

func TestApplicationCredentials_GetServerExternal(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/application/servers/external/cow_eater?include=allocations"
		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		res := `{
		  "object": "server",
		  "attributes": {
			"id": 7,
			"external_id": "cow_eater",
			"uuid": "78165af0-4835-405f-b281-07e961bfd0ad",
			"identifier": "78165af0",
			"name": "Eat Cows",
			"description": "ok....",
			"suspended": false,
			"limits": {
			  "memory": 1024,
			  "swap": 0,
			  "disk": 1000,
			  "io": 500,
			  "cpu": 0
			},
			"feature_limits": {
			  "databases": 0,
			  "allocations": 0
			},
			"user": 1,
			"node": 2,
			"allocation": 9,
			"nest": 1,
			"egg": 4,
			"pack": null,
			"container": {
			  "startup_command": "java -Xms128M -Xmx{{SERVER_MEMORY}}M -jar {{SERVER_JARFILE}}",
			  "image": "quay.io/pterodactyl/core:java",
			  "installed": true,
			  "environment": {
				"SERVER_JARFILE": "server.jar",
				"VANILLA_VERSION": "latest",
				"STARTUP": "java -Xms128M -Xmx{{SERVER_MEMORY}}M -jar {{SERVER_JARFILE}}",
				"P_SERVER_LOCATION": "test",
				"P_SERVER_UUID": "78165af0-4835-405f-b281-07e961bfd0ad"
			  }
			},
			"updated_at": "2018-12-17T20:00:16+00:00",
			"created_at": "2018-12-11T21:56:00+00:00"
		  }
		}`

		return []byte(res), nil
	}

	a := NewApplication("https://example.com", "")

	u1, _ := time.Parse(time.RFC3339, "2018-12-17T20:00:16+00:00")
	c1, _ := time.Parse(time.RFC3339, "2018-12-11T21:56:00+00:00")

	expect := &ApplicationServer{
		ID:          7,
		Name:        "Eat Cows",
		Description: "ok....",
		Limits: Limits{
			Memory:      1024,
			Swap:        0,
			Disk:        1000,
			IO:          500,
			CPU:         0,
			Databases:   0,
			Allocations: 0,
		},
		ExternalID: "cow_eater",
		UUID:       "78165af0-4835-405f-b281-07e961bfd0ad",
		Suspended:  false,
		User:       1,
		Node:       2,
		Nest:       1,
		Egg:        4,
		Pack:       0,
		Allocation: 9,
		Container: Container{
			StartupCommand: "java -Xms128M -Xmx{{SERVER_MEMORY}}M -jar {{SERVER_JARFILE}}",
			Image:          "quay.io/pterodactyl/core:java",
			Installed:      true,
			Environment: map[string]string{
				"SERVER_JARFILE":    "server.jar",
				"VANILLA_VERSION":   "latest",
				"STARTUP":           "java -Xms128M -Xmx{{SERVER_MEMORY}}M -jar {{SERVER_JARFILE}}",
				"P_SERVER_LOCATION": "test",
				"P_SERVER_UUID":     "78165af0-4835-405f-b281-07e961bfd0ad",
			},
		},
		Updated: u1,
		Created: c1,
	}

	got, err := a.GetServerExternal("cow_eater")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if !cmp.Equal(got, expect) {
		t.Errorf("Unexpected response: %s", cmp.Diff(got, expect))
	}
}

func TestApplicationCredentials_CreateServer(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/application/servers"
		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		return nil, nil
	}

	a := NewApplication("https://example.com", "")

	err := a.CreateServer(&ApplicationServer{})
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestApplicationCredentials_UpdateDetails(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/application/servers/1/details"
		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		// The original request parsed the user as string
		expectBody := `{"external_id":"cow_eater","name":"Eat Cows","user":1,"description":"ok...."}`

		if expectBody != string(data) {
			t.Errorf("Request data does not match expected: %s", string(data))
		}

		return nil, nil
	}

	a := NewApplication("https://example.com", "")

	sv := &ApplicationServer{
		ID:          1,
		Name:        "Eat Cows",
		Description: "ok....",
		ExternalID:  "cow_eater",
		User:        1,
	}

	err := a.UpdateDetails(sv)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestApplicationCredentials_UpdateBuild(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/application/servers/1/build"
		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		expectBody := `{"allocation":9,"oom_disabled":true,` +
			`"limits":{"memory":2048,"swap":-1,` +
			`"disk":10000,"io":500,"cpu":300},` +
			`"add_allocations":[15],"remove_allocations":[3],` +
			`"feature_limits":{"databases":10,"allocations":10}}`

		if expectBody != string(data) {
			t.Errorf("Request data does not match expected: %s", string(data))
		}

		return nil, nil
	}

	a := NewApplication("https://example.com", "")

	sv := &ApplicationServer{
		ID:         1,
		Allocation: 9,
		Limits: Limits{
			Memory:      2048,
			Swap:        -1,
			Disk:        10000,
			IO:          500,
			CPU:         300,
			Databases:   10,
			Allocations: 10,
		},
	}

	err := a.UpdateBuild(sv, []int{15}, []int{3})
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestApplicationCredentials_UpdateStartup(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/application/servers/1/startup"
		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		// Modified a bit from the original API example. The Map object in the request originally had 2 items
		// but would fail the test if they swapped order.
		expectBody := `{"startup":"java -Xms128M -Xmx1024M -jar paper.jar",` +
			`"environment":["SERVER_JARFILE"],"egg":1,` +
			`"pack":4,"image":"quay.io/pterodactyl/core:java","skip_scripts":false}`

		if expectBody != string(data) {
			t.Errorf("Request data does not match expected: %s", string(data))
		}

		return nil, nil
	}

	a := NewApplication("https://example.com", "")

	sv := &ApplicationServer{
		ID: 1,
		Container: Container{
			StartupCommand: "java -Xms128M -Xmx1024M -jar paper.jar",
			Image:          "quay.io/pterodactyl/core:java",
			Environment:    map[string]string{"SERVER_JARFILE": ""},
		},
		Egg:  1,
		Pack: 4,
	}

	err := a.UpdateStartup(sv)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestApplicationCredentials_SuspendServer(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/application/servers/1/suspend"
		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		return nil, nil
	}

	a := NewApplication("https://example.com", "")

	err := a.SuspendServer(1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestApplicationCredentials_UnsuspendServer(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/application/servers/1/unsuspend"
		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		return nil, nil
	}

	a := NewApplication("https://example.com", "")

	err := a.UnsuspendServer(1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestApplicationCredentials_ReinstallServer(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/application/servers/1/reinstall"
		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		return nil, nil
	}

	a := NewApplication("https://example.com", "")

	err := a.ReinstallServer(1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestApplicationCredentials_RebuildServer(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/application/servers/1/rebuild"
		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		return nil, nil
	}

	a := NewApplication("https://example.com", "")

	err := a.RebuildServer(1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestApplicationCredentials_DeleteServer(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/application/servers/1"
		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		return nil, nil
	}

	a := NewApplication("https://example.com", "")

	err := a.DeleteServer(1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestApplicationCredentials_ForceDeleteServer(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/application/servers/1/force"
		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		return nil, nil
	}

	a := NewApplication("https://example.com", "")

	err := a.ForceDeleteServer(1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}
