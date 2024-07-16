package docker

import (
    "testing"
    "github.com/SaaShup/paashup-sdk/netbox"
    "os"
)

func TestVolumeList(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")

    vols, err := VolumeList()
    if err != nil {
        t.Fatal(err)
    }
    if len(vols.Results) != 2 {
        t.Fatal("No volumes found")
    }
}

func TestVolumeListByHost(t *testing.T) {
    netbox.NETBOX_URL = os.Getenv("NETBOX_URL")
    netbox.NETBOX_TOKEN = os.Getenv("NETBOX_TOKEN")
    
    vols, err := VolumeListByHost(1)
    if err != nil {
        t.Fatal(err)
    }
    if vols.Results[0].Host.Name != "testhost" {
        t.Fatal(vols.Results[0].Host.Name)
    }
}
