package docker

import (
    "github.com/SaaShup/paashup-sdk/netbox"
    "encoding/json"
    "fmt"
)

func ImageList() (ImageListStruct, error){
    var ImageList ImageListStruct
    result, err := netbox.Request("/docker/images/", "GET", nil)

    if err != nil {
        return ImageList, err
    }
    if err := json.Unmarshal(result, &ImageList); err != nil { // Parse []byte to the go struct pointer
        return ImageList, err
    }
    return ImageList, nil
}

func ImageListByHost(hostId int) (ImageListStruct, error){
    var ImageList ImageListStruct
    result, err := netbox.Request(fmt.Sprintf("/docker/images/?host_id=%d", hostId), "GET", nil)

    if err != nil {
        return ImageList, err
    }
    if err := json.Unmarshal(result, &ImageList); err != nil {
        return ImageList, err
    }
    return ImageList, nil
}

func ImageInspect(ImageId int) (ImageComplete, error){
    var Image ImageComplete
    url := fmt.Sprintf("/docker/images/%d/", ImageId)
    result, err := netbox.Request(url, "GET", nil)

    if err != nil {
        return Image, err
    }
    if err := json.Unmarshal(result, &Image); err != nil {
        return Image, err
    }
    return Image, nil
}

func ImageSearchByName(name string, hostId int) (ImageComplete, error){
	if name == "" {
		return ImageComplete{}, fmt.Errorf("Image not found")
	}

	url := fmt.Sprintf("/docker/images/?name=%s&host_id=%d", name, hostId)
	var result ImageListStruct
	resultCall, err := netbox.Request(url, "GET", nil)

	if err != nil {
		return ImageComplete{}, err
	}

	if err := json.Unmarshal(resultCall, &result); err != nil {
		return ImageComplete{}, err
	}

	if result.Count == 1 {
		return result.Results[0], nil
	} else {
		return ImageComplete{}, fmt.Errorf("Image not found")
	}
}

func ImageCreate(Image ImageCreateStruct) (ImageComplete, error){
    var ImageResponse ImageComplete
    jsonStr, _ := json.Marshal(Image)
    result, err := netbox.Request("/docker/images/", "POST", jsonStr)

    if err != nil {
        return ImageResponse, err
    }
    if err := json.Unmarshal(result, &ImageResponse); err != nil {
        return ImageResponse, err
    }
    return ImageResponse, nil
}

func ImageDelete(ImageId int) error{
    url := fmt.Sprintf("/docker/images/%d/", ImageId)
    _, err := netbox.Request(url, "DELETE", nil)

    if err != nil {
        return err
    }
    return nil
}
