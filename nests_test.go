package fossil

import (
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

//***** Testing *****//

func TestApplicationCredentials_GetNests(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/application/nests"
		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		res := `{
		  "object": "list",
		  "data": [
			{
			  "object": "nest",
			  "attributes": {
				"id": 1,
				"uuid": "179a06c9-b5bf-4798-8c0e-9f78e8f1f67f",
				"author": "support@pterodactyl.io",
				"name": "Minecraft",
				"description": "Minecraft - the classic game from Mojang. With support for Vanilla MC, Spigot, and many others!",
				"created_at": "2018-03-18T15:14:37+00:00",
				"updated_at": "2018-03-18T15:14:37+00:00"
			  }
			},
			{
			  "object": "nest",
			  "attributes": {
				"id": 2,
				"uuid": "c389ab01-8b30-4c0f-b2a2-01e79b3611d2",
				"author": "support@pterodactyl.io",
				"name": "Source Engine",
				"description": "Includes support for most Source Dedicated Server games.",
				"created_at": "2018-03-18T15:14:37+00:00",
				"updated_at": "2018-03-18T15:14:37+00:00"
			  }
			},
			{
			  "object": "nest",
			  "attributes": {
				"id": 3,
				"uuid": "8b47c310-ff04-4c10-84c8-133e63f1bf61",
				"author": "support@pterodactyl.io",
				"name": "Voice Servers",
				"description": "Voice servers such as Mumble and Teamspeak 3.",
				"created_at": "2018-03-18T15:14:37+00:00",
				"updated_at": "2018-03-18T15:14:37+00:00"
			  }
			},
			{
			  "object": "nest",
			  "attributes": {
				"id": 4,
				"uuid": "f92cd279-a916-4ede-84ee-900f51048e31",
				"author": "support@pterodactyl.io",
				"name": "Rust",
				"description": "Rust - A game where you must fight to survive.",
				"created_at": "2018-03-18T15:14:37+00:00",
				"updated_at": "2018-03-18T15:14:37+00:00"
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

	u1, _ := time.Parse(time.RFC3339, "2018-03-18T15:14:37+00:00")
	c1, _ := time.Parse(time.RFC3339, "2018-03-18T15:14:37+00:00")

	u2, _ := time.Parse(time.RFC3339, "2018-03-18T15:14:37+00:00")
	c2, _ := time.Parse(time.RFC3339, "2018-03-18T15:14:37+00:00")

	u3, _ := time.Parse(time.RFC3339, "2018-03-18T15:14:37+00:00")
	c3, _ := time.Parse(time.RFC3339, "2018-03-18T15:14:37+00:00")

	u4, _ := time.Parse(time.RFC3339, "2018-03-18T15:14:37+00:00")
	c4, _ := time.Parse(time.RFC3339, "2018-03-18T15:14:37+00:00")

	expect := []*Nest{
		{
			ID:          1,
			UUID:        "179a06c9-b5bf-4798-8c0e-9f78e8f1f67f",
			Author:      "support@pterodactyl.io",
			Name:        "Minecraft",
			Description: "Minecraft - the classic game from Mojang. With support for Vanilla MC, Spigot, and many others!",
			CreatedAt:   c1,
			UpdatedAt:   u1,
		},
		{
			ID:          2,
			UUID:        "c389ab01-8b30-4c0f-b2a2-01e79b3611d2",
			Author:      "support@pterodactyl.io",
			Name:        "Source Engine",
			Description: "Includes support for most Source Dedicated Server games.",
			CreatedAt:   c2,
			UpdatedAt:   u2,
		},
		{
			ID:          3,
			UUID:        "8b47c310-ff04-4c10-84c8-133e63f1bf61",
			Author:      "support@pterodactyl.io",
			Name:        "Voice Servers",
			Description: "Voice servers such as Mumble and Teamspeak 3.",
			CreatedAt:   c3,
			UpdatedAt:   u3,
		},
		{
			ID:          4,
			UUID:        "f92cd279-a916-4ede-84ee-900f51048e31",
			Author:      "support@pterodactyl.io",
			Name:        "Rust",
			Description: "Rust - A game where you must fight to survive.",
			CreatedAt:   c4,
			UpdatedAt:   u4,
		},
	}

	got, err := a.GetNests()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if !cmp.Equal(got, expect) {
		t.Errorf("Unexpected response: %s", cmp.Diff(got, expect))
	}
}

func TestApplicationCredentials_GetNest(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/application/nests/1"
		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		res := `{
		  "object": "nest",
		  "attributes": {
			"id": 1,
			"uuid": "179a06c9-b5bf-4798-8c0e-9f78e8f1f67f",
			"author": "support@pterodactyl.io",
			"name": "Minecraft",
			"description": "Minecraft - the classic game from Mojang. With support for Vanilla MC, Spigot, and many others!",
			"created_at": "2018-03-18T15:14:37+00:00",
			"updated_at": "2018-03-18T15:14:37+00:00"
		  }
		}`

		return []byte(res), nil
	}

	a := NewApplication("https://example.com", "")

	u1, _ := time.Parse(time.RFC3339, "2018-03-18T15:14:37+00:00")
	c1, _ := time.Parse(time.RFC3339, "2018-03-18T15:14:37+00:00")

	expect := &Nest{
		ID:          1,
		UUID:        "179a06c9-b5bf-4798-8c0e-9f78e8f1f67f",
		Author:      "support@pterodactyl.io",
		Name:        "Minecraft",
		Description: "Minecraft - the classic game from Mojang. With support for Vanilla MC, Spigot, and many others!",
		CreatedAt:   c1,
		UpdatedAt:   u1,
	}

	got, err := a.GetNest(1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if !cmp.Equal(got, expect) {
		t.Errorf("Unexpected response: %s", cmp.Diff(got, expect))
	}
}

func TestApplicationCredentials_GetEggs(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/application/nests/1/eggs"
		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		res := `{
			  "object": "list",
			  "data": [
				{
				  "object": "egg",
				  "attributes": {
					"id": 1,
					"uuid": "db6323d7-d62f-4278-bb66-db06484b1089",
					"nest": 1,
					"author": "support@pterodactyl.io",
					"description": "Spigot is the most widely-used modded Minecraft server software in the world. It powers many of the top Minecraft server networks around to ensure they can cope with their huge player base and ensure the satisfaction of their players. Spigot works by reducing and eliminating many causes of lag, as well as adding in handy features and settings that help make your job of server administration easier.",
					"docker_image": "quay.io/pterodactyl/core:java-glibc",
					"config": {
					  "files": {
						"server.properties": {
						  "parser": "properties",
						  "find": {
							"server-ip": "0.0.0.0",
							"enable-query": "true",
							"server-port": "{{server.build.default.port}}",
							"query.port": "{{server.build.default.port}}"
						  }
						}
					  },
					  "startup": {
						"done": ")! For help, type ",
						"userInteraction": [
						  "Go to eula.txt for more info."
						]
					  },
					  "stop": "stop",
					  "logs": {
						"custom": false,
						"location": "logs/latest.log"
					  },
					  "extends": null
					},
					"startup": "java -Xms128M -Xmx{{SERVER_MEMORY}}M -jar {{SERVER_JARFILE}}",
					"script": {
					  "privileged": true,
					  "install": "TEST_INSTALL",
					  "entry": "ash",
					  "container": "alpine:3.7",
					  "extends": null
					},
					"created_at": "2018-03-18T15:14:37+00:00",
					"updated_at": "2018-07-08T00:56:48+00:00"
				  }
				},
				{
				  "object": "egg",
				  "attributes": {
					"id": 2,
					"uuid": "2f9bfca1-4a41-4302-83e9-0aa868719eb6",
					"nest": 1,
					"author": "support@pterodactyl.io",
					"description": "Minecraft Forge Server. Minecraft Forge is a modding API (Application Programming Interface), which makes it easier to create mods, and also make sure mods are compatible with each other.",
					"docker_image": "quay.io/pterodactyl/core:java",
					"config": {
					  "files": {
						"server.properties": {
						  "parser": "properties",
						  "find": {
							"server-ip": "0.0.0.0",
							"enable-query": "true",
							"server-port": "{{server.build.default.port}}",
							"query.port": "{{server.build.default.port}}"
						  }
						}
					  },
					  "startup": {
						"done": ")! For help, type ",
						"userInteraction": [
						  "Go to eula.txt for more info."
						]
					  },
					  "stop": "stop",
					  "logs": {
						"custom": false,
						"location": "logs/latest.log"
					  },
					  "extends": null
					},
					"startup": "java -Xms128M -Xmx{{SERVER_MEMORY}}M -jar {{SERVER_JARFILE}}",
					"script": {
					  "privileged": true,
					  "install": "TEST_INSTALL",
					  "entry": "ash",
					  "container": "frolvlad/alpine-oraclejdk8:cleaned",
					  "extends": null
					},
					"created_at": "2018-03-18T15:14:37+00:00",
					"updated_at": "2018-07-08T00:56:48+00:00"
				  }
				}
			  ]
			}`

		return []byte(res), nil
	}

	a := NewApplication("https://example.com", "")

	u1, _ := time.Parse(time.RFC3339, "2018-07-08T00:56:48+00:00")
	c1, _ := time.Parse(time.RFC3339, "2018-03-18T15:14:37+00:00")

	u2, _ := time.Parse(time.RFC3339, "2018-07-08T00:56:48+00:00")
	c2, _ := time.Parse(time.RFC3339, "2018-03-18T15:14:37+00:00")

	expect := []*Egg{
		{
			ID:          1,
			UUID:        "db6323d7-d62f-4278-bb66-db06484b1089",
			Nest:        1,
			Author:      "support@pterodactyl.io",
			Description: "Spigot is the most widely-used modded Minecraft server software in the world. It powers many of the top Minecraft server networks around to ensure they can cope with their huge player base and ensure the satisfaction of their players. Spigot works by reducing and eliminating many causes of lag, as well as adding in handy features and settings that help make your job of server administration easier.",
			DockerImage: "quay.io/pterodactyl/core:java-glibc",
			Config:      EggConfig{
				Startup:      EggStartup{
					Done:            ")! For help, type ",
					UserInteraction: []string{"Go to eula.txt for more info."},
				},
				Stop:         "stop",
				CustomConfig: nil,
				Extends:      "",
			},
			Startup:     "java -Xms128M -Xmx{{SERVER_MEMORY}}M -jar {{SERVER_JARFILE}}",
			Script:      EggScript{
				Privileged: true,
				Install:    "TEST_INSTALL",
				Entry:      "ash",
				Container:  "alpine:3.7",
				Extends:    "",
			},
			CreatedAt:   c1,
			UpdatedAt:   u1,
		},
		{
			ID:          2,
			UUID:        "2f9bfca1-4a41-4302-83e9-0aa868719eb6",
			Nest:        1,
			Author:      "support@pterodactyl.io",
			Description: "Minecraft Forge Server. Minecraft Forge is a modding API (Application Programming Interface), which makes it easier to create mods, and also make sure mods are compatible with each other.",
			DockerImage: "quay.io/pterodactyl/core:java",
			Config:      EggConfig{
				Startup:      EggStartup{
					Done:            ")! For help, type ",
					UserInteraction: []string{"Go to eula.txt for more info."},
				},
				Stop:         "stop",
				CustomConfig: nil,
				Extends:      "",
			},
			Startup:     "java -Xms128M -Xmx{{SERVER_MEMORY}}M -jar {{SERVER_JARFILE}}",
			Script:      EggScript{
				Privileged: true,
				Install:    "TEST_INSTALL",
				Entry:      "ash",
				Container:  "frolvlad/alpine-oraclejdk8:cleaned",
				Extends:    "",
			},
			CreatedAt:   c2,
			UpdatedAt:   u2,
		},
	}

	got, err := a.GetEggs(1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if !cmp.Equal(got, expect) {
		t.Errorf("Unexpected response: %s", cmp.Diff(got, expect))
	}
}

func TestApplicationCredentials_GetEgg(t *testing.T) {
	query = func(url, token, method string, data []byte) ([]byte, error) {
		expectURL := "https://example.com/api/application/nests/1/eggs/1"
		if expectURL != url {
			t.Errorf("Request url does not match expected: %s", url)
		}

		res := `{
		  "object": "egg",
		  "attributes": {
			"id": 1,
			"uuid": "db6323d7-d62f-4278-bb66-db06484b1089",
			"nest": 1,
			"author": "support@pterodactyl.io",
			"description": "Spigot is the most widely-used modded Minecraft server software in the world. It powers many of the top Minecraft server networks around to ensure they can cope with their huge player base and ensure the satisfaction of their players. Spigot works by reducing and eliminating many causes of lag, as well as adding in handy features and settings that help make your job of server administration easier.",
			"docker_image": "quay.io/pterodactyl/core:java-glibc",
			"config": {
			  "files": {
				"server.properties": {
				  "parser": "properties",
				  "find": {
					"server-ip": "0.0.0.0",
					"enable-query": "true",
					"server-port": "{{server.build.default.port}}",
					"query.port": "{{server.build.default.port}}"
				  }
				}
			  },
			  "startup": {
				"done": ")! For help, type ",
				"userInteraction": [
				  "Go to eula.txt for more info."
				]
			  },
			  "stop": "stop",
			  "logs": {
				"custom": false,
				"location": "logs/latest.log"
			  },
			  "extends": null
			},
			"startup": "java -Xms128M -Xmx{{SERVER_MEMORY}}M -jar {{SERVER_JARFILE}}",
			"script": {
			  "privileged": true,
			  "install": "TEST_INSTALL",
			  "entry": "ash",
			  "container": "alpine:3.7",
			  "extends": null
			},
			"created_at": "2018-03-18T15:14:37+00:00",
			"updated_at": "2018-07-08T00:56:48+00:00"
		  }
		}`

		return []byte(res), nil
	}

	a := NewApplication("https://example.com", "")

	u1, _ := time.Parse(time.RFC3339, "2018-07-08T00:56:48+00:00")
	c1, _ := time.Parse(time.RFC3339, "2018-03-18T15:14:37+00:00")

	expect := &Egg{
		ID:          1,
		UUID:        "db6323d7-d62f-4278-bb66-db06484b1089",
		Nest:        1,
		Author:      "support@pterodactyl.io",
		Description: "Spigot is the most widely-used modded Minecraft server software in the world. It powers many of the top Minecraft server networks around to ensure they can cope with their huge player base and ensure the satisfaction of their players. Spigot works by reducing and eliminating many causes of lag, as well as adding in handy features and settings that help make your job of server administration easier.",
		DockerImage: "quay.io/pterodactyl/core:java-glibc",
		Config:      EggConfig{
			Startup:      EggStartup{
				Done:            ")! For help, type ",
				UserInteraction: []string{"Go to eula.txt for more info."},
			},
			Stop:         "stop",
			CustomConfig: nil,
			Extends:      "",
		},
		Startup:     "java -Xms128M -Xmx{{SERVER_MEMORY}}M -jar {{SERVER_JARFILE}}",
		Script:      EggScript{
			Privileged: true,
			Install:    "TEST_INSTALL",
			Entry:      "ash",
			Container:  "alpine:3.7",
			Extends:    "",
		},
		CreatedAt:   c1,
		UpdatedAt:   u1,
	}

	got, err := a.GetEgg(1, 1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if !cmp.Equal(got, expect) {
		t.Errorf("Unexpected response: %s", cmp.Diff(got, expect))
	}
}
