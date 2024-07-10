package docker

import (
    "github.com/SaaShup/paashup-sdk/netbox"
    "encoding/json"
    "fmt"
)

func NetworkList() (NetworkListStruct, error){
    var NetworkList NetworkListStruct
    result, err := netbox.Request("networks/", "GET", nil)

    if err != nil {
        return NetworkList, err
    }
    if err := json.Unmarshal(result, &NetworkList); err != nil { // Parse []byte to the go struct pointer
        return NetworkList, err
    }
    return NetworkList, nil
}

func NetworkListByHost(hostId int) (NetworkListStruct, error){
    var NetworkList NetworkListStruct
    result, err := netbox.Request(fmt.Sprintf("networks/?host_id=%d", hostId), "GET", nil)

    if err != nil {
        return NetworkList, err
    }
    if err := json.Unmarshal(result, &NetworkList); err != nil {
        return NetworkList, err
    }
    return NetworkList, nil
}

func NetworkInspect(networkId int) (NetworkComplete, error){
    var Network NetworkComplete
    url := fmt.Sprintf("networks/%d/", networkId)
    result, err := netbox.Request(url, "GET", nil)

    if err != nil {
        return Network, err
    }
    if err := json.Unmarshal(result, &Network); err != nil {
        return Network, err
    }
    return Network, nil
}

func NetworkSearchByName(name string, hostId int) (NetworkComplete, error){
	if name == "" {
		return NetworkComplete{}, fmt.Errorf("Network not found")
	}

	url := fmt.Sprintf("networks/?name=%s&host_id=%d", name, hostId)
	var result NetworkListStruct
	resultCall, err := netbox.Request(url, "GET", nil)

	if err != nil {
		return NetworkComplete{}, err
	}

	if err := json.Unmarshal(resultCall, &result); err != nil {
		return NetworkComplete{}, err
	}

	if result.Count == 1 {
		return result.Results[0], nil
	} else {
		return NetworkComplete{}, fmt.Errorf("Network not found")
	}
}

func NetworkCreate(Network NetworkCreateStruct) (NetworkComplete, error){
    var NetworkResponse NetworkComplete
    jsonStr, _ := json.Marshal(Network)
    result, err := netbox.Request("networks/", "POST", jsonStr)

    if err != nil {
        return NetworkResponse, err
    }
    if err := json.Unmarshal(result, &NetworkResponse); err != nil {
        return NetworkResponse, err
    }
    return NetworkResponse, nil
}

func NetworkDelete(networkId int) error{
    url := fmt.Sprintf("networks/%d/", networkId)
    _, err := netbox.Request(url, "DELETE", nil)

    if err != nil {
        return err
    }
    return nil
}
