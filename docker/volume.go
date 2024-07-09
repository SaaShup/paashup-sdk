package docker

import (
    "github.com/SaaShup/paashup-sdk/netbox"
    "encoding/json"
    "fmt"
)

func VolumeList() (VolumeListStruct, error){
    var volumeList VolumeListStruct
    result, err := netbox.Request("/volumes", "GET", nil)

    if err != nil {
        return volumeList, err
    }
    if err := json.Unmarshal(result, &volumeList); err != nil { // Parse []byte to the go struct pointer
        return volumeList, err
    }
    return volumeList, nil
}

func VolumeListByHost(hostId int) (VolumeListStruct, error){
    var volumeList VolumeListStruct
    result, err := netbox.Request(fmt.Sprintf("/volumes?host_id=%d", hostId), "GET", nil)

    if err != nil {
        return volumeList, err
    }
    if err := json.Unmarshal(result, &volumeList); err != nil {
        return volumeList, err
    }
    return volumeList, nil
}

func VolumeInspect(volumeId int) (VolumeComplete, error){
    var volume VolumeComplete
    url := fmt.Sprintf("/volumes/%d/", volumeId)
    result, err := netbox.Request(url, "GET", nil)

    if err != nil {
        return volume, err
    }
    if err := json.Unmarshal(result, &volume); err != nil {
        return volume, err
    }
    return volume, nil
}

func VolumeSearchByName(name string, hostId int) (VolumeComplete, error){
	if name == "" {
		return VolumeComplete{}, fmt.Errorf("Volume not found")
	}

	url := fmt.Sprintf("/volumes/?name=%s&host_id=%d", name, hostId)
	var result VolumeListStruct
	resultCall, err := netbox.Request(url, "GET", nil)

	if err != nil {
		return VolumeComplete{}, err
	}

	if err := json.Unmarshal(resultCall, &result); err != nil {
		return VolumeComplete{}, err
	}

	if result.Count == 1 {
		return result.Results[0], nil
	} else {
		return VolumeComplete{}, fmt.Errorf("Volume not found")
	}
}

func VolumeCreate(volume VolumeCreateStruct) (VolumeComplete, error){
    var volumeResponse VolumeComplete
    jsonStr, _ := json.Marshal(volume)
    result, err := netbox.Request("/volumes", "POST", jsonStr)

    if err != nil {
        return volumeResponse, err
    }
    if err := json.Unmarshal(result, &volumeResponse); err != nil {
        return volumeResponse, err
    }
    return volumeResponse, nil
}

func VolumeDelete(volumeId int) error{
    url := fmt.Sprintf("/volumes/%d/", volumeId)
    _, err := netbox.Request(url, "DELETE", nil)

    if err != nil {
        return err
    }
    return nil
}
