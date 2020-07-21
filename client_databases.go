package fossil

import (
	"encoding/json"
	"fmt"
)

//***** Structures *****//

// GetDatabases fetches all the associated databases for a server
func (c *ClientCredentials) GetDatabases(sid int) (dbs []*Database, err error) {
	bytes, err := c.query(fmt.Sprintf("servers/%d/databases", sid), "GET", nil)
	if err != nil {
		return
	}

	var wrapper struct {
		Data []struct {
			Database *Database `json:"attributes"`
		}
	}

	err = json.Unmarshal(bytes, &wrapper)
	if err != nil {
		return
	}

	for _, db := range wrapper.Data {
		dbs = append(dbs, db.Database)
	}

	return dbs, nil
}

// CreateDatabase creates a new database based on the provided information
func (c *ClientCredentials) CreateDatabase(sid string) (db *Database, password string, err error) {
	bytes, err := c.query(fmt.Sprintf("servers/%s/databases", sid), "POST", nil)

	var dbResponse struct {
		Attributes struct {
			ID   string `json:"id"`
			Host *Host `json:"host"`
			Name            string `json:"name"`
			Username        string `json:"username"`
			ConnectionsFrom string `json:"connections_from"`
			MaxConnections  int    `json:"max_connections"`
			Relationships   struct {
				Password struct {
					Object     string `json:"object"`
					Attributes struct {
						Password string `json:"password"`
					} `json:"attributes"`
				} `json:"password"`
			} `json:"relationships"`
		} `json:"attributes"`
	}

	err = json.Unmarshal(bytes, &dbResponse)
	if err != nil {
		return
	}

	db = &Database{
		ID:              dbResponse.Attributes.ID,
		Host:            dbResponse.Attributes.Host,
		Name:            dbResponse.Attributes.Name,
		Username:        dbResponse.Attributes.Username,
		ConnectionsFrom: dbResponse.Attributes.ConnectionsFrom,
		MaxConnections:  dbResponse.Attributes.MaxConnections,
	}

	password = dbResponse.Attributes.Relationships.Password.Attributes.Password

	return db, password, nil
}