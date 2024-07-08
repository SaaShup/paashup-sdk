package netbox

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strings"

)

var NETBOX_URL string
var NETBOX_TOKEN string

func Request(endpoint string, method string, jsonStr []byte) ([]byte, error) {
	netboxUrl := strings.TrimRight(NETBOX_URL, "/")
	client := &http.Client{}

	req, err := http.NewRequest(method, fmt.Sprintf("%s/api/plugins/docker/%s", netboxUrl, endpoint), ioutil.NopCloser(bytes.NewBuffer(jsonStr)))

	if err != nil {
		return nil, err
	}

	req.ContentLength = int64(len(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", NETBOX_TOKEN))
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	// ...
	return ioutil.ReadAll(res.Body)
}
