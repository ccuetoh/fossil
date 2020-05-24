package fossil

import (
	"encoding/json"
	"time"
)

//***** Structures *****//

// Power States
const (
	ON = "start"
	OFF = "stop"
	RESTART = "restart"
	KILL = "kill"
)

// Servers

// ClientServer defines Pterodactyl server as a user would see it. It is fetched and interacted with the client token.
type ClientServer struct {
	ID          string
	Name        string
	Description string
	Limits      *Limits
	Allocation  Allocation
	IsOwner     bool
}

// AdminServer defines Pterodactyl server as an administrator would see it. It is fetched and interacted with
// the api token.
type AdminServer struct {
	ID          string
	Name        string
	Description string
	Limits      *Limits
	ExternalID  string
	UUID        string
	Suspended   bool
	User        int
	Node        int
	Nest        int
	Egg         int
	Allocation  Allocation
	Container   *Container
}

// jsonServer is the API definition for the server, and contains all the data in it's original form.
// It's used as the target struct in the marshalling/unmarshalling of API requests or responses.
type jsonServer struct {
	ID            int     `json:"id"`
	ExternalID    string  `json:"external_id"`
	UUID          string  `json:"uuid"`
	Identifier    string  `json:"identifier"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Suspended     bool    `json:"suspended"`
	ServerOwner   bool    `json:"server_owner"`
	Limits        *Limits `json:"limits"`
	FeatureLimits struct {
		Databases   int `json:"databases"`
		Allocations int `json:"allocations"`
	} `json:"feature_limits"`
	User          int        `json:"user"`
	Node          int        `json:"node"`
	Allocation    int        `json:"allocation"`
	Nest          int        `json:"nest"`
	Egg           int        `json:"egg"`
	Container     *Container `json:"container"`
	UpdatedAt     time.Time  `json:"updated_at"`
	CreatedAt     time.Time  `json:"created_at"`
	Relationships struct {
		Allocations struct {
			Data   []struct {
				Allocation *Allocation `json:"attributes"`
			} `json:"data"`
		} `json:"allocations"`
	} `json:"relationships"`
}

// jsonServerPage contains a page of jsonServers and the pagination data.
// It's used as the target struct in the marshalling/unmarshalling of API requests or responses.
type jsonServerPage struct {
	Data []struct {
		Server *jsonServer `json:"attributes"`
	} `json:"data"`
	Meta struct {
		Pagination struct {
			Total       int `json:"total"`
			Count       int `json:"count"`
			PerPage     int `json:"per_page"`
			CurrentPage int `json:"current_page"`
			TotalPages  int `json:"total_pages"`
			Links       struct {
				Previous string `json:"previous,omitempty"`
				Next     string `json:"next,omitempty"`
			} `json:"links"`
		} `json:"pagination"`
	} `json:"meta"`
}

// Limits contains all the allocated usage limits set for a server
type Limits struct {
	Memory    int `json:"memory"`
	Swap      int `json:"swap"`
	Disk      int `json:"disk"`
	IO        int `json:"io"`
	CPU       int `json:"cpu"`
	Databases int
}

// Container holds all the Docker Image settings
type Container struct {
	StartupCommand string            `json:"startup_command"`
	Image          string            `json:"image"`
	Installed      bool              `json:"installed"`
	Environment    map[string]string `json:"environment"`
}

type Allocation struct {
	Primary bool   `json:"primary"`
	IP      string `json:"ip"`
	Alias   string `json:"alias"`
	Port    int    `json:"port"`
}

// Status

// ServerStatus contains the client-visible server usage information
type ServerStatus struct {
	State   string
	Memory  Memory
	CPU     CPU
	Disk    Disk
	Players Players
}

// Memory holds the usage and limits of the server memory
type Memory struct {
	Used  uint64
	Limit uint64
}

// CPU holds the usage, core status and limits of the server CPU
type CPU struct {
	Current float32
	Cores   []float32
	Limit   uint64
}

// Memory holds the usage and limits of the server disk
type Disk struct {
	Current uint64
	Limit   uint64
}

// Players holds the number of users in a server and the max amount permitted. Some Pterodactyl API-based
// software DO NOT provide this information.
type Players struct {
	Current uint64
	Limit   uint64
}

//***** Converters *****//

// asClientServer parses a jsonServer into a *ClientServer
func (s *jsonServer) asClientServer() *ClientServer{
	cs := &ClientServer{
		ID:          s.Identifier,
		Name:        s.Name,
		Description: s.Description,
		Limits:      s.Limits,
		IsOwner:     s.ServerOwner,
	}
	cs.Limits.Databases = s.FeatureLimits.Databases

	alloc := s.Relationships.Allocations.Data
	if len(alloc) > 0{
		cs.Allocation = *alloc[0].Allocation
	}

	return cs
}

// asClientServers parses a jsonServerPage into an slice of *ClientServers
func (sp *jsonServerPage) asClientServers() (servers []*ClientServer){
	for _, d := range sp.Data{
		servers = append(servers, d.Server.asClientServer())
	}

	return servers
}

// asAdminServer parses a jsonServer into an *AdminServer
func (s *jsonServer) asAdminServer() *AdminServer{
	as := &AdminServer{
		ID:          s.Identifier,
		Name:        s.Name,
		Description: s.Description,
		Limits:      s.Limits,
		ExternalID:  s.ExternalID,
		UUID:        s.UUID,
		Suspended:   s.Suspended,
		User:        s.User,
		Node:        s.Node,
		Nest:        s.Nest,
		Egg:         s.Egg,
		Container:   s.Container,
	}

	as.Limits.Databases = s.FeatureLimits.Databases

	alloc := s.Relationships.Allocations.Data
	if len(alloc) > 0{
		as.Allocation = *alloc[0].Allocation
	}

	return as
}

// asAdminServers parses a jsonServerPage into an slice of *AdminServer
func (sp *jsonServerPage) asAdminServers() (servers []*AdminServer, ){
	for _, d := range sp.Data{
		servers = append(servers, d.Server.asAdminServer())
	}

	return servers
}

//***** String *****//

func (s *ClientServer) String() string {
	return s.Name
}
func (s *AdminServer) String() string {
	return s.Name
}

//***** Pagination *****//

// getAll fetches all the existing pages for a server list. The original page is kept as index 0
func (sp *jsonServerPage) getAll(token string) (pages []*jsonServerPage, err error) {
	pages = append(pages, sp)
	for pages[len(pages)-1].Meta.Pagination.Links.Next != ""{
		url := sp.Meta.Pagination.Links.Next + "&include=allocations"
		bytes, err := queryCallback(url, token, "GET", nil)
		if err != nil {
			return nil, err
		}

		var page jsonServerPage
		err = json.Unmarshal(bytes, &page)
		if err != nil {
			return nil, err
		}

		pages = append(pages, &page)
	}

	return pages, nil
}
