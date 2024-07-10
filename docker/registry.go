package docker

import (
    "github.com/SaaShup/paashup-sdk/netbox"
    "encoding/json"
    "fmt"
)

func RegistryList() (RegistryListStruct, error){
    var RegistryList RegistryListStruct
    result, err := netbox.Request("registries/", "GET", nil)

    if err != nil {
        return RegistryList, err
    }
    if err := json.Unmarshal(result, &RegistryList); err != nil { // Parse []byte to the go struct pointer
        return RegistryList, err
    }
    return RegistryList, nil
}

func RegistryListByHost(hostId int) (RegistryListStruct, error){
    var RegistryList RegistryListStruct
    result, err := netbox.Request(fmt.Sprintf("registries/?host_id=%d", hostId), "GET", nil)

    if err != nil {
        return RegistryList, err
    }
    if err := json.Unmarshal(result, &RegistryList); err != nil {
        return RegistryList, err
    }
    return RegistryList, nil
}

func RegistryInspect(registryId int) (RegistryComplete, error){
    var Registry RegistryComplete
    url := fmt.Sprintf("registries/%d/", registryId)
    result, err := netbox.Request(url, "GET", nil)

    if err != nil {
        return Registry, err
    }
    if err := json.Unmarshal(result, &Registry); err != nil {
        return Registry, err
    }
    return Registry, nil
}

func RegistrySearchByName(name string, hostId int) (RegistryComplete, error){
	if name == "" {
		return RegistryComplete{}, fmt.Errorf("Registry not found")
	}

	url := fmt.Sprintf("registries/?name=%s&host_id=%d", name, hostId)
	var result RegistryListStruct
	resultCall, err := netbox.Request(url, "GET", nil)

	if err != nil {
		return RegistryComplete{}, err
	}

	if err := json.Unmarshal(resultCall, &result); err != nil {
		return RegistryComplete{}, err
	}

	if result.Count == 1 {
		return result.Results[0], nil
	} else {
		return RegistryComplete{}, fmt.Errorf("Registry not found")
	}
}

func RegistryCreate(Registry RegistryCreateStruct) (RegistryComplete, error){
    var RegistryResponse RegistryComplete
    jsonStr, _ := json.Marshal(Registry)
    result, err := netbox.Request("registries/", "POST", jsonStr)

    if err != nil {
        return RegistryResponse, err
    }
    if err := json.Unmarshal(result, &RegistryResponse); err != nil {
        return RegistryResponse, err
    }
    return RegistryResponse, nil
}

func RegistryDelete(registryId int) error{
    url := fmt.Sprintf("registries/%d/", registryId)
    _, err := netbox.Request(url, "DELETE", nil)

    if err != nil {
        return err
    }
    return nil
}
