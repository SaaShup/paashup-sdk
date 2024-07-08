package docker

import (
    "github.com/SaaShup/paashup-sdk/netbox"
    "encoding/json"
    "fmt"
    "strings"
)

func ContainerList() (ContainerListStruct, error){
    resultCall, err := netbox.Request("containers/", "GET", nil)

    if err != nil {
        return ContainerListStruct{}, err
    }

    var result ContainerListStruct

    if err := json.Unmarshal(resultCall, &result); err != nil { // Parse []byte to the go struct pointer
        return ContainerListStruct{}, err
    }

    return result, nil
}

func ContainerListByHost(hostId int) (ContainerListStruct, error){
    resultCall, err := netbox.Request(fmt.Sprintf("containers/?host_id=%d", hostId), "GET", nil)

    if err != nil {
        return ContainerListStruct{}, err
    }

    var result ContainerListStruct

    if err := json.Unmarshal(resultCall, &result); err != nil { // Parse []byte to the go struct pointer
        return ContainerListStruct{}, err
    }

    return result, nil
}

func ContainerSearchByName(host HostComplete, containerName string) (Container, error) {
	for _, container := range host.Containers {
		if container.Name == containerName {
			return container, nil
		}
	}
	return Container{}, fmt.Errorf("Container not found")
}

func containerOperation(container Container, operation string) (Container, error){
	url := fmt.Sprintf("containers/%d/", container.Id)
	operationS := &operationType{Operation: operation}
	jsonStr, _ := json.Marshal(operationS)

	resultCall, err := netbox.Request(url, "PATCH", jsonStr)

	if err != nil {
		return Container{}, err
	}

	var result Container

	if err := json.Unmarshal(resultCall, &result); err != nil { // Parse []byte to the go struct pointer
		return Container{}, err
	}

    return result, nil
}

func ContainerStart(container Container) (Container, error){
    return containerOperation(container, "start")
}

func ContainerStop(container Container) (Container, error){
    return containerOperation(container, "stop")
}

func ContainerInspect(containerId int) (ContainerComplete, error){
	url := fmt.Sprintf("containers/%d/", containerId)
	resultCall, err := netbox.Request(url, "GET", nil)

	if err != nil {
		return ContainerComplete{}, err
	}

	var result ContainerComplete
	if err := json.Unmarshal(resultCall, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
    return result, nil
}

func ContainerExec(containerId int, cmd string) (string, error) {

	url := fmt.Sprintf("containers/%d/exec/", containerId)
	command := &command{Cmd: strings.Fields(cmd)}
	jsonStr, _ := json.Marshal(command)

	resultCall, err := netbox.Request(url, "POST", jsonStr)

	if err != nil {
		return "", err
	}

	var result Exec

	if err := json.Unmarshal(resultCall, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return result.Stdout, nil

}

func ContainerLogs(containerId int) (string, error) {
	url := fmt.Sprintf("containers/%d/logs/", containerId)
	resultCall, err := netbox.Request(url, "GET", nil)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s\n", resultCall), nil
}