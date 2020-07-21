package fossil

import (
	"encoding/json"
	"fmt"
)

//***** Structures *****//

// Database contains all the information of a specific Pterodactyl database
type Database struct {
	ID              string `json:"id"`
	Host            *Host  `json:"host"`
	Name            string `json:"name"`
	Username        string `json:"username"`
	ConnectionsFrom string `json:"connections_from"`
	MaxConnections  int    `json:"max_connections"`
}

// Host holds the database's host information
type Host struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

//***** Requests *****//

// GetDatabases fetches all the associated databases for a server
func (c *ApplicationCredentials) GetDatabases(sid int) (dbs []*Database, err error) {
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

// GetDatabase fetches, if present, the database matching the id in the server's databases
func (c *ApplicationCredentials) GetDatabase(sid int, dbid int) (db *Database, err error) {
	bytes, err := c.query(fmt.Sprintf("servers/%d/databases/%d", sid, dbid), "GET", nil)
	if err != nil {
		return
	}

	var wrapper struct {
		Database *Database `json:"attributes"`
	}

	err = json.Unmarshal(bytes, &wrapper)
	if err != nil {
		return
	}

	return wrapper.Database, nil
}

// CreateDatabase creates a new database based on the provided information
func (c *ApplicationCredentials) CreateDatabase(sid int, db *Database) (err error) {
	type databaseCreate struct {
		Database string `json:"database"`
		Remote   string `json:"remote"`
		Host     int    `json:"host"`
	}

	dbStruct := databaseCreate{
		// #TODO
	}

	bytes, err := json.Marshal(dbStruct)
	if err != nil {
		return err
	}
	_, err = c.query(fmt.Sprintf("servers/%d/databases", sid), "POST", bytes)
	return
}

// ResetDatabasePassword resets the password for the specified database of the specified server
func (c *ApplicationCredentials) ResetDatabasePassword(sid int, dbid int) (err error) {
	_, err = c.query(fmt.Sprintf("servers/%d/databases/%d/reset-password", sid, dbid), "POST", nil)
	return
}

// DeleteDatabase marks the specified database in the specified server for deletion
func (c *ApplicationCredentials) DeleteDatabase(sid int, dbid int) (err error) {
	_, err = c.query(fmt.Sprintf("servers/%d/databases/%d", sid, dbid), "DELETE", nil)
	return
}
