package fossil

import (
	"encoding/json"
	"fmt"
	"time"
)

//***** Structures *****//

// Nest represents the information regarding a nest
type Nest struct {
	ID          int       `json:"id"`
	UUID        string    `json:"uuid"`
	Author      string    `json:"author"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// jsonNestPage is a slate to correctly parse responses from the API into JSON
type jsonNestPage struct {
	Object string `json:"object"`
	Data   []struct {
		Object     string `json:"object"`
		Nest *Nest `json:"attributes"`
	} `json:"data"`
	Meta Meta `json:"meta"`
}

// Egg represents the information regarding an egg
type Egg struct {
	ID          int    `json:"id"`
	UUID        string `json:"uuid"`
	Nest        int    `json:"nest"`
	Author      string `json:"author"`
	Description string `json:"description"`
	DockerImage string `json:"docker_image"`
	Config      struct {
		Startup struct {
			Done            string   `json:"done"`
			UserInteraction []string `json:"userInteraction"`
		} `json:"startup"`
		Stop string `json:"stop"`
		Extends string `json:"extends"`
	} `json:"config"`
	Startup string `json:"startup"`
	Script  struct {
		Privileged bool   `json:"privileged"`
		Install    string `json:"install"`
		Entry      string `json:"entry"`
		Container  string `json:"container"`
		Extends    string `json:"extends"`
	} `json:"script"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//***** Converters *****//

// asUserSlice parses a jsonUserPage into a []*User
func (np *jsonNestPage) asNestSlice() (nests []*Nest) {
	for _, n := range np.Data {
		nests = append(nests, n.Nest)
	}

	return nests
}

//***** Pagination *****//

// getAll fetches all the existing pages for a nest. The original page is kept as index 0
func (np *jsonNestPage) getAll(token string) (pages []*jsonNestPage, err error) {
	pages = append(pages, np)
	for pages[len(pages)-1].Meta.Pagination.Links.Next != "" {
		url := np.Meta.Pagination.Links.Next
		bytes, err := queryURL(url, token, "GET", nil)
		if err != nil {
			return nil, err
		}

		var page jsonNestPage
		err = json.Unmarshal(bytes, &page)
		if err != nil {
			return nil, err
		}

		pages = append(pages, &page)
	}

	return pages, nil
}

//***** Requests *****//

// GetNests fetches all available nests
func (c *ApplicationCredentials) GetNests() (nests []*Nest, err error) {
	bytes, err := c.query("nests", "GET", nil)
	if err != nil {
		return
	}

	// Get the initial page
	var page jsonNestPage
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
		nests = append(nests, page.asNestSlice()...)
	}

	return
}

// GetNest fetches, if present, a specific nest.
func (c *ApplicationCredentials) GetNest(id int) (nest *Nest, err error) {
	bytes, err := c.query(fmt.Sprintf("nests/%d", id), "GET", nil)
	if err != nil {
		return
	}

	var wrapper struct {
		Nest Nest `json:"attributes"`
	}

	err = json.Unmarshal(bytes, &wrapper)
	if err != nil {
		return
	}

	return &wrapper.Nest, nil
}

// GetEggs fetches all eggs inside a nest
func (c *ApplicationCredentials) GetEggs(nestID int) (eggs []*Egg, err error) {
	bytes, err := c.query(fmt.Sprintf("nests/%d/eggs", nestID), "GET", nil)
	if err != nil {
		return
	}

	var wrapper struct {
		Data []struct{
			Egg *Egg `json:"attributes"`
		} `json:"data"`
	}

	err = json.Unmarshal(bytes, &wrapper)
	if err != nil {
		return
	}

	for _, d := range wrapper.Data {
		eggs = append(eggs, d.Egg)
	}

	return
}

// GetEgg searches for a specific eggs inside a nest
func (c *ApplicationCredentials) GetEgg(nestID int, eggID int) (egg *Egg, err error) {
	bytes, err := c.query(fmt.Sprintf("nests/%d/eggs/%d", nestID, eggID), "GET", nil)
	if err != nil {
		return
	}

	var wrapper struct {
		Egg *Egg `json:"attributes"`
	}

	err = json.Unmarshal(bytes, &wrapper)
	if err != nil {
		return
	}

	return wrapper.Egg, nil
}