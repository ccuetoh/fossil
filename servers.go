package fossil

import (
	"encoding/json"
	"fmt"
	"time"
)

//***** Structures *****//

// Power States
const (
	ON      = "start"
	OFF     = "stop"
	RESTART = "restart"
	KILL    = "kill"
)

// Servers

// ClientServer defines Pterodactyl server as a user would see it. It is fetched and interacted with the client token.
type ClientServer struct {
	ID                string
	Name              string
	Description       string
	Limits            Limits
	AllocationDetails []Allocation
	IsOwner           bool
}

// ApplicationServer defines Pterodactyl server as an administrator would see it. It is fetched and interacted with
// the api token.
type ApplicationServer struct {
	ID                 int
	Name               string
	Description        string
	Limits             Limits
	ExternalID         string
	UUID               string
	Suspended          bool
	User               int
	Node               int
	Nest               int
	Egg                int
	Pack               int
	Allocation         int
	AllocationsDetails []Allocation
	Container          Container
	Updated            time.Time
	Created            time.Time
}

// jsonServer is the API definition for the server, and contains all the data in it's original form.
// It's used as the target struct in the marshalling/unmarshalling of API requests or responses.
type jsonServer struct {
	ID            int    `json:"id"`
	ExternalID    string `json:"external_id"`
	UUID          string `json:"uuid"`
	Identifier    string `json:"identifier"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Suspended     bool   `json:"suspended"`
	ServerOwner   bool   `json:"server_owner"`
	Limits        Limits `json:"limits"`
	FeatureLimits struct {
		Databases   int `json:"databases"`
		Allocations int `json:"allocations"`
	} `json:"feature_limits"`
	User          int       `json:"user"`
	Node          int       `json:"node"`
	Allocation    int       `json:"allocation"`
	Nest          int       `json:"nest"`
	Egg           int       `json:"egg"`
	Pack          int       `json:"pack"`
	Container     Container `json:"container"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
	Relationships struct {
		Allocations struct {
			Data []struct {
				Allocation *Allocation `json:"attributes"`
			} `json:"data"`
		} `json:"allocations"`
	} `json:"relationships"`
}

// jsonServerCreation stores the server info in an API-ready format for server creation
type jsonServerCreation struct {
	ExternalID    string            `json:"external_id"`
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	Limits        Limits            `json:"limits"`
	DockerImage   string            `json:"docker_image,omitempty"`
	Startup       string            `json:"startup"`
	Environment   map[string]string `json:"environment"`
	SkipScripts   bool              `json:"skip_scripts"`
	FeatureLimits struct {
		Databases   int `json:"databases"`
		Allocations int `json:"allocations"`
	} `json:"feature_limits"`
	User       int `json:"user"`
	Node       int `json:"node"`
	Allocation struct {
		Default    int   `json:"default,omitempty"`
		Additional []int `json:"additional,omitempty"`
	} `json:"allocation"`
	Nest int `json:"nest"`
	Egg  int `json:"egg"`
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
	Memory      int `json:"memory"`
	Swap        int `json:"swap"`
	Disk        int `json:"disk"`
	IO          int `json:"io"`
	CPU         int `json:"cpu"`
	Databases   int `json:"-"`
	Allocations int `json:"-"`
}

// Container holds all the Docker Image settings
type Container struct {
	StartupCommand string            `json:"startup_command"`
	Image          string            `json:"image"`
	Installed      bool              `json:"installed"`
	Environment    map[string]string `json:"environment"`
}

// Allocation holds all the information relating to the allocation data of a server
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

// Disk holds the usage and limits of the server disk
type Disk struct {
	Used  uint64
	Limit uint64
}

// Players holds the number of users in a server and the max amount permitted. Some Pterodactyl API-based
// software DO NOT provide this information.
type Players struct {
	Current uint64
	Limit   uint64
}

//***** Converters *****//

// asClientServer parses a jsonServer into a *ClientServer
func (s *jsonServer) asClientServer() *ClientServer {
	cs := &ClientServer{
		ID:          s.Identifier,
		Name:        s.Name,
		Description: s.Description,
		Limits:      s.Limits,
		IsOwner:     s.ServerOwner,
	}
	cs.Limits.Databases = s.FeatureLimits.Databases

	for _, alloc := range s.Relationships.Allocations.Data {
		cs.AllocationDetails = append(cs.AllocationDetails, *alloc.Allocation)
	}

	return cs
}

// asClientServers parses a jsonServerPage into an slice of *ClientServers
func (sp *jsonServerPage) asClientServers() (servers []*ClientServer) {
	for _, d := range sp.Data {
		servers = append(servers, d.Server.asClientServer())
	}

	return servers
}

// asApplicationServer parses a jsonServer into an *ApplicationServer
func (s *jsonServer) asApplicationServer() *ApplicationServer {
	as := &ApplicationServer{
		ID:          s.ID,
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
		Pack:        s.Pack,
		Container:   s.Container,
		Updated:     s.UpdatedAt,
		Created:     s.CreatedAt,
		Allocation:  s.Allocation,
	}

	as.Limits.Databases = s.FeatureLimits.Databases

	for _, alloc := range s.Relationships.Allocations.Data {
		as.AllocationsDetails = append(as.AllocationsDetails, *alloc.Allocation)
	}

	return as
}

// asApplicationServers parses a jsonServerPage into an slice of *ApplicationServer
func (sp *jsonServerPage) asApplicationServers() (servers []*ApplicationServer) {
	for _, d := range sp.Data {
		servers = append(servers, d.Server.asApplicationServer())
	}

	return servers
}

// asJsonServerCreation parses a ApplicationServer into a JSON-ready *jsonServerCreation
func (s *ApplicationServer) asJsonServerCreation() *jsonServerCreation {
	js := &jsonServerCreation{
		ExternalID:  s.ExternalID,
		Name:        s.Name,
		Description: s.Description,
		Limits: Limits{
			Memory: s.Limits.Memory,
			Swap:   s.Limits.Swap,
			Disk:   s.Limits.Disk,
			IO:     s.Limits.IO,
			CPU:    s.Limits.CPU,
		},
		User:        s.User,
		Node:        s.Node,
		Nest:        s.Nest,
		Egg:         s.Egg,
		DockerImage: s.Container.Image,
		Startup:     s.Container.StartupCommand,
		Environment: s.Container.Environment,
	}

	js.FeatureLimits.Databases = s.Limits.Databases
	js.Allocation.Default = s.Allocation

	for _, alloc := range s.AllocationsDetails {
		if s.Allocation == alloc.Port {
			continue
		}
		js.Allocation.Additional = append(js.Allocation.Additional, alloc.Port)
	}

	return js
}

//***** String *****//

func (s *ClientServer) String() string {
	return s.Name
}
func (s *ApplicationServer) String() string {
	return s.Name
}

//***** Pagination *****//

// getAll fetches all the existing pages for a server list. The original page is kept as index 0
func (sp *jsonServerPage) getAll(token string) (pages []*jsonServerPage, err error) {
	pages = append(pages, sp)
	for pages[len(pages)-1].Meta.Pagination.Links.Next != "" {
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

//***** Requests *****//

// GetServer fetches the server with the given Internal ID if it exists
func (c *ApplicationCredentials) GetServer(internalId int) (sv *ApplicationServer, err error) {
	bytes, err := c.query(fmt.Sprintf("servers/%d?include=allocations", internalId), "GET", nil)
	if err != nil {
		return
	}

	var wrapper struct {
		Server jsonServer `json:"attributes"`
	}

	err = json.Unmarshal(bytes, &wrapper)
	if err != nil {
		return
	}

	return wrapper.Server.asApplicationServer(), nil
}

// GetServerExternal fetches the server with the given External ID if it exists
func (c *ApplicationCredentials) GetServerExternal(externalId string) (sv *ApplicationServer, err error) {
	bytes, err := c.query("servers/external/"+externalId+"?include=allocations", "GET", nil)
	if err != nil {
		return
	}

	var wrapper struct {
		Server jsonServer `json:"attributes"`
	}

	err = json.Unmarshal(bytes, &wrapper)
	if err != nil {
		return
	}

	return wrapper.Server.asApplicationServer(), nil
}

// GetServers fetches the servers of all the users
func (c *ApplicationCredentials) GetServers() (svs []*ApplicationServer, err error) {
	bytes, err := c.query("servers?include=allocations", "GET", nil)
	if err != nil {
		return
	}

	// We get the initial page
	var page jsonServerPage
	err = json.Unmarshal(bytes, &page)
	if err != nil {
		return
	}

	// We'll search for the remaining pages if present
	pages, err := page.getAll(c.Token)
	if err != nil {
		return
	}

	for _, page := range pages {
		svs = append(svs, page.asApplicationServers()...)
	}

	return
}

// CreateServer creates a new server
func (c *ApplicationCredentials) CreateServer(sv *ApplicationServer) (err error) {
	bytes, err := json.Marshal(sv.asJsonServerCreation())
	if err != nil {
		return err
	}

	_, err = c.query("servers", "POST", bytes)
	if err != nil {
		return
	}

	return nil
}

// UpdateDetails modifies the server name, user, external id and description
func (c *ApplicationCredentials) UpdateDetails(sv *ApplicationServer) (err error) {
	type details struct {
		ExternalID  string `json:"external_id"`
		Name        string `json:"name"`
		User        int    `json:"user"`
		Description string `json:"description,omitempty"`
	}

	d := details{
		ExternalID:  sv.ExternalID,
		Name:        sv.Name,
		User:        sv.User,
		Description: sv.Description,
	}

	bytes, err := json.Marshal(d)
	if err != nil {
		return err
	}

	_, err = c.query(fmt.Sprintf("servers/%d/details", sv.ID), "PATCH", bytes)
	if err != nil {
		return
	}

	return nil
}

// UpdateBuild modifies the server's limit and allocation configuration
func (c *ApplicationCredentials) UpdateBuild(sv *ApplicationServer, addAlloc []int, removeAlloc []int) (err error) {
	type build struct {
		Allocation        int     `json:"allocation,omitempty"`
		OOM               bool    `json:"oom_disabled"`
		Limits            *Limits `json:"limits,omitempty"`
		AddAllocations    []int   `json:"add_allocations,omitempty"`
		RemoveAllocations []int   `json:"remove_allocations,omitempty"`
		FeatureLimits     struct {
			Databases   int `json:"databases"`
			Allocations int `json:"allocations,omitempty"`
		} `json:"feature_limits"`
	}

	b := build{
		OOM:               true,
		Allocation:        sv.Allocation,
		Limits:            &sv.Limits,
		AddAllocations:    addAlloc,
		RemoveAllocations: removeAlloc,
	}

	b.FeatureLimits.Allocations = sv.Limits.Allocations
	b.FeatureLimits.Databases = sv.Limits.Databases

	bytes, err := json.Marshal(b)
	if err != nil {
		return err
	}

	_, err = c.query(fmt.Sprintf("servers/%d/build", sv.ID), "PATCH", bytes)
	if err != nil {
		return
	}

	return nil
}

// UpdateStartup modifies the server's startup parameters, egg and image configuration
func (c *ApplicationCredentials) UpdateStartup(sv *ApplicationServer) (err error) {
	type startup struct {
		Startup     string   `json:"startup"`
		Environment []string `json:"environment"`
		Egg         int      `json:"egg"`
		Pack        int      `json:"pack"`
		Image       string   `json:"image"`
		SkipScripts bool     `json:"skip_scripts"`
	}

	// Since Pterodactyl doesn't like to have a stable definition for it's parameters we need Environment
	// as a slice of the env of the environmental variables and not as a map.
	//
	// A note on this approach: Map operations are NOT thread safe
	env := make([]string, 0, len(sv.Container.Environment))
	for k := range sv.Container.Environment {
		env = append(env, k)
	}

	su := startup{
		Startup:     sv.Container.StartupCommand,
		Environment: env,
		Egg:         sv.Egg,
		Image:       sv.Container.Image,
		Pack:        sv.Pack,
	}

	bytes, err := json.Marshal(su)
	if err != nil {
		return err
	}

	_, err = c.query(fmt.Sprintf("servers/%d/startup", sv.ID), "PATCH", bytes)
	return
}

// SuspendServer marks a server as suspended
func (c *ApplicationCredentials) SuspendServer(sid int) (err error) {
	_, err = c.query(fmt.Sprintf("servers/%d/suspend", sid), "POST", nil)
	return
}

// UnsuspendServer marks a server as active
func (c *ApplicationCredentials) UnsuspendServer(sid int) (err error) {
	_, err = c.query(fmt.Sprintf("servers/%d/unsuspend", sid), "POST", nil)
	return
}

// RebuildServer starts a server rebuild
func (c *ApplicationCredentials) RebuildServer(sid int) (err error) {
	_, err = c.query(fmt.Sprintf("servers/%d/rebuild", sid), "POST", nil)
	return
}

// ReinstallServer marks a server for reinstallation
func (c *ApplicationCredentials) ReinstallServer(sid int) (err error) {
	_, err = c.query(fmt.Sprintf("servers/%d/reinstall", sid), "POST", nil)
	return
}

// DeleteServer marks a server for deletion
func (c *ApplicationCredentials) DeleteServer(sid int) (err error) {
	_, err = c.query(fmt.Sprintf("servers/%d", sid), "DELETE", nil)
	return
}

// ForceDeleteServer forcefully deletes a server. This is an ungraceful way to delete the server, and when
// possible DeleteServer should be preferred.
func (c *ApplicationCredentials) ForceDeleteServer(sid int) (err error) {
	_, err = c.query(fmt.Sprintf("servers/%d/force", sid), "DELETE", nil)
	return
}
