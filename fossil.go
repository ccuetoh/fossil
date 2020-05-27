// Package fossil provides a wrapper for the Pterodactyl and WISP APIs.
package fossil

//***** Credentials *****//

// Credentials is the base object for ClientCredentials and ApplicationCredentials, and should not be used
// independently.
type Credentials struct {
	URL   string
	Token string
}

// ClientCredentials are user-specific, and can only be used to access and modify servers associated
// with that user. They do not allow administrator-level control of the servers.
type ClientCredentials Credentials

// ApplicationCredentials allow you to authenticate as a administrator, and access nodes, servers
// users with creation and destruction privileges.
type ApplicationCredentials Credentials

// NewClient creates a new ClientCredentials object used to interact with Pterodactyl as a user.
func NewClient(url, clientToken string) *ClientCredentials {
	return &ClientCredentials{
		URL:   url,
		Token: clientToken,
	}
}

// NewApplication Creates a new ApplicationCredentials object used to interact with Pterodactyl as administrator.
func NewApplication(url, apiToken string) *ApplicationCredentials {
	return &ApplicationCredentials{
		URL:   url,
		Token: apiToken,
	}
}
