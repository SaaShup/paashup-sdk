# Welcome to SaaShup/paashup-sdk

This is the official SDK for Paashup. This SDK is designed to help you interact with the Paashup Platform.

## Import inside your project

```bash
import "github.com/SaaShup/paashup-sdk/docker"
import "github.com/SaaShup/paashup-sdk/netbox"
```

## Usage

```go

import (
    "fmt"
    "github.com/SaaShup/paashup-sdk/docker"
    "github.com/SaaShup/paashup-sdk/netbox"
)

func main() {
    netbox.NETBOX_URL = NETBOX_URL
    netbox.NETBOX_TOKEN = NETBOX_TOKEN

    fmt.Println(docker.ContainerInspect(containerId))
}
```
