package fossil

import (
	"encoding/json"
	"fmt"
	"time"
)

//***** Structures *****//

// Location represents a server location
type Location struct {
	ID        int       `json:"id"`
	ShortName string    `json:"short"`
	LongName  string    `json:"long"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// jsonLocationPage is a slate to correctly parse responses from the API into JSON
type jsonLocationPage struct {
	Data []struct {
		Location *Location `json:"attributes"`
	} `json:"data"`
	Meta Meta `json:"meta"`
}

//***** Converters *****//

// asUserSlice parses a jsonUserPage into a []*User
func (lp *jsonLocationPage) asLocationSlice() (locations []*Location) {
	for _, l := range lp.Data {
		locations = append(locations, l.Location)
	}

	return locations
}

//***** Pagination *****//

// getAll fetches all the existing pages for a location. The original page is kept as index 0
func (lp *jsonLocationPage) getAll(token string) (pages []*jsonLocationPage, err error) {
	pages = append(pages, lp)
	for pages[len(pages)-1].Meta.Pagination.Links.Next != "" {
		url := lp.Meta.Pagination.Links.Next
		bytes, err := queryURL(url, token, "GET", nil)
		if err != nil {
			return nil, err
		}

		var page jsonLocationPage
		err = json.Unmarshal(bytes, &page)
		if err != nil {
			return nil, err
		}

		pages = append(pages, &page)
	}

	return pages, nil
}

//***** Requests *****//

// GetLocations fetches all available locations
func (c *ApplicationCredentials) GetLocations() (locations []*Location, err error) {
	bytes, err := c.query("locations", "GET", nil)
	if err != nil {
		return
	}

	// Get the initial page
	var page jsonLocationPage
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
		locations = append(locations, page.asLocationSlice()...)
	}

	return
}

// GetLocation fetches the location with the given ID
func (c *ApplicationCredentials) GetLocation(id int) (loc *Location, err error) {
	bytes, err := c.query(fmt.Sprintf("locations/%d", id), "GET", nil)
	if err != nil {
		return
	}

	var wrapper struct {
		Location *Location `json:"attributes"`
	}

	err = json.Unmarshal(bytes, &wrapper)
	if err != nil {
		return
	}

	return wrapper.Location, nil
}

// CreateLocation makes a new location with the provided names
func (c *ApplicationCredentials) CreateLocation(shortName, longName string) (loc *Location, err error) {
	type names struct {
		Short string `json:"short"`
		Long  string `json:"long"`
	}

	b := names{
		Short: shortName,
		Long:  longName,
	}

	nms, err := json.Marshal(b)
	if err != nil {
		return
	}

	bytes, err := c.query("locations", "POST", nms)
	if err != nil {
		return
	}

	var fWrapper struct {
		Location *Location `json:"attributes"`
	}

	err = json.Unmarshal(bytes, &fWrapper)
	if err != nil {
		return
	}

	return fWrapper.Location, nil
}

// UpdateLocationName modifies the short and long names of a location
func (c *ApplicationCredentials) UpdateLocationName(loc *Location) (err error) {
	type names struct {
		Short string `json:"short"`
		Long  string `json:"long"`
	}

	b := names{
		Short: loc.ShortName,
		Long:  loc.LongName,
	}

	nms, err := json.Marshal(b)
	if err != nil {
		return
	}

	_, err = c.query(fmt.Sprintf("locations/%d", loc.ID), "PATCH", nms)
	return
}

// DeleteLocation marks a server as suspended
func (c *ApplicationCredentials) DeleteLocation(lid int) (err error) {
	_, err = c.query(fmt.Sprintf("locations/%d", lid), "DELETE", nil)
	return
}
