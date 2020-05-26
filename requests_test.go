package fossil

import (
	"strings"
	"testing"
)

func TestQueryCallback(t *testing.T) {
	res, err := queryCallback("https://reqbin.com/echo/get/json", "", "GET", nil)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	expect := `{"success":"true"}`

	if strings.TrimSpace(string(res)) != expect{
		t.Errorf("Response data unexpected: %s", res)
	}
}
