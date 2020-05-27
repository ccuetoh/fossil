package fossil

import (
	"encoding/json"
	"fmt"
	"time"
)

// User holds user information
type User struct {
	ID                      int       `json:"id"`
	ExternalID              string    `json:"external_id"`
	UUID                    string    `json:"uuid"`
	Username                string    `json:"username"`
	Email                   string    `json:"email"`
	FirstName               string    `json:"first_name"`
	LastName                string    `json:"last_name"`
	Language                string    `json:"language"`
	RootAdmin               bool      `json:"root_admin"`
	TwoFactorAuthentication bool      `json:"2fa"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}

// jsonUserPage contains a page of Users and the pagination data.
// It's used as the target struct in the marshalling/unmarshalling of API requests or responses.
type jsonUserPage struct {
	Data []struct {
		User *User `json:"attributes"`
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

//***** Pagination *****//

// getAll fetches all the existing pages for a user list. The original page is kept as index 0
func (up *jsonUserPage) getAll(token string) (pages []*jsonUserPage, err error) {
	pages = append(pages, up)
	for pages[len(pages)-1].Meta.Pagination.Links.Next != "" {
		url := up.Meta.Pagination.Links.Next
		bytes, err := queryURL(url, token, "GET", nil)
		if err != nil {
			return nil, err
		}

		var page jsonUserPage
		err = json.Unmarshal(bytes, &page)
		if err != nil {
			return nil, err
		}

		pages = append(pages, &page)
	}

	return pages, nil
}

//***** Converters *****//

// asUserSlice parses a jsonUserPage into a []*User
func (up *jsonUserPage) asUserSlice() (users []*User) {
	for _, u := range up.Data {
		users = append(users, u.User)
	}

	return users
}

//***** Requests *****//

// GetUsers fetches all the registered users from the API
func (c *ApplicationCredentials) GetUsers() (users []*User, err error) {
	bytes, err := c.query("users", "GET", nil)
	if err != nil {
		return
	}

	// Get the initial page
	var page jsonUserPage
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
		users = append(users, page.asUserSlice()...)
	}

	return
}

// GetUser fetches, if present, the user with the matching Internal ID
func (c *ApplicationCredentials) GetUser(id int) (user *User, err error) {
	bytes, err := c.query(fmt.Sprintf("users/%d", id), "GET", nil)
	if err != nil {
		return
	}

	var wrapper struct {
		User *User `json:"attributes"`
	}

	err = json.Unmarshal(bytes, &wrapper)
	if err != nil {
		return
	}

	return wrapper.User, nil
}

// GetUserExternal fetches, if present, the user with the matching External ID
func (c *ApplicationCredentials) GetUserExternal(eid string) (user *User, err error) {
	bytes, err := c.query(fmt.Sprintf("users/external/%s", eid), "GET", nil)
	if err != nil {
		return
	}

	var wrapper struct {
		User *User `json:"attributes"`
	}

	err = json.Unmarshal(bytes, &wrapper)
	if err != nil {
		return
	}

	return wrapper.User, nil
}


// CreateUser makes a new account with the provided data. The password argument can be optionally set.
func (c *ApplicationCredentials) CreateUser(u *User, password ...string) (err error) {
	type wrapper struct {
		ExternalID string `json:"external_id,omitempty"`
		Username   string `json:"username"`
		Email      string `json:"email"`
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		Password   string `json:"password,omitempty"`
		RootAdmin  bool   `json:"root_admin"`
		Language   string `json:"language,omitempty"`
	}

	var pw string
	if len(password) > 0{
		pw = password[0]
	}

	usrStruct := &wrapper{
		ExternalID: u.ExternalID,
		Username:   u.Username,
		Email:      u.Email,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		Password:   pw, 			// If empty the parser will take care of it
		RootAdmin:  u.RootAdmin,
		Language:   u.Language,
	}

	bytes, err := json.Marshal(usrStruct)
	if err != nil {
		return err
	}

	_, err = c.query("users", "POST", bytes)
	return
}

// UpdateUser modifies the user as per the passed object. Be aware that not all parameters can be
// modified. Modifiable parameters include: External ID, Username, First name, Last name, Password, Root admin
// and Language. The password parameter can be optionally set.
func (c *ApplicationCredentials) UpdateUser(u *User, password ...string) (err error) {
	type wrapper struct {
		ExternalID string `json:"external_id,omitempty"`
		Username   string `json:"username"`
		Email      string `json:"email"`
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		Password   string `json:"password,omitempty"`
		RootAdmin  bool   `json:"root_admin,omitempty"`
		Language   string `json:"language,omitempty"`
	}

	var pw string
	if len(password) > 0{
		pw = password[0]
	}

	usrStruct := &wrapper{
		ExternalID: u.ExternalID,
		Username:   u.Username,
		Email:      u.Email,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		Password:   pw, 			// If empty the parser will take care of it
		RootAdmin:  u.RootAdmin,
		Language:   u.Language,
	}

	bytes, err := json.Marshal(usrStruct)
	if err != nil {
		return err
	}

	_, err = c.query(fmt.Sprintf("users/%d", u.ID), "PATCH", bytes)
	return
}

// DeleteUser marks a user for deletion.
func (c *ApplicationCredentials) DeleteUser(id int) (err error) {
	_, err = c.query(fmt.Sprintf("users/%d", id), "DELETE", nil)
	return
}