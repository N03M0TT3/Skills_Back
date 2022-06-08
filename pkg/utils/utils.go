package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func parseBody(r *http.Request, x interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, x)
	if err != nil {
		return err
	}
	return nil
}
