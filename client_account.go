package fossil

import (
	"encoding/json"
	"time"
)

//***** Structs *****//

// APIKey represents an API Key emitted to the user
type APIKey struct {
	Identifier  string    `json:"identifier"`
	Description string    `json:"description"`
	AllowedIPs  []string  `json:"allowed_ips"`
	LastUsed    time.Time `json:"last_used_at"`
	Created     time.Time `json:"created_at"`
}

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
			URL string `json:"image_url_data"`
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

// Account settings

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

// GetAPIKeys returns all the API keys emitted to the user
func (c *ClientCredentials) GetAPIKeys() (keys []*APIKey, err error) {
	bytes, err := c.query("account/api-keys", "GET", nil)
	if err != nil {
		return
	}

	var wrapper struct {
		Data []struct{
			APIKey APIKey `json:"attributes"`
		} `json:"data"`
	}

	err = json.Unmarshal(bytes, &wrapper)
	if err != nil {
		return
	}

	for _, keyholder := range wrapper.Data{
		keys = append(keys, &keyholder.APIKey)
	}

	return
}

// NewAPIKey creates a new token and returns the key data and secret
func (c *ClientCredentials) NewAPIKey(description string) (key *APIKey, secret string, err error) {
	keyStruct := struct {
		Description string `json:"description"`
	}{Description: description}

	rqBytes, err := json.Marshal(keyStruct)
	if err != nil {
		return nil, "", err
	}

	resBytes, err := c.query("account/api-keys", "POST", rqBytes)
	if err != nil {
		return
	}

	var wrapper struct {
		APIKey APIKey `json:"attributes"`
		Meta struct {
			SecretToken string `json:"secret_token"`
		} `json:"meta"`
	}

	err = json.Unmarshal(resBytes, &wrapper)
	if err != nil {
		return
	}

	return &wrapper.APIKey, wrapper.Meta.SecretToken, nil
}


// DeleteAPIKey deletes an existing API Key based on it's ID
func (c *ClientCredentials) DeleteAPIKey(id string) (err error) {
	_, err = c.query("account/api-keys/" + id, "DELETE", nil)
	return
}
