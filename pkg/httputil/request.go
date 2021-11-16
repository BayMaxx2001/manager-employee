package httputil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func BindJsonReq(r *http.Request, obj interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return json.Unmarshal(b, obj)
}

func BindJSON(obj interface{}) string {
	out, err := json.Marshal(obj)
	if err != nil {
		return ""
	}

	return string(out)
}
