package docker

type Exec struct {
	Stdout string `json:"stdout"`
}

type Host struct {
	Id                 int    `json:"id" yaml:"id"`
	Url                string `json:"url" yaml:"url"`
	Display            string `json:"display" yaml:"display"`
	Name               string `json:"name" yaml:"name"`
	State              string `json:"state" yaml:"state"`
	Agent_version      string `json:"agent_version" yaml:"agent_version"`
	Docker_api_version string `json:"docker_api_version" yaml:"docker_api_version"`
	Endpoint           string `json:"endpoint" yaml:"endpoint"`
}

type Image struct {
	Id      int    `json:"id" yaml:"id"`
	Url     string `json:"url" yaml:"url"`
	Display string `json:"display" yaml:"display"`
	Name    string `json:"name" yaml:"name"`
	Version string `json:"version" yaml:"version"`
	Size    int    `json:"size" yaml:"size"`
	ImageID string `json:"ImageID" yaml:"ImageID"`
	Digest  string `json:"Digest" yaml:"Digest"`
}

type Volume struct {
	Id      int    `json:"id" yaml:"id"`
	Url     string `json:"url" yaml:"url"`
	Display string `json:"display" yaml:"display"`
	Name    string `json:"name" yaml:"name"`
	Driver  string `json:"driver" yaml:"driver"`
}

type Network struct {
	Id        int    `json:"id" yaml:"id"`
	Url       string `json:"url" yaml:"url"`
	Display   string `json:"display" yaml:"display"`
	Name      string `json:"name" yaml:"name"`
	Driver    string `json:"driver" yaml:"driver"`
	NetworkID string `json:"NetworkID" yaml:"NetworkID"`
	State     string `json:"state" yaml:"state"`
}

type Registry struct {
	Id            int    `json:"id" yaml:"id"`
	Url           string `json:"url" yaml:"url"`
	Display       string `json:"display" yaml:"display"`
	Name          string `json:"name" yaml:"name"`
	ServerAddress string `json:"serveraddress" yaml:"serveraddress"`
	Username      string `json:"username" yaml:"username"`
	Password      string `json:"password" yaml:"password"`
	Email         string `json:"email" yaml:"email"`
}

type Container struct {
	Id             int    `json:"id" yaml:"id"`
	Url            string `json:"url" yaml:"url"`
	Display        string `json:"display" yaml:"display"`
	Name           string `json:"name" yaml:"name"`
	ContainerID    string `json:"ContainerID" yaml:"ContainerID"`
	State          string `json:"state" yaml:"state"`
	Status         string `json:"status" yaml:"status"`
	Restart_policy string `json:"restart_policy" yaml:"restart_policy"`
	Operation      string `json:"operation" yaml:"operation"`
	Hostname       string `json:"hostname" yaml:"hostname"`
}

type Mount struct {
	Source    string `json:"source" yaml:"source"`
	Read_only bool   `json:"read_only" yaml:"read_only"`
	Volume    Volume `json:"volume" yaml:"volume"`
}

type VolumeComplete struct {
    Volume
    Mounts []Mount `json:"mounts" yaml:"mounts"`
    Host  Host    `json:"host" yaml:"host"`
    Custom_fields CustomField `json:"custom_fields" yaml:"custom_fields"`
    Tags []Tag `json:"tags" yaml:"tags"`
}

type VolumeCreateStruct struct {
    Name   string `json:"name" yaml:"name"`
    Driver string `json:"driver" yaml:"driver"`
    Host   int    `json:"host" yaml:"host"`
}

type VolumeListStruct struct {
    Count    int              `json:"count" yaml:"count"`
    Next     string           `json:"next" yaml:"next"`
    Previous string           `json:"previous" yaml:"previous"`
    Results  []VolumeComplete `json:"results" yaml:"results"`
}

type NetworkComplete struct {
    Network
    Host  Host  `json:"host" yaml:"host"`
    network_settings []struct {
        Id     int    `json:"id" yaml:"id"`
        Network Network `json:"network" yaml:"network"`
    }
    Custom_fields CustomField `json:"custom_fields" yaml:"custom_fields"`
    Tags []Tag `json:"tags" yaml:"tags"`
}

type NetworkCreateStruct struct {
    Name   string `json:"name" yaml:"name"`
    Driver string `json:"driver" yaml:"driver"`
    Host   int    `json:"host" yaml:"host"`
    NetworkId string `json:"NetworkID" yaml:"NetworkID"`
    State string `json:"state" yaml:"state"`
}

type NetworkListStruct struct {
    Count    int              `json:"count" yaml:"count"`
    Next     string           `json:"next" yaml:"next"`
    Previous string           `json:"previous" yaml:"previous"`
    Results  []NetworkComplete `json:"results" yaml:"results"`
}

type RegistryComplete struct {
    Registry
    Host  Host  `json:"host" yaml:"host"`
    Images []Image `json:"images" yaml:"images"`
}

type RegistryCreateStruct struct {
    Name          string `json:"name" yaml:"name"`
    ServerAddress string `json:"serveraddress" yaml:"serveraddress"`
    Username      string `json:"username" yaml:"username"`
    Password      string `json:"password" yaml:"password"`
    Email         string `json:"email" yaml:"email"`
    Host          int    `json:"host" yaml:"host"`
}

type RegistryListStruct struct {
    Count    int              `json:"count" yaml:"count"`
    Next     string           `json:"next" yaml:"next"`
    Previous string           `json:"previous" yaml:"previous"`
    Results  []RegistryComplete `json:"results" yaml:"results"`
}

type ImageComplete struct {
    Image
    Host  Host  `json:"host" yaml:"host"`
    Containers []Container `json:"containers" yaml:"containers"`
    Custom_fields CustomField `json:"custom_fields" yaml:"custom_fields"`
    Tags []Tag `json:"tags" yaml:"tags"`
    Registry Registry `json:"registry" yaml:"registry"`
}

type ImageCreateStruct struct {
    Name    string `json:"name" yaml:"name"`
    Version string `json:"version" yaml:"version"`
    Size    int    `json:"size" yaml:"size"`
    ImageID string `json:"ImageID" yaml:"ImageID"`
    Digest  string `json:"Digest" yaml:"Digest"`
    Host    int    `json:"host" yaml:"host"`
    Registry int    `json:"registry" yaml:"registry"`
}

type ImageListStruct struct {
    Count    int              `json:"count" yaml:"count"`
    Next     string           `json:"next" yaml:"next"`
    Previous string           `json:"previous" yaml:"previous"`
    Results  []ImageComplete `json:"results" yaml:"results"`
}

type Port struct {
    Public_port  int    `json:"public_port" yaml:"public_port"`
    Private_port int    `json:"private_port" yaml:"private_port"`
    Type         string `json:"type" yaml:"type"`
}

type Env struct {
    Var_name string `json:"var_name" yaml:"var_name"`
    Value    string `json:"value" yaml:"value"`
}

type Label struct {
    Key   string `json:"key" yaml:"key"`
    Value string `json:"value" yaml:"value"`
}

type ContainerComplete struct {
	Container
    Ports []Port `json:"ports" yaml:"ports"`
	Env  []Env   `json:"env" yaml:"env"`
    Labels []Label `json:"labels" yaml:"labels"`
	Mounts []Mount `json:"mounts" yaml:"mounts"`
	Binds []struct {
		Host_path      string `json:"host_path" yaml:"host_path"`
		Container_path string `json:"container_path" yaml:"container_path"`
		Read_only      bool   `json:"read_only" yaml:"read_only"`
	}
	Network_settings []struct {
		Network Network `json:"network" yaml:"network"`
	}
	Created       string      `json:"created" yaml:"created"`
	Custom_fields CustomField `json:"custom_fields" yaml:"custom_fields"`

	Last_updated string   `json:"last_updated" yaml:"last_updated"`
	Tags         []string `json:"tags" yaml:"tags"`

	Host  Host  `json:"host" yaml:"host"`
	Image Image `json:"image" yaml:"image"`
}

type ContainerCreateStruct struct {
    Name           string `json:"name" yaml:"name"`
    ContainerID    string `json:"ContainerID" yaml:"ContainerID"`
    State          string `json:"state" yaml:"state"`
    Status         string `json:"status" yaml:"status"`
    Restart_policy string `json:"restart_policy" yaml:"restart_policy"`
    Operation      string `json:"operation" yaml:"operation"`
    Host           int    `json:"host" yaml:"host"`
    Image          int    `json:"image" yaml:"image"`
    NetworkSettings []Network `json:"network_settings" yaml:"network_settings"`
    Mounts         []Mount `json:"mounts" yaml:"mounts"`
    Binds          []struct {
        Host_path      string `json:"host_path" yaml:"host_path"`
        Container_path string `json:"container_path" yaml:"container_path"`
        Read_only      bool   `json:"read_only" yaml:"read_only"`
    }
    Env             []Env `json:"env" yaml:"env"`
    Labels          []Label `json:"labels" yaml:"labels"`
    Ports           []Port `json:"ports" yaml:"ports"`
}

type ContainerListStruct struct {
	Count    int                 `json:"count" yaml:"count"`
	Next     string              `json:"next" yaml:"next"`
	Previous string              `json:"previous" yaml:"previous"`
	Results  []ContainerComplete `json:"results" yaml:"results"`
}

type CustomField interface{}

type Tag struct {
    Name string `json:"name" yaml:"name"`
    Slug string `json:"slug" yaml:"slug"`
    Color string `json:"color" yaml:"color"`
}

type HostComplete struct {
	Token struct {
		Id            int    `json:"id" yaml:"id"`
		Url           string `json:"url" yaml:"url"`
		Display       string `json:"display" yaml:"display"`
		Key           string `json:"key" yaml:"key"`
		Write_enabled bool   `json:"write_enabled" yaml:"write_enabled"`
	}
	Netbox_base_url string      `json:"netbox_base_url" yaml:"netbox_base_url"`
	Custom_fields   CustomField `json:"custom_fields" yaml:"custom_fields"`
	Last_updated    string      `json:"last_updated" yaml:"last_updated"`
	Tags            []string    `json:"tags" yaml:"tags"`
	Images          []Image     `json:"images" yaml:"images"`
	Volumes         []Volume    `json:"volumes" yaml:"volumes"`
	Networks        []Network   `json:"networks" yaml:"networks"`
	Containers      []Container `json:"containers" yaml:"containers"`
	Registries      []Registry  `json:"registries" yaml:"registries"`
	Host
}

type HostCreateStruct struct {
    Endpoint string `json:"endpoint" yaml:"endpoint"`
    Name     string `json:"name" yaml:"name"`
}

type HostListStruct struct {
	Count    int            `json:"count" yaml:"count"`
	Next     string         `json:"next" yaml:"next"`
	Previous string         `json:"previous" yaml:"previous"`
	Results  []HostComplete `json:"results" yaml:"results"`
}

type operationType struct {
	Operation string `json:"operation"`
}

type command struct {
	Cmd []string `json:"cmd"`
}
