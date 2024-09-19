package docker

import (
    "testing"
    "github.com/SaaShup/paashup-sdk/netbox"
    "github.com/stretchr/testify/assert"
    "os"
)

func TestHostCreate(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")
    assert := assert.New(t)

    host, err := HostCreate(HostCreateStruct{Name: "testhost", Endpoint: "http://test:test@test.com:1880/"})
    assert.Nil(err)
    assert.Equal(host.Name, "testhost")
}

func TestRegistryCreate(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")
    assert := assert.New(t)

    host, err := HostSearchByName("testhost")
    registry, err := RegistryCreate(RegistryCreateStruct{Name: "testregistry", ServerAddress: "http://test.com:5000", Username: "test", Password: "test", Host: host.Id})
    assert.Nil(err)
    assert.Equal(registry.Name, "testregistry")
}

func TestImageCreate(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")
    assert := assert.New(t)

    host, err := HostSearchByName("testhost")
    registry, err := RegistrySearchByName("testregistry", host.Id)

    image, err := ImageCreate(ImageCreateStruct{Name: "testimage", Host: host.Id, Registry: registry.Id, ImageID: "0", Version: "0"})
    assert.Nil(err)
    assert.Equal(image.Name, "testimage")
}

func TestContainerCreate(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")
    assert := assert.New(t)

    host, err := HostSearchByName("testhost")
    image, err := ImageSearchByName("testimage", host.Id)

    containerS := ContainerCreateStruct{Name: "testcontainer", Host: host.Id, Image: image.Id, State: "none", Operation: "create"}
    containerS.Restart_policy = "no"
    containerS.Labels = []Label{}
    containerS.Env = []Env{}
    containerS.Ports = []Port{}
    containerS.Mounts = []Mount{}
    containerS.NetworkSettings = []Network{}

    container, err := ContainerCreate(containerS)
    assert.Nil(err)
    assert.Equal(container.Name, "testcontainer")
}

func TestVolumeCreate(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")
    assert := assert.New(t)

    host, err := HostSearchByName("testhost")

    volume, err := VolumeCreate(VolumeCreateStruct{Name: "testvolume", Host: host.Id, Driver: "local"})
    assert.Nil(err)
    assert.Equal(volume.Name, "testvolume")
}

func TestVolumeList(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")
    assert := assert.New(t)

    vols, err := VolumeList()
    assert.Nil(err)
    assert.NotZero(len(vols.Results))
}

func TestVolumeListByHost(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")
    assert := assert.New(t)

    host, err := HostSearchByName("testhost")
    vols, err := VolumeListByHost(host.Id)
    assert.Nil(err)
    assert.NotZero(len(vols.Results))
    assert.Equal(vols.Results[0].Host.Name, "testhost")
}

func TestContainerDelete(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")
    assert := assert.New(t)

    host, err := HostSearchByName("testhost")
    container, err := ContainerSearchByName(host, "testcontainer")
    err = ContainerDelete(container.Id)
    assert.Nil(err) 
}

func TestVolumeDelete(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")
    assert := assert.New(t)

    host, err := HostSearchByName("testhost")
    volume, err := VolumeSearchByName("testvolume", host.Id)
    err = VolumeDelete(volume.Id)
    assert.Nil(err)
}

func TestImageDelete(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")
    assert := assert.New(t)

    host, err := HostSearchByName("testhost")
    image, err := ImageSearchByName("testimage", host.Id)
    err = ImageDelete(image.Id)
    assert.Nil(err)
}

func TestRegistryDelete(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")
    assert := assert.New(t)

    host, err := HostSearchByName("testhost")
    registry, err := RegistrySearchByName("testregistry", host.Id)
    err = RegistryDelete(registry.Id)
    assert.Nil(err)
}

func TestHostDelete(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")
    assert := assert.New(t)

    host, err := HostSearchByName("testhost")
    err = HostDelete(host.Id)
    assert.Nil(err)
}
