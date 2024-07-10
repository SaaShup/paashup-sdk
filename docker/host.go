package docker

import (
    "github.com/SaaShup/paashup-sdk/netbox"
    "encoding/json"
    "fmt"
)

func HostList() (HostListStruct, error){
    var hostList HostListStruct
    result, err := netbox.Request("hosts/", "GET", nil)

    if err != nil {
        return hostList, err
    }
    if err := json.Unmarshal(result, &hostList); err != nil { // Parse []byte to the go struct pointer
        return HostListStruct{}, err
    }
    return hostList, nil
}

func HostSearchByName(name string) (HostComplete, error){
	if name == "" {
		return HostComplete{}, fmt.Errorf("Host not found")
	}

	url := fmt.Sprintf("hosts/?name=%s", name)
	var result HostListStruct
	resultCall, err := netbox.Request(url, "GET", nil)

	if err != nil {
		return HostComplete{}, err
	}

	if err := json.Unmarshal(resultCall, &result); err != nil { // Parse []byte to the go struct pointer
        return HostComplete{}, err
	}

	if result.Count == 1 {
		return result.Results[0], nil
	} else {
		return HostComplete{}, fmt.Errorf("Host not found")
	}
}
