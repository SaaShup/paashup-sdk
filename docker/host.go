package docker

import (
    "github.com/SaaShup/paashup-sdk/netbox"
    "encoding/json"
    "fmt"
)

func HostList() (HostListStruct, error){
    var hostList HostListStruct
    result, err := netbox.Request("/docker/hosts/", "GET", nil)

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

	url := fmt.Sprintf("/docker/hosts/?name=%s", name)
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

func HostCreate(Host HostCreateStruct) (HostComplete, error){
    var HostResponse HostComplete
    jsonStr, _ := json.Marshal(Host)

    result, err := netbox.Request("/docker/hosts/", "POST", jsonStr)

    if err != nil {
        return HostResponse, err
    }
    if err := json.Unmarshal(result, &HostResponse); err != nil {
        return HostResponse, err
    }
    return HostResponse, nil
}

