package fossil

import (
	"encoding/json"
)

//***** Requests *****//

// User info

// WhoAmI fetches information about the current account
func (c *ClientCredentials) WhoAmI() (me *Me, err error) {
	bytes, err := c.query("account", "GET", nil)
	if err != nil {
		return
	}

	var wrapper struct {
		Me Me `json:"attributes"`
	}

	err = json.Unmarshal(bytes, &wrapper)
	if err != nil {
		return
	}

	return &wrapper.Me, nil
}

// Two-factor authentication

// Get2FAImageURL gets the two-factor authentication QR code for the setup process
func (c *ClientCredentials) Get2FAImageURL() (imageData string, err error) {
	bytes, err := c.query("account/two-factor", "GET", nil)
	if err != nil {
		return
	}

	var wrapper struct {
		Data struct{
			URL string `json:"imageurldata"`
		} `json:"data"`
	}

	err = json.Unmarshal(bytes, &wrapper)
	if err != nil {
		return
	}

	return wrapper.Data.URL, nil
}

// Enable2FA setups two-factor authentication in the account
func (c *ClientCredentials) Enable2FA(code string) (err error) {
	codeStruct := struct {
		Code string `json:"code"`
	}{Code: code}

	bytes, err := json.Marshal(codeStruct)
	if err != nil {
		return err
	}

	_, err = c.query("account/two-factor", "POST", bytes)
	return
}

// Disable2FA removes two-factor authentication from the account
func (c *ClientCredentials) Disable2FA(password string) (err error) {
	passStruct := struct {
		Password string `json:"password"`
	}{Password: password}

	bytes, err := json.Marshal(passStruct)
	if err != nil {
		return err
	}

	_, err = c.query("account/two-factor", "DELETE", bytes)
	return
}

// UpdateEmail modifies the email address of the account
func (c *ClientCredentials) UpdateEmail(email string, password string) (err error) {
	emailStruct := struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}{Password: password, Email: email}

	bytes, err := json.Marshal(emailStruct)
	if err != nil {
		return err
	}

	_, err = c.query("account/email", "PUT", bytes)
	return
}

// UpdatePassword modifies the password of the account
func (c *ClientCredentials) UpdatePassword(oldPassword string, newPassword string) (err error) {
	passStruct := struct {
		OldPassword        string `json:"current_password"`
		NewPassword        string `json:"password"`
		NewPasswordConfirm string `json:"password_confirmation"`
	}{OldPassword: oldPassword, NewPassword: newPassword, NewPasswordConfirm: newPassword}

	bytes, err := json.Marshal(passStruct)
	if err != nil {
		return err
	}

	_, err = c.query("account/password", "PUT", bytes)
	return
}


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
