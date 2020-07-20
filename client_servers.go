package fossil

import (
	"encoding/json"
)

//***** Structures *****//

// Servers

// ClientServer defines Pterodactyl server as a user would see it from the server list.\
// It is fetched and interacted with the client token.
type ClientServer struct {
	ID           string
	UUID         string
	Name         string
	Description  string
	Node         string
	Limits       Limits
	SFTPDetails  SFTPDetails
	Allocation   Allocation
	IsOwner      bool
	IsInstalling bool
}

// ClientServerDetail defines Pterodactyl server as a user would see it from the server details view.
// It is fetched and interacted with the client token.
type ClientServerDetail struct {
	ID           string
	UUID         string
	Name         string
	Description  string
	Node         string
	Limits       Limits
	SFTPDetails  SFTPDetails
	Allocations  []*AllocationDetails
	IsOwner      bool
	IsInstalling bool
	IsSuspended  bool
	Permissions  []string
}

// AllocationDetails extended allocation info
type AllocationDetails struct {
	ID        int    `json:"id"`
	IP        string `json:"ip"`
	IPAlias   string `json:"ip_alias"`
	Port      int    `json:"port"`
	Notes     string `json:"notes"`
	IsDefault bool   `json:"is_default"`
}

//***** Converters *****//

// asClientServer parses a jsonServer into a *ClientServer
func (s *jsonServer) asClientServer() *ClientServer {
	cs := &ClientServer{
		ID:           s.Identifier,
		Name:         s.Name,
		Description:  s.Description,
		Limits:       s.Limits,
		IsOwner:      s.ServerOwner,
		UUID:         s.UUID,
		Node:         s.Node,
		IsInstalling: s.IsInstalling,
		SFTPDetails:  s.SFTPDetails,
		Allocation:   s.Allocation,
	}
	cs.Limits.Databases = s.FeatureLimits.Databases
	cs.Limits.Allocations = s.FeatureLimits.Allocations
	cs.Limits.Backups = s.FeatureLimits.Backups

	return cs
}

// asClientServers parses a jsonServerPage into an slice of *ClientServers
func (sp *jsonServerPage) asClientServers() (servers []*ClientServer) {
	for _, d := range sp.Data {
		servers = append(servers, d.Server.asClientServer())
	}

	return servers
}

// asClientServers parses a jsonServer into a *ClientServerDetails
func (s *jsonServer) asClientServerDetail() *ClientServerDetail {
	csd := &ClientServerDetail{
		ID:           s.Identifier,
		UUID:         s.UUID,
		Name:         s.Name,
		Description:  s.Description,
		Node:         s.Node,
		Limits:       s.Limits,
		SFTPDetails:  s.SFTPDetails,
		IsInstalling: s.IsInstalling,
		IsSuspended:  s.IsSuspended,
	}

	csd.Limits.Databases = s.FeatureLimits.Databases
	csd.Limits.Allocations = s.FeatureLimits.Allocations
	csd.Limits.Backups = s.FeatureLimits.Backups

	for _, data := range s.Relationships.Allocations.Data {
		csd.Allocations = append(csd.Allocations, data.Allocation)
	}

	return csd
}

//***** Requests *****//

// Servers

// GetServers fetches all the servers of the client
func (c *ClientCredentials) GetServers() (svs []*ClientServer, err error) {
	bytes, err := c.query("?include=allocations", "GET", nil)
	if err != nil {
		return
	}

	// Get the initial page
	var page jsonServerPage
	err = json.Unmarshal(bytes, &page)
	if err != nil {
		return
	}

	// Search for the remaining pages if present
	pages, err := page.getAll(c.Token)
	if err != nil {
		return
	}

	for _, page := range pages {
		svs = append(svs, page.asClientServers()...)
	}

	return
}

// GetServer fetches the server with the given ID if it exists
func (c *ClientCredentials) GetServer(id string) (sv *ClientServerDetail, err error) {
	bytes, err := c.query("servers/"+id, "GET", nil)
	if err != nil {
		return
	}

	var wrapper struct {
		Server jsonServer `json:"attributes"`
		Meta   struct {
			IsOwner         bool     `json:"is_server_owner"`
			UserPermissions []string `json:"user_permissions"`
		} `json:"meta"`
	}

	err = json.Unmarshal(bytes, &wrapper)
	if err != nil {
		return
	}

	sv = wrapper.Server.asClientServerDetail()
	sv.Permissions = wrapper.Meta.UserPermissions
	sv.IsOwner = wrapper.Meta.IsOwner

	return sv, nil
}

// GetServerStatus fetches the server's status and usage
func (c *ClientCredentials) GetServerStatus(id string) (ss *ServerStatus, err error) {
	bytes, err := c.query("servers/"+id+"/utilization", "GET", nil)
	if err != nil {
		return
	}

	var wrapper struct {
		Utilization *ServerStatus `json:"attributes"`
	}

	err = json.Unmarshal(bytes, &wrapper)
	if err != nil {
		return
	}

	return wrapper.Utilization, nil
}

// ExecuteCommand allows the execution of a console command on the specified server
func (c *ClientCredentials) ExecuteCommand(id string, cmd string) (err error) {
	type wrapper struct {
		Command string `json:"command"`
	}

	cmdWrapper := wrapper{Command: cmd}

	rq, err := json.Marshal(cmdWrapper)
	if err != nil {
		return
	}

	_, err = c.query("servers/"+id+"/command", "POST", rq)
	if err != nil {
		return
	}

	return nil
}

// SetPowerState changes the power state of a server. Will result in error if the server is already in that state
// or is unable to change state.
func (c *ClientCredentials) SetPowerState(id string, state string) (err error) {
	type wrapper struct {
		Signal string `json:"signal"`
	}

	signalWrapper := wrapper{Signal: state}

	rq, err := json.Marshal(signalWrapper)
	if err != nil {
		return
	}

	_, err = c.query("servers/"+id+"/power", "POST", rq)
	if err != nil {
		return
	}

	return nil
}


