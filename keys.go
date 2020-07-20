package fossil

import (
	"encoding/json"
	"time"
)

//***** Requests *****//

// APIKey represents an API Key emitted to the user
type APIKey struct {
	Identifier  string    `json:"identifier"`
	Description string    `json:"description"`
	AllowedIPs  []string  `json:"allowed_ips"`
	LastUsed    time.Time `json:"last_used_at"`
	Created     time.Time `json:"created_at"`
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