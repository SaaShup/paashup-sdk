package docker

import (
    "github.com/SaaShup/paashup-sdk/netbox"
    "encoding/json"
    "fmt"
)

func listContainers() (ContainerList, error){
    resultCall, err := netbox.request(nil, "containers/", "GET", nil)

    if err != nil {
        return ContainerList{}, err
    }

    var result ContainerList

    if err := json.Unmarshal(resultCall, &result); err != nil { // Parse []byte to the go struct pointer
        return ContainerList{}, err
    }

    return result, nil
}

func listContainersByHost(hostId int) (ContainerList, error){
    resultCall, err := netbox.request(nil, fmt.Sprintf("containers/?host_id=%d", hostId), "GET", nil)

    if err != nil {
        return ContainerList{}, err
    }

    var result ContainerList

    if err := json.Unmarshal(resultCall, &result); err != nil { // Parse []byte to the go struct pointer
        return ContainerList{}, err
    }

    return result, nil
}
