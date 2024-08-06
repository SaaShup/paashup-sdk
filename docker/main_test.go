package docker

import (
    "testing"
    "github.com/SaaShup/paashup-sdk/netbox"
    "os"
)

func TestHostCreate(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")

    host, err := HostCreate(HostCreateStruct{Name: "testhost", Endpoint: "http://test:test@test.com:1880/"})
    if err != nil {
        t.Fatal(err)
    }
    if host.Name != "testhost" {
        t.Fatal(host.Name)
    }
}

func TestRegistryCreate(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")

    host, err := HostSearchByName("testhost")
    registry, err := RegistryCreate(RegistryCreateStruct{Name: "testregistry", ServerAddress: "http://test.com:5000", Username: "test", Password: "test", Host: host.Id})
    if err != nil {
        t.Fatal(err)
    }
    if registry.Name != "testregistry" {
        t.Fatal(registry.Name)
    }
}

func TestImageCreate(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")

    host, err := HostSearchByName("testhost")
    registry, err := RegistrySearchByName("testregistry", host.Id)

    image, err := ImageCreate(ImageCreateStruct{Name: "testimage", Host: host.Id, Registry: registry.Id, ImageID: "0", Version: "0"})
    if err != nil {
        t.Fatal(err)
    }
    if image.Name != "testimage" {
        t.Fatal(image.Name)
    }
}

func TestContainerCreate(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")

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
    if err != nil {
        t.Fatal(err)
    }
    if container.Name != "testcontainer" {
        t.Fatal(container.Name)
    }
}

func TestVolumeCreate(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")

    host, err := HostSearchByName("testhost")

    volume, err := VolumeCreate(VolumeCreateStruct{Name: "testvolume", Host: host.Id, Driver: "local"})
    if err != nil {
        t.Fatal(err)
    }
    if volume.Name != "testvolume" {
        t.Fatal(volume.Name)
    }
}

func TestVolumeList(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")

    vols, err := VolumeList()
    if err != nil {
        t.Fatal(err)
    }
    if len(vols.Results) == 0 {
        t.Fatal("No volumes found")
    }
}

func TestVolumeListByHost(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")
   
    host, err := HostSearchByName("testhost")
    vols, err := VolumeListByHost(host.Id)
    if err != nil {
        t.Fatal(err)
    }
    if len(vols.Results) > 0 {
        if vols.Results[0].Host.Name != "testhost" {
            t.Fatal(vols.Results[0].Host.Name)
        }
    } else {
        t.Fatal("No volumes found")
    }
}

func TestContainerDelete(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")

    host, err := HostSearchByName("testhost")
    container, err := ContainerSearchByName(host, "testcontainer")
    err = ContainerDelete(container.Id)
    if err != nil {
        t.Fatal(err)
    }
    
}

func TestVolumeDelete(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")

    host, err := HostSearchByName("testhost")
    volume, err := VolumeSearchByName("testvolume", host.Id)
    err = VolumeDelete(volume.Id)
    if err != nil {
        t.Fatal(err)
    }
}

func TestImageDelete(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")

    host, err := HostSearchByName("testhost")
    image, err := ImageSearchByName("testimage", host.Id)
    err = ImageDelete(image.Id)
    if err != nil {
        t.Fatal(err)
    }
}

func TestRegistryDelete(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")
    
    host, err := HostSearchByName("testhost")
    registry, err := RegistrySearchByName("testregistry", host.Id)
    err = RegistryDelete(registry.Id)
    if err != nil {
        t.Fatal(err)
    }
}

func TestHostDelete(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")

    host, err := HostSearchByName("testhost")
    err = HostDelete(host.Id)
    if err != nil {
        t.Fatal(err)
    }

}
