package fossil

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Using the function through a variable allows stub testing
var query = queryCallback

//***** Queries *****//

func (c *ApplicationCredentials) query(endpoint, method string, data []byte) ([]byte, error) {
	target := fmt.Sprintf("%s/api/application/%s", c.URL, endpoint)
	return query(target, c.Token, method, data)
}

func (c *ClientCredentials) query(endpoint, method string, data []byte) ([]byte, error) {
	target := fmt.Sprintf("%s/api/client/%s", c.URL, endpoint)
	return query(target, c.Token, method, data)
}

func queryCallback(url, token, method string, data []byte) ([]byte, error) {
	client := &http.Client{}
	rq, _ := http.NewRequest(method, url, bytes.NewBuffer(data))

	rq.Header.Set("Authorization", "Bearer "+token)
	rq.Header.Set("Accept", "Application/vnd.pterodactyl.v1+json")
	rq.Header.Set("Content-Type", "application/json")

	rp, err := client.Do(rq)
	if err != nil {
		return nil, err
	}

	var body []byte
	// Success status range
	if rp.StatusCode < 200 || rp.StatusCode > 226 {
		// Response was a not-success code
		body, _ = ioutil.ReadAll(rp.Body)

		// No additional error info given
		if body == nil {
			return nil, errors.New("remote server responded with status " + rp.Status)
		}

		rqErr, err := parseError(body)
		// Info was there but unable to be decoded
		if err != nil {
			msg := fmt.Sprintf("remote server responded with status %s."+
				" aditionaly another error occurred while decoding the error: %s", rp.Status, err.Error())
			return nil, errors.New(msg)
		}

		// Was able to parse the error details
		msg := fmt.Sprintf("remote server responded with status %s (%s): %s",
			rqErr.Status, rqErr.Code, rqErr.Detail)
		return nil, errors.New(msg)
	}

	if rp.Body != nil {
		body, _ = ioutil.ReadAll(rp.Body)
	}

	rp.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	err = rp.Body.Close()
	if err != nil {
		return body, errors.New("unable to close response body")
	}

	return body, nil
}

//***** Errors *****//

// requestError contains error details for Pterodactyl requests errors
type requestError struct {
	Code   string `json:"code"` // For some reason the code is given as a string
	Status string `json:"status"`
	Detail string `json:"detail"`
}

// jsonRequestErrors allows the JSON provided to be decoded
type jsonRequestErrors struct {
	Errors []*requestError `json:"errors"`
}

// parseError takes a Pterodactyl-formatted error and parses it into a struct
func parseError(bytes []byte) (*requestError, error) {
	var rqErrors jsonRequestErrors
	err := json.Unmarshal(bytes, &rqErrors)
	if err != nil {
		return nil, err
	}

	if len(rqErrors.Errors) < 1 {
		return nil, errors.New("no error details given")
	}

	return rqErrors.Errors[0], nil
}
