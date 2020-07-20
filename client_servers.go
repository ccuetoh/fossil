package fossil

import (
	"encoding/json"
)

//***** Requests *****//

// GetServer fetches the server with the given ID if it exists
func (c *ClientCredentials) GetServer(id string) (sv *ClientServer, err error) {
	bytes, err := c.query("servers/"+id+"?include=allocations", "GET", nil)
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

	return wrapper.Server.asClientServer(), nil
}

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


